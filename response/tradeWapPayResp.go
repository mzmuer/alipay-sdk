package response

// 手机网站支付
type (
	TradeWapPayResp struct {
		BaseResponse
		OutTradeNo      string `json:"out_trade_no"`      // 商户网站唯一订单号
		TradeNo         string `json:"trade_no"`          // 该交易在支付宝系统中的交易流水号。最长64位。
		TotalAmount     int64  `json:"total_amount"`      // 该笔订单的资金总额，单位为RMB-Yuan。取值范围为[0.01，100000000.00]。
		SellerId        string `json:"seller_id"`         // 收款支付宝账号对应的支付宝唯一用户号。 以2088开头的纯16位数字
		MerchantOrderNo string `json:"merchant_order_no"` // 商户原始订单号，最大长度限制32位
	}
)
