package alipay

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mzmuer/alipay-sdk/response"
)

type Pay struct {
	AppId       string
	Charset     string // utf-8
	Version     string // 1.0
	SignType    string // RSA2
	PublicKey   string
	PrivateKey  string
	Signer      *signer
	SignChecker *signChecker
	Format      string // json
	EncryptType string // AES
	isSandBox   bool
}

func NewPay(appId, publicKey, privateKey string, isSandBox bool) *Pay {
	return &Pay{
		AppId:       appId,
		Charset:     "utf-8",
		Version:     "1.0",
		PublicKey:   publicKey,
		SignChecker: &signChecker{publicKey: publicKey},
		SignType:    SignTypeRSA2,
		Format:      "json",
		EncryptType: EncryptTypeAes,
		PrivateKey:  privateKey,
		Signer:      &signer{PrivateKey: privateKey},
		isSandBox:   isSandBox,
	}
}

func (p *Pay) Execute(method, notifyUrl string, bizContent interface{}) (response.Response, error) {
	r := Request{
		Method:     method,
		NotifyUrl:  notifyUrl,
		BizContent: bizContent,
	}

	requestParams, err := p.getRequestHolderWithSign(&r, "", "")
	if err != nil {
		return nil, err
	}

	gateway := Gateway
	if p.isSandBox {
		gateway = SandboxGateway
	}

	b, err := doPost(gateway, requestParams)
	if err != nil {
		return nil, err
	}

	resp, err := _parseResponse(bizContent, b)
	if err != nil {
		return nil, err
	}

	if resp.IsSuccess() ||
		(!resp.IsSuccess() && resp.GetSign() != "") {
		match, err := p.checkResponseSign(resp.GetRawParams(), resp.GetSign())
		if err != nil {
			return nil, err
		}

		if !match { // 签名不匹配
			return nil, fmt.Errorf("sign check fail: check Sign and Data Fail")
		}
	}

	return resp, nil
}

func _parseResponse(anchoring interface{}, data []byte) (response.Response, error) {
	switch anchoring.(type) {
	case TradeCreateReq:
		resp := response.TradeCreateResp{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return nil, err
		}

		// 解析到结构
		err = json.Unmarshal(resp.RawResp, &resp.Resp)
		return &resp, err
	case TradeRefundReq:
		resp := response.TradeRefundResp{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return nil, err
		}

		// 解析到结构
		err = json.Unmarshal(resp.RawResp, &resp.Resp)
		return &resp, err
	case TradeRefundQueryReq:
		resp := response.TradeRefundQueryResp{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return nil, err
		}

		// 解析到结构
		err = json.Unmarshal(resp.RawResp, &resp.Resp)
		return &resp, err
	case FundTransToaccountReq:
		resp := response.FundTransToaccountResp{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return nil, err
		}

		// 解析到结构
		err = json.Unmarshal(resp.RawResp, &resp.Resp)
		return &resp, err
	case FundTransOrderQueryReq:
		resp := response.FundTransOrderQueryResp{}
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return nil, err
		}

		// 解析到结构
		err = json.Unmarshal(resp.RawResp, &resp.Resp)
		return &resp, err
	default:
		return nil, fmt.Errorf("未知的请求类型")
	}
}

func (p *Pay) getRequestHolderWithSign(r *Request, accessToken, appAuthToken string) (map[string]string, error) {
	params := map[string]string{}

	// 必选参数
	params[Method] = r.Method
	params[Version] = p.Version
	params[AppId] = p.AppId
	params[SignType] = p.SignType
	params[TerminalType] = r.TerminalType
	params[TerminalInfo] = r.TerminalInfo
	params[NotifyUrl] = r.NotifyUrl
	params[ReturnUrl] = r.ReturnUrl
	params[Charset] = p.Charset
	params[Timestamp] = time.Now().Format("2006-01-02 15:03:04")
	if r.NeedEncrypt {
		params[EncryptType] = r.EncryptType
	}

	// 可选参数
	params[Format] = p.Format
	params[AccessToken] = accessToken
	params[AlipaySdk] = SdkVersion
	params[ProdCode] = r.ProdCode

	// app参数
	bizContent, err := json.Marshal(r.BizContent)
	if err != nil {
		return nil, err
	}

	params[BizContentKey] = string(bizContent)

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

	// 签名 - 必选参数
	if p.SignType != "" {
		signContent := GetSignatureContent(params)
		params[Sign], err = p.Signer.Sign(signContent, p.SignType, p.Charset)
		if err != nil {
			return nil, err
		}
	} else {
		params[Sign] = ""
	}

	return params, nil
}

// --
func (p *Pay) checkResponseSign(sourceContent string, signature string) (bool, error) {
	if p.SignChecker == nil {
		return true, nil
	}

	return p.SignChecker.Check(sourceContent, signature, p.SignType, p.Charset)
}
