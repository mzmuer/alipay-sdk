package request

// 发送模板消息接口
type (
	SendTemplateMsgReq struct{ BaseRequest }

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
