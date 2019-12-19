package request

//https://docs.open.alipay.com/api_1/alipay.trade.create
// 统一收单交易创建接口
type (
	TradeCreateReq struct{ BaseRequest }

	GoodsDetail struct {
		GoodsID        string `json:"goods_id"`
		GoodsName      string `json:"goods_name"`
		Quantity       int    `json:"quantity"`
		Price          int    `json:"price"`
		GoodsCategory  string `json:"goods_category,omitempty"`
		CategoriesTree string `json:"categories_tree,omitempty"`
		Body           string `json:"body,omitempty"`
		ShowURL        string `json:"show_url,omitempty"`
	}

	TradeCreateBizModel struct {
		Body           string        `json:"body"`            // 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
		Subject        string        `json:"subject"`         // 商品的标题/交易标题/订单标题/订单关键字等。
		OutTradeNo     string        `json:"out_trade_no"`    // 商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
		TimeoutExpress string        `json:"timeout_express"` // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
		TotalAmount    string        `json:"total_amount"`    // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
		BuyerLogonId   string        `json:"buyer_logon_id"`  // 买家支付宝账号，和buyer_id不能同时为空
		BuyerId        string        `json:"buyer_id"`        // 买家的支付宝唯一用户号（2088开头的16位纯数字）
		GoodsDetail    []GoodsDetail `json:"goods_detail"`
	}
)

func (*TradeCreateReq) GetMethod() string {
	return "alipay.trade.create"
}
