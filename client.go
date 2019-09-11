package alipay

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/mzmuer/alipay-sdk/request"
	"github.com/mzmuer/alipay-sdk/response"
)

type Client struct {
	AppId       string
	Charset     string // utf-8
	SignType    string // RSA2
	Signer      *signer
	SignChecker *signChecker
	Format      string // json
	EncryptType string // AES

	// 公钥证书相关
	appPubCertSN       string
	alipayRootCertSN   string
	alipayPubCertSN    string
	alipayPublicKeyMap map[string]*rsa.PublicKey

	isSandBox bool
}

func NewClient(appId string, publicKey, privateKey string, isSandBox bool) (*Client, error) {
	signChecker, err := NewSignChecker([]byte(publicKey))
	if err != nil {
		return nil, err
	}

	signer, err := NewSigner([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	c := Client{
		AppId:       appId,
		Charset:     "utf-8",
		SignChecker: signChecker,
		SignType:    SignTypeRSA2,
		Format:      "json",
		EncryptType: EncryptTypeAes,
		Signer:      signer,
		isSandBox:   isSandBox,
	}

	return &c, nil
}

func NewCertClient(appId, privateKey, appPubCert, alipayRootCert, alipayPubCert string, isSandBox bool) (*Client, error) {
	signer, err := NewSigner([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	c := Client{
		AppId:              appId,
		Charset:            "utf-8",
		SignType:           SignTypeRSA2,
		Format:             "json",
		EncryptType:        EncryptTypeAes,
		Signer:             signer,
		isSandBox:          isSandBox,
		alipayPublicKeyMap: make(map[string]*rsa.PublicKey),
	}

	err = c.loadAppPubCertSN(appPubCert)
	if err != nil {
		return nil, err
	}

	err = c.loadAliPayRootCert(alipayRootCert)
	if err != nil {
		return nil, err
	}

	err = c.loadAliPayPublicCert(alipayPubCert)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Client) Execute(r request.Request, result response.Response) (string, error) {
	return c._execute(r, result, "", "")
}

func (c *Client) ExecuteP1(r request.Request, result response.Response, accessToken string) (string, error) {
	return c._execute(r, result, accessToken, "")
}

func (c *Client) ExecuteP2(r request.Request, result response.Response, accessToken, appAuthToken string) (string, error) {
	return c._execute(r, result, accessToken, appAuthToken)
}

func (c *Client) _execute(r request.Request, result response.Response, accessToken, appAuthToken string) (string, error) {
	// 构造请求map请求
	requestParams, err := c.getRequestHolderWithSign(r, accessToken, appAuthToken)
	if err != nil {
		return "", err
	}

	gateway := Gateway
	if c.isSandBox {
		gateway = SandboxGateway
	}

	b, err := doPost(gateway, requestParams)
	if err != nil {
		return "", err
	}

	if err = response.ParseResponse(r.GetMethod(), b, result); err != nil {
		return string(b), err
	}

	if result.GetSign() != "" {
		if result.GetAlipayCertSn() != "" {
			_, err = c.checkCertResponseSign(result)
		} else {
			_, err = c.checkResponseSign(result)
		}

		if err != nil {
			return string(b), err
		}
		//if !match { // 签名不匹配
		//	return "", fmt.Errorf("sign check fail: check Sign and Data Fail")
		//}
	}

	return "", nil
}

// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func (c *Client) RsaCheckV1(params map[string]string, charset, signType string) (bool, error) {
	if c.SignChecker == nil {
		return true, nil
	}

	sign := params["sign"]
	delete(params, "sign")
	delete(params, "sign_type")

	return c.SignChecker.Check(getSignatureContent(params), sign, signType, charset)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func (c *Client) RsaCheckV2(params map[string]string, charset, signType string) (bool, error) {
	if c.SignChecker == nil {
		return true, nil
	}

	sign := params["sign"]
	delete(params, "sign")

	return c.SignChecker.Check(getSignatureContent(params), sign, signType, charset)
}

// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func (c *Client) RsaCertCheckV1(params map[string]string, charset, signType string) (bool, error) {
	sn := params["alipay_cert_sn"]

	k, ok := c.alipayPublicKeyMap[sn]
	if !ok && params["sub_code"] == "" {
		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
	}

	sign := params["sign"]
	delete(params, "sign")
	delete(params, "sign_type")

	return NewSignCheckerWithPublicKey(k).Check(getSignatureContent(params), sign, signType, charset)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func (c *Client) RsaCertCheckV2(params map[string]string, charset, signType string) (bool, error) {
	sn := params["alipay_cert_sn"]

	k, ok := c.alipayPublicKeyMap[sn]
	if !ok && params["sub_code"] == "" {
		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
	}

	sign := params["sign"]
	delete(params, "sign")

	return NewSignCheckerWithPublicKey(k).Check(getSignatureContent(params), sign, signType, charset)
}

// 构造请求map
func (c *Client) getRequestHolderWithSign(r request.Request, accessToken, appAuthToken string) (map[string]string, error) {
	params := map[string]string{}

	// 必选参数
	params[Method] = r.GetMethod()
	params[Version] = r.GetApiVersion()
	params[AppId] = c.AppId
	params[SignType] = c.SignType
	params[TerminalType] = r.GetTerminalType()
	params[TerminalInfo] = r.GetTerminalInfo()
	params[NotifyUrl] = r.GetNotifyUrl()
	params[ReturnUrl] = r.GetReturnUrl()
	params[Charset] = c.Charset
	params[Timestamp] = time.Now().Format("2006-01-02 15:03:04")
	if r.GetNeedEncrypt() {
		params[EncryptType] = c.EncryptType
	}

	if c.appPubCertSN != "" {
		params[AppCertSn] = c.appPubCertSN
	}

	if c.alipayRootCertSN != "" {
		params[AlipayRootCertSn] = c.alipayRootCertSN
	}

	// 可选参数
	params[Format] = c.Format
	params[AccessToken] = accessToken
	params[AlipaySdk] = SdkVersion
	params[ProdCode] = r.GetProdCode()

	// app参数
	// 必须先添加额外参数
	for key, v := range r.GetTextParams() {
		params[key] = v
	}

	if params[BizContentKey] == "" && r.GetBizModel() != nil {
		bizContent, err := json.Marshal(r.GetBizModel())
		if err != nil {
			return nil, err
		}

		params[BizContentKey] = string(bizContent)
	}

	if r.GetNeedEncrypt() {
		if c.EncryptType == "" {
			return nil, fmt.Errorf("加密类型错误")
		}

		params[EncryptType] = c.EncryptType
		// TODO: 对r.BizContent一波加密操作
		// params[BizContentKey] = encryptContent
	}

	if appAuthToken != "" {
		params[AppAuthToken] = appAuthToken
	}

	// 签名 - 必选参数
	if c.SignType != "" {
		var err error
		signContent := getSignatureContent(params)
		params["sign"], err = c.Signer.Sign(signContent, c.SignType, c.Charset)
		if err != nil {
			return nil, err
		}
	} else {
		params["sign"] = ""
	}

	return params, nil
}

// --
func (c *Client) checkResponseSign(resp response.Response) (bool, error) {
	if c.SignChecker == nil {
		return true, nil
	}

	return c.SignChecker.Check(resp.GetRawParams(), resp.GetSign(), c.SignType, c.Charset)
}

func (c *Client) checkCertResponseSign(resp response.Response) (bool, error) {
	k, ok := c.alipayPublicKeyMap[resp.GetAlipayCertSn()]
	if !ok && resp.IsSuccess() {
		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
	}

	return NewSignCheckerWithPublicKey(k).Check(resp.GetRawParams(), resp.GetSign(), c.SignType, c.Charset)
}

// 加载应用公钥证书sn
func (c *Client) loadAppPubCertSN(s string) error {
	cert, err := parseCertificate(s)
	if err != nil {
		return err
	}

	c.appPubCertSN = _getCertSN(cert)
	return nil
}

// 加载支付宝根证书sn
func (c *Client) loadAliPayRootCert(s string) error {
	var (
		certStrList = strings.SplitAfter(s, "-----END CERTIFICATE-----")
		certSNList  = make([]string, 0, len(certStrList))
	)

	for _, v := range certStrList {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		block, _ := pem.Decode([]byte(v))
		if block == nil {
			return fmt.Errorf("failed to parse certificate PEM")
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			// TODO: 暂时先忽略错误，第一个证书是sm2的椭圆曲线，go本身不支持
			//fmt.Println(err)
			//return fmt.Errorf("failed to parse certificate: " + err.Error())
		}

		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			certSNList = append(certSNList, _getCertSN(cert))
		}
	}

	c.alipayRootCertSN = strings.Join(certSNList, "_")
	return nil
}

// 加载支付宝公钥证书sn
func (c *Client) loadAliPayPublicCert(s string) error {
	cert, err := parseCertificate(s)
	if err != nil {
		return err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return fmt.Errorf("支付宝公钥证书类型错误，无法获取到public key")
	}

	c.alipayPubCertSN = _getCertSN(cert)
	c.alipayPublicKeyMap[c.alipayPubCertSN] = key
	return nil
}

// 获取证书的sn
func _getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
