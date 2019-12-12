package response

// 创建商品接口
type (
	MerchantExpandItemOpenCreateResp struct {
		BaseResponse
		ItemId string `json:"item_id"`
	}
)
