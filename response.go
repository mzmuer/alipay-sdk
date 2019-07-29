package alipay

type (
	commonResponse struct {
		Code    string `json:"code"`     // 网关返回码
		Msg     string `json:"msg"`      // 网关返回码描述
		SubCode string `json:"sub_code"` // 业务返回码
		SubMsg  string `json:"sub_msg"`  // 业务返回码描述
		Sign    string `json:"sign"`     // 签名
	}

	TradeWapPayResp struct {
		commonResponse
		OutTradeNo      string `json:"out_trade_no"`      // 商户网站唯一订单号
		TradeNo         string `json:"trade_no"`          // 该交易在支付宝系统中的交易流水号
		TotalAmount     string `json:"total_amount"`      // 该笔订单的资金总额，单位为RMB-Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位
		SellerId        string `json:"seller_id"`         // 收款支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字
		MerchantOrderNo string `json:"merchant_order_no"` // 商户原始订单号
	}
)
