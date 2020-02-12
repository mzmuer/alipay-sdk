package request

// 卡券模板更新接口
type (
	PassTemplateUpdateReq struct{ BaseRequest }

	PassTemplateUpdateBizModel struct {
		TplId      string `json:"tpl_id"`      // 商户用于控制模版的唯一性。（可以使用时间戳保证唯一性）
		TplContent string `json:"tpl_content"` // 模板内容信息，遵循JSON规范，详情参见tpl_content参数说明：https://doc.open.alipay.com/doc2/detail.htm?treeId=193&articleId=105249&docType=1#tpl_content
	}
)

func (PassTemplateUpdateReq) GetMethod() string {
	return "alipay.pass.template.update"
}
