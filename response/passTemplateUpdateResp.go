package response

// 卡券模板创建接口响应
type PassTemplateUpdateResp struct {
	baseResponse
	Success bool   `json:"success"` // 操作成功标识【true：成功；false：失败】
	Result  string `json:"result"`  // 接口调用返回结果信息(json字串): errorCode：处理结果码（错误码）； errorMsg：处理结果说明（错误说明）； tpl_id：模板编号，预期在调发券接口时必须传入； tpl_params：动态参数（变量）列表，预期在调发券接口时传入；
}
