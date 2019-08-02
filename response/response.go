package response

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
