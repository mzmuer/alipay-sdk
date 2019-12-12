package response

// 订单数据同步接口
type (
	MerchantOrderSyncResp struct {
		BaseResponse
		OrderId string `json:"order_id"`
	}
)
