package request

// 单笔转账到支付宝账户接口
type (
	FundTransToaccountReq struct{ BaseRequest }

	FundTransToaccountBizModel struct {
		OutBizNo      string `json:"out_biz_no"`      // 商户转账唯一订单号。发起转账来源方定义的转账单据ID，用于将转账回执通知给来源方。
		PayeeType     string `json:"payee_type"`      // 收款方账户类型。可取值： 1、ALIPAY_USERID：支付宝账号对应的支付宝唯一用户号。以2088开头的16位纯数字组成。 2、ALIPAY_LOGONID：支付宝登录号，支持邮箱和手机号格式。
		PayeeAccount  string `json:"payee_account"`   // 收款方账户。与payee_type配合使用。付款方和收款方不能是同一个账户。
		Amount        string `json:"amount"`          // 转账金额，单位：元。只支持2位小数，小数点前最大支持13位，金额必须大于等于0.1元。
		PayerShowName string `json:"payer_show_name"` // 付款方姓名。显示在收款方的账单详情页。如果该字段不传，则默认显示付款方的支付宝认证姓名或单位名称。
		PayeeRealName string `json:"payee_real_name"` // 收款方真实姓名。如果本参数不为空，则会校验该账户在支付宝登记的实名是否与收款方真实姓名一致。
		Remark        string `json:"remark"`          // 转账备注。当付款方为企业账户，且转账金额达到（大于等于）50000元，remark不能为空。
	}
)

func (FundTransToaccountReq) GetMethod() string {
	return "alipay.fund.trans.toaccount.transfer"
}
