package request

// 支付宝会员授权信息查询接口
type (
	UserInfoShareReq struct{ baseRequest }
)

func (*UserInfoShareReq) GetMethod() string {
	return "alipay.user.info.share"
}
