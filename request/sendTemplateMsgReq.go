package request

// 查询转账订单接口
type (
	SendTemplateMsgReq struct{ baseRequest }

	SendTemplateMsgBizModel struct {
		ToUserId       string                       `json:"to_user_id"`
		FormId         string                       `json:"form_id"`
		UserTemplateId string                       `json:"user_template_id"`
		Page           string                       `json:"page"`
		Data           map[string]map[string]string `json:"data"`
	}
)

func (*SendTemplateMsgReq) GetMethod() string {
	return "alipay.open.app.mini.templatemessage.send"
}
