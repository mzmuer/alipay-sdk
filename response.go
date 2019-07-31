package alipay

import "encoding/json"

type (
	commonResponse struct {
		Code    string `json:"code"`     // 网关返回码
		Msg     string `json:"msg"`      // 网关返回码描述
		SubCode string `json:"sub_code"` // 业务返回码
		SubMsg  string `json:"sub_msg"`  // 业务返回码描述
	}

	// TODO: 自定义json解析

	// 统一收单交易创建接口
	TradeCreateResp struct {
		Resp struct {
			commonResponse
			OutTradeNo string `json:"out_trade_no"` // 商户网站唯一订单号
			TradeNo    string `json:"trade_no"`     // 该交易在支付宝系统中的交易流水号
		} `json:"-"`
		RawResp json.RawMessage `json:"alipay_trade_create_response"`
		Sign    string          `json:"sign"` // 签名
	}
)

type Response interface {
	IsSuccess() bool
	GetSign() string
	GetRawParams() string
}

func (r *TradeCreateResp) IsSuccess() bool {
	return r.Resp.SubCode == ""
}

func (r *TradeCreateResp) GetSign() string {
	return r.Sign
}

func (r *TradeCreateResp) GetRawParams() string {
	return string(r.RawResp)
}
