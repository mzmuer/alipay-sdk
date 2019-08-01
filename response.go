package alipay

import (
	"encoding/json"
)

type commonResponse struct {
	Code    string `json:"code"`     // 网关返回码
	Msg     string `json:"msg"`      // 网关返回码描述
	SubCode string `json:"sub_code"` // 业务返回码
	SubMsg  string `json:"sub_msg"`  // 业务返回码描述
}

// TODO: 自定义json解析

type Response interface {
	GetSubCode() string
	IsSuccess() bool
	GetSign() string
	GetRawParams() string
}

// 统一收单交易创建接口
type TradeCreateResp struct {
	Resp struct {
		commonResponse
		OutTradeNo string `json:"out_trade_no"` // 商户网站唯一订单号
		TradeNo    string `json:"trade_no"`     // 该交易在支付宝系统中的交易流水号
	} `json:"-"`
	RawResp json.RawMessage `json:"alipay_trade_create_response"`
	Sign    string          `json:"sign"` // 签名
}

func (r *TradeCreateResp) GetSubCode() string {
	return r.Resp.SubCode
}

func (r *TradeCreateResp) IsSuccess() bool {
	return r.Resp.SubCode == ""
}

func (r *TradeCreateResp) GetSign() string {
	return r.Sign
}

func (r *TradeCreateResp) GetRawParams() string {
	return string(r.RawResp)
}

// 统一收单交易退款接口响应
type TradeRefundResp struct {
	Resp struct {
		commonResponse
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
	} `json:"-"`
	RawResp json.RawMessage `json:"alipay_trade_refund_response"`
	Sign    string          `json:"sign"` // 签名
}

func (r *TradeRefundResp) GetSubCode() string {
	return r.Resp.SubCode
}

func (r *TradeRefundResp) IsSuccess() bool {
	return r.Resp.SubCode == ""
}

func (r *TradeRefundResp) GetSign() string {
	return r.Sign
}

func (r *TradeRefundResp) GetRawParams() string {
	return string(r.RawResp)
}

// 统一收单交易退款查询响应
type TradeRefundQueryResp struct {
	Resp struct {
		commonResponse
		// optional
		OutTradeNo                   string          `json:"out_trade_no"`                    // 商户网站唯一订单号
		TradeNo                      string          `json:"trade_no"`                        // 该交易在支付宝系统中的交易流水号
		OutRequestNo                 string          `json:"out_request_no"`                  // 本笔退款对应的退款请求号
		RefundReason                 string          `json:"refund_reason"`                   // 发起退款时，传入的退款原因
		TotalAmount                  string          `json:"total_amount"`                    // 该笔退款所对应的交易的订单金额
		RefundAmount                 string          `json:"refund_amount"`                   // 本次退款请求，对应的退款金额
		GmtRefundPay                 string          `json:"gmt_refund_pay"`                  // 退款时间；默认不返回该信息，需与支付宝约定后配置返回；
		RefundStatus                 string          `json:"refund_status"`                   // 为空或为REFUND_SUCCESS，则代表退款成功
		RefundDetailItemList         json.RawMessage `json:"refund_detail_item_list"`         // 退款使用的资金渠道
		SendBackFee                  string          `json:"send_back_fee"`                   // 本次商户实际退回金额；默认不返回该信息，需与支付宝约定后配置返回；
		RefundSettlementId           string          `json:"refund_settlement_id"`            // 退款清算编号，用于清算对账使用；只在银行间联交易场景下返回该信息
		PresentRefundBuyerAmount     string          `json:"present_refund_buyer_amount"`     // 本次退款金额中买家退款金额
		PresentRefundDiscountAmount  string          `json:"present_refund_discount_amount"`  // 本次退款金额中平台优惠退款金额
		PresentRefundMdiscountAmount string          `json:"present_refund_mdiscount_amount"` // 本次退款金额中商家优惠退款金额
	} `json:"-"`
	RawResp json.RawMessage `json:"alipay_trade_fastpay_refund_query_response"`
	Sign    string          `json:"sign"` // 签名
}

func (r *TradeRefundQueryResp) GetSubCode() string {
	return r.Resp.SubCode
}

func (r *TradeRefundQueryResp) IsSuccess() bool {
	return r.Resp.SubCode == ""
}

func (r *TradeRefundQueryResp) GetSign() string {
	return r.Sign
}

func (r *TradeRefundQueryResp) GetRawParams() string {
	return string(r.RawResp)
}
