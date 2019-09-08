package alipay

import (
	"crypto/md5"
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
	Version     string // 1.0
	SignType    string // RSA2
	Signer      *signer
	SignChecker *signChecker
	Format      string // json
	EncryptType string // AES

	// 公钥证书相关参数
	appPubCertSN     string
	alipayRootCertSN string
	alipayPubCertSN  string

	isSandBox bool
}

func NewClient(appId string, publicKey, privateKey []byte, isSandBox bool) (*Client, error) {
	signChecker, err := NewSignChecker(publicKey)
	if err != nil {
		return nil, err
	}

	signer, err := NewSigner(privateKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		AppId:       appId,
		Charset:     "utf-8",
		Version:     "1.0",
		SignChecker: signChecker,
		SignType:    SignTypeRSA2,
		Format:      "json",
		EncryptType: EncryptTypeAes,
		Signer:      signer,
		isSandBox:   isSandBox,
	}, nil
}

// 加载应用公钥证书
func (c *Client) LoadAppPublicCert(b []byte) error {
	block, _ := pem.Decode(b)
	if block == nil {
		return fmt.Errorf("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse certificate: " + err.Error())
	}

	c.appPubCertSN = _getCertSN(cert)
	return nil
}

// 加载支付宝根证书
func (c *Client) LoadAliPayRootCert(b []byte) error {
	var certStrList = strings.SplitAfter(string(b), "-----END CERTIFICATE-----")

	var certSNList = make([]string, 0, len(certStrList))
	for _, v := range certStrList {
		block, _ := pem.Decode([]byte(v))
		if block == nil {
			return fmt.Errorf("failed to parse certificate PEM")
		}
		cert, _ := x509.ParseCertificate(block.Bytes)
		//if err != nil {
		//	return fmt.Errorf("failed to parse certificate: " + err.Error())
		//}

		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			certSNList = append(certSNList, _getCertSN(cert))
		}
	}

	c.alipayRootCertSN = strings.Join(certSNList, "_")
	return nil
}

// 加载支付宝公钥证书
func (c *Client) LoadAliPayPublicCert(b []byte) error {
	block, _ := pem.Decode(b)
	if block == nil {
		return fmt.Errorf("failed to parse certificate PEM")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse certificate: " + err.Error())
	}

	//key, ok := cert.PublicKey.(*rsa.PublicKey)
	//if ok == false {
	//	return nil
	//}

	c.alipayPubCertSN = _getCertSN(cert)
	return nil
}

func (p *Client) Execute(method, notifyUrl string, bizContent interface{}, params map[string]string, result response.Response) error {
	r := request.Request{
		Method:     method,
		NotifyUrl:  notifyUrl,
		BizContent: bizContent,
		Params:     params,
	}

	// 签名请求
	requestParams, err := p.getRequestHolderWithSign(&r, "", "")
	if err != nil {
		return err
	}

	gateway := Gateway
	if p.isSandBox {
		gateway = SandboxGateway
	}

	b, err := doPost(gateway, requestParams)
	if err != nil {
		return err
	}

	if err = response.ParseResponse(method, b, result); err != nil {
		return err
	}

	if result.IsSuccess() && result.GetSign() != "" {
		match, err := p.checkResponseSign(result.GetRawParams(), result.GetSign())
		if err != nil {
			return err
		}

		if !match { // 签名不匹配
			return fmt.Errorf("sign check fail: check Sign and Data Fail")
		}
	}

	return nil
}

//
//func _parseResponse(method string, data []byte) (response.Response, error) {
//	var (
//		jsKey  = strings.ReplaceAll(method, ".", "_") + "_response"
//		tmpMap = map[string]json.RawMessage{}
//	)
//
//	err := json.Unmarshal(data, &tmpMap)
//	if err != nil {
//		return nil, err
//	}
//
//	switch method {
//	case AlipayTradeCreate:
//		resp := response.TradeCreateResp{}
//		resp.RawResp = string(tmpMap[jsKey])
//		resp.Sign = string(tmpMap["sign"])
//
//		err := json.Unmarshal(tmpMap[jsKey], &resp)
//		if err != nil {
//			return nil, err
//		}
//
//		return &resp, err
//	case AlipayTradeRefund:
//		resp := response.TradeRefundResp{}
//		err := json.Unmarshal(data, &resp)
//		if err != nil {
//			return nil, err
//		}
//
//		// 解析到结构
//		err = json.Unmarshal(resp.RawResp, &resp.Resp)
//		return &resp, err
//	case AlipayTradeFastpayRefundQuery:
//		resp := response.TradeRefundQueryResp{}
//		err := json.Unmarshal(data, &resp)
//		if err != nil {
//			return nil, err
//		}
//
//		// 解析到结构
//		err = json.Unmarshal(resp.RawResp, &resp.Resp)
//		return &resp, err
//	case AlipayFundTransToaccountTransfer:
//		resp := response.FundTransToaccountResp{}
//		err := json.Unmarshal(data, &resp)
//		if err != nil {
//			return nil, err
//		}
//
//		// 解析到结构
//		err = json.Unmarshal(resp.RawResp, &resp.Resp)
//		return &resp, err
//	case AlipayFundTransOrderQuery:
//		resp := response.FundTransOrderQueryResp{}
//		err := json.Unmarshal(data, &resp)
//		if err != nil {
//			return nil, err
//		}
//
//		// 解析到结构
//		err = json.Unmarshal(resp.RawResp, &resp.Resp)
//		return &resp, err
//	default:
//		return nil, fmt.Errorf("暂不支持的未知的请求方法[%s]", method)
//	}
//}

func _getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}

func (c *Client) getRequestHolderWithSign(r *request.Request, accessToken, appAuthToken string) (map[string]string, error) {
	params := map[string]string{}

	// 必选参数
	params[Method] = r.Method
	params[Version] = c.Version
	params[AppId] = c.AppId
	params[SignType] = c.SignType
	params[TerminalType] = r.TerminalType
	params[TerminalInfo] = r.TerminalInfo
	params[NotifyUrl] = r.NotifyUrl
	params[ReturnUrl] = r.ReturnUrl
	params[Charset] = c.Charset
	params[Timestamp] = time.Now().Format("2006-01-02 15:03:04")
	if r.NeedEncrypt {
		params[EncryptType] = r.EncryptType
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
	params[ProdCode] = r.ProdCode

	// app参数
	if r.BizContent != nil {
		bizContent, err := json.Marshal(r.BizContent)
		if err != nil {
			return nil, err
		}

		params[BizContentKey] = string(bizContent)
	}

	if r.NeedEncrypt {
		if r.EncryptType == "" {
			return nil, fmt.Errorf("加密类型错误")
		}

		params[EncryptType] = r.EncryptType
		// TODO: 对r.BizContent一波加密操作
		// params[BizContentKey] = encryptContent
	}

	if appAuthToken != "" {
		params[AppAuthToken] = appAuthToken
	}

	// 额外参数
	for key, v := range r.Params {
		params[key] = v
	}

	// 签名 - 必选参数
	if c.SignType != "" {
		var err error
		signContent := GetSignatureContent(params)
		params[Sign], err = c.Signer.Sign(signContent, c.SignType, c.Charset)
		if err != nil {
			return nil, err
		}
	} else {
		params[Sign] = ""
	}

	return params, nil
}

// --
func (p *Client) checkResponseSign(sourceContent string, signature string) (bool, error) {
	if p.SignChecker == nil {
		return true, nil
	}

	return p.SignChecker.Check(sourceContent, signature, p.SignType, p.Charset)
}
