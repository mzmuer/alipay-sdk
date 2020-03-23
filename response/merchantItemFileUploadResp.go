package response

// 商品文件上传接口
type MerchantItemFileUploadResp struct {
	BaseResponse
	MaterialId  string `json:"material_id"`  // 文件在商品中心的素材标识
	MaterialKey string `json:"material_key"` // 文件在商品中心的素材标示，创建/更新商品时使用
}
