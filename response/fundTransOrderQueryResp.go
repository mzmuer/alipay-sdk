package response

// 查询转账订单接口
type FundTransOrderQueryResp struct {
	baseResponse
	OrderId string `json:"order_id"` // 支付宝转账单据号，查询失败不返回。

	//转账单据状态。
	//SUCCESS：成功（配合"单笔转账到银行账户接口"产品使用时, 同一笔单据多次查询有可能从成功变成退票状态）；
	//FAIL：失败（具体失败原因请参见error_code以及fail_reason返回值）；
	//INIT：等待处理；
	//DEALING：处理中；
	//REFUND：退票（仅配合"单笔转账到银行账户接口"产品使用时会涉及, 具体退票原因请参见fail_reason返回值）；
	//UNKNOWN：状态未知。
	Status         string `json:"status"`
	PayDate        string `json:"pay_date"`         // 支付时间，格式为yyyy-MM-dd HH:mm:ss，转账失败不返回。
	ArrivalTimeEnd string `json:"arrival_time_end"` // 预计到账时间，转账到银行卡专用，格式为yyyy-MM-dd HH:mm:ss，转账受理失败不返回。
	OrderFee       string `json:"order_fee"`        // 预计收费金额（元），转账到银行卡专用，数字格式，精确到小数点后2位，转账失败或转账受理失败不返回。
	FailReason     string `json:"fail_reason"`      // 查询到的订单状态为FAIL失败或REFUND退票时，返回具体的原因。
	OutBizNo       string `json:"out_biz_no"`       // 发起转账来源方定义的转账单据号。该参数的赋值均以查询结果中的out_biz_no为准。如果查询失败，不返回该参数。
	ErrorCode      string `json:"error_code"`       // 查询失败时，本参数为错误代码。 查询成功不返回。 对于退票订单，不返回该参数。
}
