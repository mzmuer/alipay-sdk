package response

// 统一收单线下交易预创建
type TradePrecreateResp struct {
	BaseResponse
	OutTradeNo string `json:"out_trade_no"` // 商户网站唯一订单号
	QrCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
}
