package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Pay struct {
	AppId       string
	Charset     string // utf-8
	Version     string // 1.0
	SignType    string // RSA2
	PublicKey   string
	PrivateKey  string
	Signer      signer
	SignChecker signChecker
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
		SignChecker: signChecker{publicKey: publicKey},
		SignType:    SignTypeRSA2,
		Format:      "json",
		EncryptType: EncryptTypeAes,
		PrivateKey:  privateKey,
		Signer:      signer{PrivateKey: privateKey},
		isSandBox:   isSandBox,
	}
}

func (p *Pay) Execute(method, notifyUrl string, bizContent interface{}) (interface{}, error) {
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

	resp, err := doPost(gateway, requestParams)
	if err != nil {
		return nil, err
	}

	ioutil.WriteFile("xxx.html", resp, 0777)

	fmt.Println("====", string(resp))

	return resp, err
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
		signContent := getSignatureContent(params)
		params[Sign], err = p.Signer.Sign(signContent, p.SignType, p.Charset)
		if err != nil {
			return nil, err
		}
	} else {
		params[Sign] = ""
	}

	return params, nil
}

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
