package request

// 卡券模板创建接口接口
type (
	PassTemplateAddReq struct{ BaseRequest }

	PassTemplateAddBizModel struct {
		UniqueId   string `json:"unique_id"`   // 商户用于控制模版的唯一性。（可以使用时间戳保证唯一性）
		TplContent string `json:"tpl_content"` // 模板内容信息，遵循JSON规范，详情参见tpl_content参数说明：https://doc.open.alipay.com/doc2/detail.htm?treeId=193&articleId=105249&docType=1#tpl_content
	}
)

func (PassTemplateAddReq) GetMethod() string {
	return "alipay.pass.template.add"
}
