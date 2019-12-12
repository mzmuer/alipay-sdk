package response

// 卡券实例更新接口响应
type PassInstanceUpdateResp struct {
	BaseResponse
	Success bool   `json:"success"` // 操作成功标识【true：成功；false：失败】
	Result  string `json:"result"`  // 接口调用返回结果信息
}
