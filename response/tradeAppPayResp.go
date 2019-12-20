package response

type TradeAppPayResp struct {
	BaseResponse
	OutTradeNo      string `json:"out_trade_no"` // 商户网站唯一订单号
	TradeNo         string `json:"trade_no"`     // 该交易在支付宝系统中的交易流水号
	TotalAmount     string `json:"total_amount"`
	SellerID        string `json:"seller_id"`
	MerchantOrderNo string `json:"merchant_order_no"`
}
