package alipay

import (
	"crypto/rsa"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/mzmuer/alipay-sdk/constant"
	"github.com/mzmuer/alipay-sdk/request"
	"github.com/mzmuer/alipay-sdk/response"
	"github.com/mzmuer/alipay-sdk/signature"
	"github.com/mzmuer/alipay-sdk/utils"
	"github.com/tjfoc/gmsm/sm2"
)

type Client struct {
	AppId       string
	Charset     string // utf-8
	SignType    string // RSA2
	Signer      *signature.Signer
	SignChecker *signature.SignChecker
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
	signChecker, err := signature.NewSignChecker([]byte(publicKey))
	if err != nil {
		return nil, err
	}

	signer, err := signature.NewSigner([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	c := Client{
		AppId:       appId,
		Charset:     "utf-8",
		SignChecker: signChecker,
		SignType:    constant.SignTypeRSA2,
		Format:      "json",
		EncryptType: constant.EncryptTypeAes,
		Signer:      signer,
		isSandBox:   isSandBox,
	}

	return &c, nil
}

func NewCertClient(appId, privateKey, appPubCert, alipayRootCert, alipayPubCert string, isSandBox bool) (*Client, error) {
	signer, err := signature.NewSigner([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	c := Client{
		AppId:              appId,
		Charset:            "utf-8",
		SignType:           constant.SignTypeRSA2,
		Format:             "json",
		EncryptType:        constant.EncryptTypeAes,
		Signer:             signer,
		isSandBox:          isSandBox,
		alipayPublicKeyMap: make(map[string]*rsa.PublicKey),
	}

	if err = c.loadAppPubCertSN(appPubCert); err != nil {
		return nil, err
	}

	if err = c.loadAliPayRootCert(alipayRootCert); err != nil {
		return nil, err
	}

	if err = c.loadAliPayPublicCert(alipayPubCert); err != nil {
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

	gateway := constant.Gateway
	if c.isSandBox {
		gateway = constant.SandboxGateway
	}

	var b []byte
	if uRequest, ok := r.(request.UploadRequest); ok {
		if b, err = utils.DoPostUploadFile(gateway, requestParams, uRequest.GetFileParams()); err != nil {
			return "", err
		}
	} else {
		if b, err = utils.DoPost(gateway, requestParams); err != nil {
			return "", err
		}
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

	sign := params[constant.Sign]
	delete(params, constant.Sign)
	delete(params, constant.SignType)

	return c.SignChecker.Check(utils.GetSignatureContent(params), sign, signType, charset)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func (c *Client) RsaCheckV2(params map[string]string, charset, signType string) (bool, error) {
	if c.SignChecker == nil {
		return true, nil
	}

	sign := params[constant.Sign]
	delete(params, constant.Sign)

	return c.SignChecker.Check(utils.GetSignatureContent(params), sign, signType, charset)
}

// 此方法会去掉sign_type做验签，暂时除生活号（原服务窗）激活开发者模式外都使用V1
func (c *Client) RsaCertCheckV1(params map[string]string, charset, signType string) (bool, error) {
	sn := params["alipay_cert_sn"]

	k, ok := c.alipayPublicKeyMap[sn]
	if !ok && params["sub_code"] == "" {
		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
	}

	sign := params[constant.Sign]
	delete(params, constant.Sign)
	delete(params, constant.SignType)

	return signature.NewSignCheckerWithPublicKey(k).Check(utils.GetSignatureContent(params), sign, signType, charset)
}

// 此方法不会去掉sign_type验签，用于生活号（原服务窗）激活开发者模式
func (c *Client) RsaCertCheckV2(params map[string]string, charset, signType string) (bool, error) {
	sn := params["alipay_cert_sn"]

	k, ok := c.alipayPublicKeyMap[sn]
	if !ok && params["sub_code"] == "" {
		return false, fmt.Errorf("cert check fail: ALIPAY_CERT_SN is Empty")
	}

	sign := params[constant.Sign]
	delete(params, constant.Sign)

	return signature.NewSignCheckerWithPublicKey(k).Check(utils.GetSignatureContent(params), sign, signType, charset)
}

// 构造请求map
func (c *Client) getRequestHolderWithSign(r request.Request, accessToken, appAuthToken string) (map[string]string, error) {
	params := map[string]string{}

	// 必选参数
	params[constant.Method] = r.GetMethod()
	params[constant.Version] = r.GetApiVersion()
	params[constant.AppId] = c.AppId
	params[constant.SignType] = c.SignType
	params[constant.TerminalType] = r.GetTerminalType()
	params[constant.TerminalInfo] = r.GetTerminalInfo()
	params[constant.NotifyUrl] = r.GetNotifyUrl()
	params[constant.ReturnUrl] = r.GetReturnUrl()
	params[constant.Charset] = c.Charset
	params[constant.Timestamp] = time.Now().Format("2006-01-02 15:03:04")
	if r.GetNeedEncrypt() {
		params[constant.EncryptType] = c.EncryptType
	}

	if c.appPubCertSN != "" {
		params[constant.AppCertSn] = c.appPubCertSN
	}

	if c.alipayRootCertSN != "" {
		params[constant.AlipayRootCertSn] = c.alipayRootCertSN
	}

	// 可选参数
	params[constant.Format] = c.Format
	params[constant.AccessToken] = accessToken
	params[constant.AlipaySdk] = constant.SdkVersion
	params[constant.ProdCode] = r.GetProdCode()

	// app参数
	// 必须先添加额外参数
	for key, v := range r.GetTextParams() {
		params[key] = v
	}

	if params[constant.BizContentKey] == "" && r.GetBizModel() != nil {
		bizContent, err := json.Marshal(r.GetBizModel())
		if err != nil {
			return nil, err
		}

		params[constant.BizContentKey] = string(bizContent)
	}

	if r.GetNeedEncrypt() {
		if c.EncryptType == "" {
			return nil, fmt.Errorf("加密类型错误")
		}

		params[constant.EncryptType] = c.EncryptType
		// TODO: 对r.BizContent一波加密操作
		// params[BizContentKey] = encryptContent
	}

	if appAuthToken != "" {
		params[constant.AppAuthToken] = appAuthToken
	}

	// 签名 - 必选参数
	if c.SignType != "" {
		var err error
		signContent := utils.GetSignatureContent(params)
		params[constant.Sign], err = c.Signer.Sign(signContent, c.SignType, c.Charset)
		if err != nil {
			return nil, err
		}
	} else {
		params[constant.Sign] = ""
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

	return signature.NewSignCheckerWithPublicKey(k).Check(resp.GetRawParams(), resp.GetSign(), c.SignType, c.Charset)
}

// 加载应用公钥证书sn
func (c *Client) loadAppPubCertSN(s string) error {
	cert, err := utils.ParseCertificate(s)
	if err != nil {
		return err
	}

	c.appPubCertSN = utils.GetCertSN(cert)
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

		// TODO: 自己实现SM2国密
		cert, err := sm2.ParseCertificate(block.Bytes)
		if err != nil {
			return fmt.Errorf("failed to parse certificate: " + err.Error())
		}

		if cert != nil && (cert.SignatureAlgorithm == sm2.SHA256WithRSA || cert.SignatureAlgorithm == sm2.SHA1WithRSA) {
			certSNList = append(certSNList, utils.GetCertSN(cert))
		}
	}

	c.alipayRootCertSN = strings.Join(certSNList, "_")
	return nil
}

// 加载支付宝公钥证书sn
func (c *Client) loadAliPayPublicCert(s string) error {
	cert, err := utils.ParseCertificate(s)
	if err != nil {
		return err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return fmt.Errorf("支付宝公钥证书类型错误，无法获取到public key")
	}

	c.alipayPubCertSN = utils.GetCertSN(cert)
	c.alipayPublicKeyMap[c.alipayPubCertSN] = key
	return nil
}
