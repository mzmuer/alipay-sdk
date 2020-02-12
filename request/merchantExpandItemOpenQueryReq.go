package request

// 查询商品接口
// https://docs.open.alipay.com/api_4/ant.merchant.expand.item.open.query/
type (
	MerchantExpandItemOpenQueryReq struct {
		BaseRequest
	}

	MerchantExpandItemOpenQueryBizModel struct {
		Scene      string `json:"scene"`
		TargetId   string `json:"target_id"`
		TargetType string `json:"target_type"` //8-小程序
		Status     string `json:"status"`
	}
)

func (MerchantExpandItemOpenQueryReq) GetMethod() string {
	return "ant.merchant.expand.item.open.query"
}
