package request

// 统一收单交易退款接口
type (
	TradeRefundReq struct{ BaseRequest }

	TradeRefundBizModel struct {
		OutTradeNo     string `json:"out_trade_no"`    // 订单支付时传入的商户订单号,不能和 trade_no同时为空。
		TradeNo        string `json:"trade_no"`        // 支付宝交易号，和商户订单号不能同时为空
		RefundAmount   string `json:"refund_amount"`   // 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数]
		OutRequestNo   string `json:"out_request_no"`  // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
		RefundCurrency string `json:"refund_currency"` // 订单退款币种信息
		RefundReason   string `json:"refund_reason"`   // 退款的原因说明
	}
)

func (TradeRefundReq) GetMethod() string {
	return "alipay.trade.refund"
}
