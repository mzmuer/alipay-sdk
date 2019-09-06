package response

import "encoding/json"

// TODO: 自定义json解析
type commonResponse struct {
	Code    string `json:"code"`     // 网关返回码
	Msg     string `json:"msg"`      // 网关返回码描述
	SubCode string `json:"sub_code"` // 业务返回码
	SubMsg  string `json:"sub_msg"`  // 业务返回码描述
}

type Response interface {
	GetSubCode() string
	IsSuccess() bool
	GetSign() string
	GetRawParams() string
}

// base
type BaseResp struct {
	Resp struct {
		commonResponse
	} `json:"-"`
	RawResp json.RawMessage `json:"xx"`
	Sign    string          `json:"sign"` // 签名
}

func (r *BaseResp) GetSubCode() string {
	return r.Resp.SubCode
}

func (r *BaseResp) IsSuccess() bool {
	return r.Resp.SubCode == ""
}

func (r *BaseResp) GetSign() string {
	return r.Sign
}

func (r *BaseResp) GetRawParams() string {
	return string(r.RawResp)
}
