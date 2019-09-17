package request

// 卡券实例发放接口
type (
	PassInstanceUpdateReq struct{ baseRequest }

	PassInstanceUpdateBizModel struct {
		SerialNumber string `json:"serial_number"` // 商户指定卡券唯一值，卡券JSON模板中fileInfo->serialNumber字段对应的值
		ChannelId    string `json:"channel_id"`    // 代理商代替商户发放卡券后，再代替商户更新卡券时，此值为商户的pid/appid
		TplParams    string `json:"tpl_params"`    // 模版动态参数信息：对应模板中$变量名$的动态参数，见模板创建接口返回值中的tpl_params字段
		Status       string `json:"status"`        // 券状态，支持更新为USED、CLOSED两种状态
		VerifyCode   string `json:"verify_code"`   // 核销码串值【当状态变更为USED时，建议传】。该值正常为模板中核销区域（Operation）对应的message值。
		VerifyType   string `json:"verify_type"`   // 核销方式，该值正常为模板中核销区域（Operation）对应的format值。verify_code和verify_type需同时传入。
	}
)

func (*PassInstanceUpdateReq) GetMethod() string {
	return "alipay.pass.instance.update"
}
