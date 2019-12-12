package request

// 手机网站支付
type (
	TradeWapPayReq struct{ BaseRequest }

	TradeWapPayBizModel struct {
		Body            string `json:"body"`              // 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
		Subject         string `json:"subject"`           // 商品的标题/交易标题/订单标题/订单关键字等。
		OutTradeNo      string `json:"out_trade_no"`      // 商户网站唯一订单号
		TimeoutExpress  string `json:"timeout_express"`   // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
		TotalAmount     string `json:"total_amount"`      // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
		QuitUrl         string `json:"quit_url"`          // 用户付款中途退出返回商户网站的地址
		ProductCode     string `json:"product_code"`      // 销售产品码，商家和支付宝签约的产品码 QUICK_WAP_WAY
		MerchantOrderNo string `json:"merchant_order_no"` // 商户原始订单号
	}
)

func (*TradeWapPayReq) GetMethod() string {
	return "alipay.trade.wap.pay"
}
