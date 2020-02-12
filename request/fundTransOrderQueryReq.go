package request

// 查询转账订单接口
type (
	FundTransOrderQueryReq struct{ BaseRequest }

	FundTransOrderQueryBizModel struct {
		OutBizNo string `json:"out_biz_no"` // 商户转账唯一订单号：发起转账来源方定义的转账单据ID。和支付宝转账单据号不能同时为空
		OrderId  string `json:"order_id"`   // 支付宝转账单据号：和商户转账唯一订单号不能同时为空
	}
)

func (FundTransOrderQueryReq) GetMethod() string {
	return "alipay.fund.trans.order.query"
}
