package request

//https://docs.open.alipay.com/api_1/alipay.trade.app.pay/
// 统一收单交易创建接口
type (
	TradeAppPayReq struct{ BaseRequest }

	ExtUserInfo struct {
		Name          string `json:"name"`
		Mobile        string `json:"mobile"`
		CertType      string `json:"cert_type"`
		CertNo        string `json:"cert_no"`
		MinAge        string `json:"min_age"`
		FixBuyer      string `json:"fix_buyer"`
		NeedCheckInfo string `json:"need_check_info"`
	}

	SignParams struct {
		PersonalProductCode string             `json:"personal_product_code"`
		SignScene           string             `json:"sign_scene"`
		ExternalAgreementNo string             `json:"external_agreement_no"`
		ExternalLogonId     string             `json:"external_logon_id"`
		AccessParams        AccessParams       `json:"access_params"`
		SubMerchant         SignMerchantParams `json:"sub_merchant"`
		PeriodRuleParams    PeriodRuleParams   `json:"period_rule_params"`
	}

	SignMerchantParams struct {
		MerchantID   string `json:"merchant_id"`
		MerchantType string `json:"merchant_type"`
	}

	AccessParams struct {
		Channel string `json:"channel"`
	}

	ExtendParams struct {
		SysServiceProviderID string `json:"sys_service_provider_id"`
		HbFqNum              string `json:"hb_fq_num"`
		HbFqSellerPercent    string `json:"hb_fq_seller_percent"`
		IndustryRefluxInfo   string `json:"industry_reflux_info"`
		CardType             string `json:"card_type"`
	}

	TradeAppPayGoodsDetail struct {
		GoodsID        string `json:"goods_id"`
		AlipayGoodsID  string `json:"alipay_goods_id"`
		GoodsName      string `json:"goods_name"`
		Quantity       int    `json:"quantity"`
		Price          int    `json:"price"`
		GoodsCategory  string `json:"goods_category"`
		CategoriesTree string `json:"categories_tree"`
		Body           string `json:"body"`
		ShowURL        string `json:"show_url"`
	}

	PeriodRuleParams struct {
		PeriodType    string `json:"period_type"`
		Period        int32  `json:"period"`
		ExecuteTime   string `json:"execute_time"`
		SingleAmount  int32  `json:"single_amount"`
		TotalAmount   int32  `json:"total_amount"`
		TotalPayments int32  `json:"total_payments"`
	}

	TradeAppPayBizModel struct {
		TimeoutExpress      string                 `json:"timeout_express"`
		TotalAmount         float64                `json:"total_amount"`
		ProductCode         string                 `json:"product_code"`
		Body                string                 `json:"body"`
		Subject             string                 `json:"subject"`
		OutTradeNo          string                 `json:"out_trade_no"`
		TimeExpire          string                 `json:"time_expire"`
		GoodsType           string                 `json:"goods_type"`
		PromoParams         string                 `json:"promo_params"`
		PassbackParams      string                 `json:"passback_params"`
		ExtendParams        ExtendParams           `json:"extend_params"`
		MerchantOrderNo     string                 `json:"merchant_order_no"`
		EnablePayChannels   string                 `json:"enable_pay_channels"`
		StoreID             string                 `json:"store_id"`
		SpecifiedChannel    string                 `json:"specified_channel"`
		DisablePayChannels  string                 `json:"disable_pay_channels"`
		GoodsDetail         TradeAppPayGoodsDetail `json:"goods_detail"`
		ExtUserInfo         ExtUserInfo            `json:"ext_user_info"`
		BusinessParams      string                 `json:"business_params"`
		AgreementSignParams SignParams             `json:"agreement_sign_params"`
	}
)

func (*TradeAppPayReq) GetMethod() string {
	return "alipay.trade.app.pay"
}
