package response

import (
	"encoding/json"
	"strings"
)

// TODO: 自定义json解析
type BaseResponse struct {
	Code         string `json:"code"`     // 网关返回码
	Msg          string `json:"msg"`      // 网关返回码描述
	SubCode      string `json:"sub_code"` // 业务返回码
	SubMsg       string `json:"sub_msg"`  // 业务返回码描述
	alipayCertSn string `json:"-"`        // 公钥证书sn
	rawResp      string `json:"-"`        // 原始响应内容，签名验证需要
	sign         string `json:"-"`        // 签名
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

func (r *BaseResponse) GetSubCode() string {
	return r.SubCode
}

func (r *BaseResponse) IsSuccess() bool {
	return r.SubCode == ""
}

func (r *BaseResponse) GetSign() string {
	return r.sign
}

func (r *BaseResponse) SetSign(sign string) {
	r.sign = sign
}

func (r *BaseResponse) GetRawParams() string {
	return r.rawResp
}

func (r *BaseResponse) SetRawParams(raw string) {
	r.rawResp = raw
}

func (r *BaseResponse) SetAlipayCertSn(sn string) {
	r.alipayCertSn = sn
}

func (r *BaseResponse) GetAlipayCertSn() string {
	return r.alipayCertSn
}

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
