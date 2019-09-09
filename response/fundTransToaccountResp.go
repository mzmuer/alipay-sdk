package response

// 单笔转账到支付宝账户接口响应
type FundTransToaccountResp struct {
	baseResponse
	OutBizNo string `json:"out_biz_no"` // 商户转账唯一订单号：发起转账来源方定义的转账单据号。请求时对应的参数，原样返回。
	OrderId  string `json:"order_id"`   // 支付宝转账单据号，成功一定返回，失败可能不返回也可能返回。
	PayDate  string `json:"pay_date"`   // 支付时间：格式为yyyy-MM-dd HH:mm:ss，仅转账成功返回。
}
