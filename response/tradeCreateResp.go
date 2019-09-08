package response

// 统一收单交易创建接口
type TradeCreateResp struct {
	baseResponse
	OutTradeNo string `json:"out_trade_no"` // 商户网站唯一订单号
	TradeNo    string `json:"trade_no"`     // 该交易在支付宝系统中的交易流水号
}
