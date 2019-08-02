package alipay

type (
	// 请求结构
	Request struct {
		Method       string // 请求方法
		NotifyUrl    string
		ReturnUrl    string
		ProdCode     string
		TerminalType string
		TerminalInfo string
		NeedEncrypt  bool
		EncryptType  string
		BizContent   interface{}
	}

	// ------------------- 请求实际业务结构 ---------------------------
	// 统一收单交易创建接口
	TradeCreateReq struct {
		Body           string `json:"body"`            // 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
		Subject        string `json:"subject"`         // 商品的标题/交易标题/订单标题/订单关键字等。
		OutTradeNo     string `json:"out_trade_no"`    //商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
		TimeoutExpress string `json:"timeout_express"` // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
		TotalAmount    string `json:"total_amount"`    // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
		BuyerLogonId   string `json:"buyer_logon_id"`  // 买家支付宝账号，和buyer_id不能同时为空
		BuyerId        string `json:"buyer_id"`        // 买家的支付宝唯一用户号（2088开头的16位纯数字）
	}

	// 手机网站支付
	TradeWapPayReq struct {
		Body            string `json:"body"`              // 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
		Subject         string `json:"subject"`           // 商品的标题/交易标题/订单标题/订单关键字等。
		OutTradeNo      string `json:"out_trade_no"`      // 商户网站唯一订单号
		TimeoutExpress  string `json:"timeout_express"`   // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
		TotalAmount     string `json:"total_amount"`      // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
		QuitUrl         string `json:"quit_url"`          // 用户付款中途退出返回商户网站的地址
		ProductCode     string `json:"product_code"`      // 销售产品码，商家和支付宝签约的产品码 QUICK_WAP_WAY
		MerchantOrderNo string `json:"merchant_order_no"` // 商户原始订单号
	}

	// 统一收单交易退款接口
	TradeRefundReq struct {
		OutTradeNo     string `json:"out_trade_no"`    // 订单支付时传入的商户订单号,不能和 trade_no同时为空。
		TradeNo        string `json:"trade_no"`        // 支付宝交易号，和商户订单号不能同时为空
		RefundAmount   string `json:"refund_amount"`   // 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数]
		OutRequestNo   string `json:"out_request_no"`  // 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
		RefundCurrency string `json:"refund_currency"` // 订单退款币种信息
		RefundReason   string `json:"refund_reason"`   // 退款的原因说明
	}

	// 统一收单交易退款查询
	// alipay.trade.fastpay.refund.query
	TradeRefundQueryReq struct {
		OutTradeNo   string `json:"out_trade_no"`   // 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
		TradeNo      string `json:"trade_no"`       // 支付宝交易号，和商户订单号不能同时为空
		OutRequestNo string `json:"out_request_no"` // 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
	}

	// 单笔转账到支付宝账户接口
	// alipay.fund.trans.toaccount.transfer
	FundTransToaccountReq struct {
		OutBizNo      string `json:"out_biz_no"`      // 商户转账唯一订单号。发起转账来源方定义的转账单据ID，用于将转账回执通知给来源方。
		PayeeType     string `json:"payee_type"`      // 收款方账户类型。可取值： 1、ALIPAY_USERID：支付宝账号对应的支付宝唯一用户号。以2088开头的16位纯数字组成。 2、ALIPAY_LOGONID：支付宝登录号，支持邮箱和手机号格式。
		PayeeAccount  string `json:"payee_account"`   // 收款方账户。与payee_type配合使用。付款方和收款方不能是同一个账户。
		Amount        string `json:"amount"`          // 转账金额，单位：元。只支持2位小数，小数点前最大支持13位，金额必须大于等于0.1元。
		PayerShowName string `json:"payer_show_name"` // 付款方姓名。显示在收款方的账单详情页。如果该字段不传，则默认显示付款方的支付宝认证姓名或单位名称。
		PayeeRealName string `json:"payee_real_name"` // 收款方真实姓名。如果本参数不为空，则会校验该账户在支付宝登记的实名是否与收款方真实姓名一致。
		Remark        string `json:"remark"`          // 转账备注。当付款方为企业账户，且转账金额达到（大于等于）50000元，remark不能为空。
	}

	// 查询转账订单接口
	// alipay.fund.trans.order.query
	FundTransOrderQueryReq struct {
		OutBizNo string `json:"out_biz_no"` // 商户转账唯一订单号：发起转账来源方定义的转账单据ID。和支付宝转账单据号不能同时为空
		OrderId  string `json:"order_id"`   // 支付宝转账单据号：和商户转账唯一订单号不能同时为空
	}
)
