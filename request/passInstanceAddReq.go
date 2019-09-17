package request

// 卡券实例发放接口
type (
	PassInstanceAddReq struct{ baseRequest }

	PassInstanceAddBizModel struct {
		TplId           string `json:"tpl_id"`           // 支付宝pass模版ID，即调用模板创建接口时返回的tpl_id。
		TplParams       string `json:"tpl_params"`       // 模版动态参数信息：对应模板中$变量名$的动态参数，见模板创建接口返回值中的tpl_params字段
		RecognitionType string `json:"recognition_type"` // Alipass添加对象识别类型：1–订单信息
		RecognitionInfo string `json:"amount"`           // 支付宝用户识别信息：uid发券组件
	}
)

func (*PassInstanceAddReq) GetMethod() string {
	return "alipay.pass.instance.add"
}
