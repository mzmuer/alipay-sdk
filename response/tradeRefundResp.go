package response

import "encoding/json"

// 统一收单交易退款接口响应
type TradeRefundResp struct {
	baseResponse
	// must
	OutTradeNo   string `json:"out_trade_no"`   // 商户网站唯一订单号
	TradeNo      string `json:"trade_no"`       // 该交易在支付宝系统中的交易流水号
	BuyerLogonId string `json:"buyer_logon_id"` // 用户的登录id
	FundChange   string `json:"fund_change"`    // 本次退款是否发生了资金变化
	RefundFee    string `json:"refund_fee"`     // 退款总金额
	GmtRefundPay string `json:"gmt_refund_pay"` // 退款支付时间
	BuyerUserId  string `json:"buyer_user_id"`  // 买家在支付宝的用户id

	// optional
	RefundCurrency               string          `json:"refund_currency"`                 // 退款币种信息
	StoreName                    string          `json:"store_name"`                      // 交易在支付时候的门店名称
	RefundDetailItemList         json.RawMessage `json:"refund_detail_item_list"`         // 退款使用的资金渠道
	RefundPresetPaytoolList      json.RawMessage `json:"refund_preset_paytool_list"`      // 退回的前置资产列表
	RefundSettlementId           string          `json:"refund_settlement_id"`            // 退款清算编号，用于清算对账使用；只在银行间联交易场景下返回该信息
	PresentRefundBuyerAmount     string          `json:"present_refund_buyer_amount"`     // 本次退款金额中买家退款金额
	PresentRefundDiscountAmount  string          `json:"present_refund_discount_amount"`  // 本次退款金额中平台优惠退款金额
	PresentRefundMdiscountAmount string          `json:"present_refund_mdiscount_amount"` // 本次退款金额中商家优惠退款金额
}
