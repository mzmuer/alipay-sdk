package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rsa"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/mzmuer/alipay-sdk/constant"
	"github.com/mzmuer/alipay-sdk/signature"
	sm2x509 "github.com/mzmuer/gmsm/x509"
)

var (
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersLength := len(letters)

	result := make([]rune, ln)

	for i := range result {
		result[i] = letters[rander.Intn(lettersLength)]
	}

	return string(result)
}

// 组成签名raw串
func GetSignatureContent(m map[string]string) string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	// 对keys排序
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, key := range keys {
		if m[key] == "" {
			continue
		}

		if i != 0 {
			buf.WriteString("&")
		}

		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(m[key])
	}

	return buf.String()
}

// 解析证书
func ParseCertificate(s string) (*sm2x509.Certificate, error) {
	block, _ := pem.Decode([]byte(strings.TrimSpace(s)))
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := sm2x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: " + err.Error())
	}

	return cert, nil
}

// 公共的验签方法
// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func RsaCheckV1(params map[string]string, publicKey, charset, signType string) (bool, error) {
	sign := params[constant.Sign]
	delete(params, constant.Sign)
	delete(params, constant.SignType)

	return _rsaCheckV2(params, publicKey, charset, signType, sign)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func RsaCheckV2(params map[string]string, publicKey, charset, signType string) (bool, error) {
	sign := params[constant.Sign]
	delete(params, constant.Sign)

	return _rsaCheckV2(params, publicKey, charset, signType, sign)
}

func _rsaCheckV2(params map[string]string, publicKey, charset, signType, sign string) (bool, error) {
	signChecker, err := signature.NewSignChecker([]byte(publicKey))
	if err != nil {
		return false, err
	}

	return signChecker.Check(GetSignatureContent(params), sign, signType, charset)
}

//此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func RsaCertCheckV1(params map[string]string, alipayPublicCertPath, charset, signType string) (bool, error) {
	sign := params[constant.Sign]
	delete(params, constant.Sign)
	delete(params, constant.SignType)

	return _rsaCertCheck(params, alipayPublicCertPath, charset, signType, sign)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func RsaCertCheckV2(params map[string]string, alipayPublicCertPath, charset, signType string) (bool, error) {
	sign := params[constant.Sign]
	delete(params, constant.Sign)

	return _rsaCertCheck(params, alipayPublicCertPath, charset, signType, sign)
}

// --
func _rsaCertCheck(params map[string]string, alipayPublicCertPath, charset, signType, sign string) (bool, error) {
	b, err := ioutil.ReadFile(alipayPublicCertPath)
	if err != nil {
		return false, err
	}

	cert, err := ParseCertificate(string(b))
	if err != nil {
		return false, err
	}

	if params["alipay_cert_sn"] != GetCertSN(cert) {
		return false, fmt.Errorf("支付宝公钥证书SN不匹配")
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return false, fmt.Errorf("支付宝公钥证书类型错误，无法获取到public key")
	}

	return signature.NewSignCheckerWithPublicKey(key).Check(GetSignatureContent(params), sign, signType, charset)
}

// 获取证书的sn
func GetCertSN(cert *sm2x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
