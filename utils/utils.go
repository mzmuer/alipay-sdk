package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rsa"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/mzmuer/alipay-sdk/signature"
	"github.com/tjfoc/gmsm/sm2"
)

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
func ParseCertificate(s string) (*sm2.Certificate, error) {
	block, _ := pem.Decode([]byte(strings.TrimSpace(s)))
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := sm2.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: " + err.Error())
	}

	return cert, nil
}

// 公共的验签方法
// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func RsaCheckV1(params map[string]string, publicKey, charset, signType string) (bool, error) {
	sign := params["sign"]
	delete(params, "sign")
	delete(params, "sign_type")

	return _rsaCheckV2(params, publicKey, charset, signType, sign)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func RsaCheckV2(params map[string]string, publicKey, charset, signType string) (bool, error) {
	sign := params["sign"]
	delete(params, "sign")

	return _rsaCheckV2(params, publicKey, charset, signType, sign)
}

func _rsaCheckV2(params map[string]string, publicKey, charset, signType, sign string) (bool, error) {
	signChecker, err := sign2.NewSignChecker([]byte(publicKey))
	if err != nil {
		return false, err
	}

	return signChecker.Check(getSignatureContent(params), sign, signType, charset)
}

//此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func RsaCertCheckV1(params map[string]string, alipayPublicCertPath, charset, signType string) (bool, error) {
	sign := params["sign"]
	delete(params, "sign")
	delete(params, "sign_type")

	return _rsaCertCheck(params, alipayPublicCertPath, charset, signType, sign)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func RsaCertCheckV2(params map[string]string, alipayPublicCertPath, charset, signType string) (bool, error) {
	sign := params["sign"]
	delete(params, "sign")

	return _rsaCertCheck(params, alipayPublicCertPath, charset, signType, sign)
}

// --
func _rsaCertCheck(params map[string]string, alipayPublicCertPath, charset, signType, sign string) (bool, error) {
	b, err := ioutil.ReadFile(alipayPublicCertPath)
	if err != nil {
		return false, err
	}

	cert, err := parseCertificate(string(b))
	if err != nil {
		return false, err
	}

	if params["alipay_cert_sn"] != getCertSN(cert) {
		return false, fmt.Errorf("支付宝公钥证书SN不匹配")
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return false, fmt.Errorf("支付宝公钥证书类型错误，无法获取到public key")
	}

	return sign2.NewSignCheckerWithPublicKey(key).Check(getSignatureContent(params), sign, signType, charset)
}

// 获取证书的sn
func GetCertSN(cert *sm2.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
