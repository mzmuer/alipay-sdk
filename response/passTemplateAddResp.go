package response

// 卡券模板创建接口响应
type PassTemplateAddResp struct {
	baseResponse
	Success bool   `json:"success"` // 操作成功标识【true：成功；false：失败】
	Result  string `json:"result"`  // passId：券唯一id operation：本次调用的操作类型，ADD errorCode：处理结果码（错误码）errorMsg：处理结果说明（错误说明）
}