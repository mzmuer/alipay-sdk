package request

// 支付宝会员授权信息查询接口
type (
	UserInfoShareReq struct{ BaseRequest }
)

func (UserInfoShareReq) GetMethod() string {
	return "alipay.user.info.share"
}
