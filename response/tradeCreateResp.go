package response

import "encoding/json"

// 统一收单交易创建接口
type TradeCreateResp struct {
	Resp struct {
		commonResponse
		OutTradeNo string `json:"out_trade_no"` // 商户网站唯一订单号
		TradeNo    string `json:"trade_no"`     // 该交易在支付宝系统中的交易流水号
	} `json:"-"`
	RawResp json.RawMessage `json:"alipay_trade_create_response"`
	Sign    string          `json:"sign"` // 签名
}

func (r *TradeCreateResp) GetSubCode() string {
	return r.Resp.SubCode
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
