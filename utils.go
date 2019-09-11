package alipay

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"sort"
	"strings"
)

// 组成签名raw串
func getSignatureContent(m map[string]string) string {
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
func parseCertificate(s string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(strings.TrimSpace(s)))
	if block == nil {
		return nil, fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: " + err.Error())
	}

	return cert, nil
}

// 公共的验签方法
// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func RsaCheckV1(params map[string]string, publicKey, charset, signType string) (bool, error) {
	signChecker, err := NewSignChecker([]byte(publicKey))
	if err != nil {
		return false, err
	}

	sign := params["sign"]
	delete(params, "sign")
	delete(params, "sign_type")

	return signChecker.Check(getSignatureContent(params), sign, signType, charset)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func RsaCheckV2(params map[string]string, publicKey, charset, signType string) (bool, error) {
	signChecker, err := NewSignChecker([]byte(publicKey))
	if err != nil {
		return false, err
	}

	sign := params["sign"]
	delete(params, "sign")

	return signChecker.Check(getSignatureContent(params), sign, signType, charset)
}

// TODO: 公钥证书验签，等sm2椭圆曲线先解决 --
// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
//func RsaCertCheckV1(params map[string]string, alipayPublicCertPath, charset, signType string) (bool, error) {
//	sn := params["alipay_cert_sn"]
//
//	k, ok := c.alipayPublicKeyMap[sn]
//	if !ok && params["sub_code"] == "" {
//		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
//	}
//
//	sign := params["sign"]
//	delete(params, "sign")
//	delete(params, "sign_type")
//
//	return NewSignCheckerWithPublicKey(k).Check(getSignatureContent(params), sign, signType, charset)
//}
//
//// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
//func RsaCertCheckV2(params map[string]string, alipayPublicCertPath, charset, signType string) (bool, error) {
//	sn := params["alipay_cert_sn"]
//
//	k, ok := c.alipayPublicKeyMap[sn]
//	if !ok && params["sub_code"] == "" {
//		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
//	}
//
//	sign := params["sign"]
//	delete(params, "sign")
//
//	return NewSignCheckerWithPublicKey(k).Check(getSignatureContent(params), sign, signType, charset)
//}
