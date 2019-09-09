package response

import (
	"encoding/json"
	"strings"
)

// TODO: 自定义json解析
type baseResponse struct {
	Code         string `json:"code"`     // 网关返回码
	Msg          string `json:"msg"`      // 网关返回码描述
	SubCode      string `json:"sub_code"` // 业务返回码
	SubMsg       string `json:"sub_msg"`  // 业务返回码描述
	AlipayCertSn string `json:"-"`        // 公钥证书sn
	RawResp      string `json:"-"`        // 原始响应内容，签名验证需要
	Sign         string `json:"-"`        // 签名
}

type Response interface {
	GetSubCode() string
	IsSuccess() bool
	GetSign() string
	SetSign(string)
	GetRawParams() string
	SetRawParams(string)
	SetAlipayCertSn(string)
	GetAlipayCertSn() string
}

func (r *baseResponse) GetSubCode() string {
	return r.SubCode
}

func (r *baseResponse) IsSuccess() bool {
	return r.SubCode == ""
}

func (r *baseResponse) GetSign() string {
	return r.Sign
}

func (r *baseResponse) SetSign(sign string) {
	r.Sign = sign
}

func (r *baseResponse) GetRawParams() string {
	return r.RawResp
}

func (r *baseResponse) SetRawParams(raw string) {
	r.RawResp = raw
}

func (r *baseResponse) SetAlipayCertSn(sn string) {
	r.AlipayCertSn = sn
}

func (r *baseResponse) GetAlipayCertSn() string {
	return r.AlipayCertSn
}

// ----
func ParseResponse(method string, data []byte, result Response) error {
	var (
		jsonKey = strings.ReplaceAll(method, ".", "_") + "_response"
		tmpMap  = map[string]json.RawMessage{}
	)

	err := json.Unmarshal(data, &tmpMap)
	if err != nil {
		return err
	}

	jsonRaw := tmpMap[jsonKey]
	raw := strings.Trim(string(jsonRaw), "\"")
	if raw == "" {
		// 老版本失败节点
		jsonRaw = tmpMap["error_response"]
		raw = strings.Trim(string(jsonRaw), "\"")
	}

	result.SetRawParams(raw)
	result.SetSign(strings.Trim(string(tmpMap["sign"]), "\""))
	result.SetAlipayCertSn(strings.Trim(string(tmpMap["alipay_cert_sn"]), "\""))

	return json.Unmarshal(jsonRaw, result)
}
