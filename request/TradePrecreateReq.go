package request

// https://docs.open.alipay.com/api_1/alipay.trade.precreate/
// 统一收单线下交易预创建
type (
	TradePrecreateReq struct{ BaseRequest }

	TradePrecreateBizModel struct {
		OutTradeNo  string `json:"out_trade_no"` // 商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
		TotalAmount string `json:"total_amount"` // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
		Subject     string `json:"subject"`      // 订单标题

		SellerId             string                 `json:"seller_id,omitempty"`           // 	卖家支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
		DiscountableAmount   string                 `json:"discountable_amount,omitempty"` // 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
		GoodsDetail          []TradePrecreateDetail `json:"goods_detail,omitempty"`
		Body                 string                 `json:"body,omitempty"`         // 对交易或商品的描述
		ProductCode          string                 `json:"product_code,omitempty"` // 销售产品码。 如果签约的是当面付快捷版，则传OFFLINE_PAYMENT; 其它支付宝当面付产品传FACE_TO_FACE_PAYMENT； 不传默认使用FACE_TO_FACE_PAYMENT；
		OperatorId           string                 `json:"operator_id,omitempty"`
		StoreId              string                 `json:"store_id,omitempty"`
		DisablePayChannels   string                 `json:"disable_pay_channels,omitempty"`
		EnablePayChannels    string                 `json:"enable_pay_channels,omitempty"`
		TerminalId           string                 `json:"terminal_id,omitempty"`
		TimeoutExpress       string                 `json:"timeout_express,omitempty"`         // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
		MerchantOrderNo      string                 `json:"merchant_order_no,omitempty"`       // 商户原始订单号，最大长度限制32位
		QrCodeTimeoutExpress string                 `json:"qr_code_timeout_express,omitempty"` // 该笔订单允许的最晚付款时间，逾期将关闭交易，从生成二维码开始计时。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	}

	TradePrecreateDetail struct {
		GoodsID        string `json:"goods_id"`
		GoodsName      string `json:"goods_name"`
		Quantity       int    `json:"quantity"`
		Price          int    `json:"price"`
		GoodsCategory  string `json:"goods_category,omitempty"`
		CategoriesTree string `json:"categories_tree,omitempty"`
		Body           string `json:"body,omitempty"`
		ShowURL        string `json:"show_url,omitempty"`
	}
)

func (*TradePrecreateReq) GetMethod() string {
	return "alipay.trade.precreate"
}
