package response

import (
	"encoding/json"
	"strings"
)

// TODO: 自定义json解析
type baseResponse struct {
	Code    string `json:"code"`     // 网关返回码
	Msg     string `json:"msg"`      // 网关返回码描述
	SubCode string `json:"sub_code"` // 业务返回码
	SubMsg  string `json:"sub_msg"`  // 业务返回码描述
	RawResp string `json:"-"`        // 原始响应内容，签名验证需要
	Sign    string `json:"-"`        // 签名
}

type Response interface {
	GetSubCode() string
	IsSuccess() bool
	GetSign() string
	SetSign(string)
	GetRawParams() string
	SetRawParams(string)
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

// ----
func ParseResponse(method string, data []byte, result Response) error {
	var (
		jsKey  = strings.ReplaceAll(method, ".", "_") + "_response"
		tmpMap = map[string]json.RawMessage{}
	)

	err := json.Unmarshal(data, &tmpMap)
	if err != nil {
		return err
	}

	result.SetRawParams(strings.Trim(string(tmpMap[jsKey]), "\""))
	result.SetSign(strings.Trim(string(tmpMap["sign"]), "\""))

	return json.Unmarshal(tmpMap[jsKey], result)
}
