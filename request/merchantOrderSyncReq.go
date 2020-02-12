package request

// 订单数据同步接口
//https://docs.open.alipay.com/api_4/alipay.merchant.order.sync
type (
	MerchantOrderSyncReq struct {
		BaseRequest
	}

	OrderExtInfo struct {
		ExtKey   string `json:"ext_key"`
		ExtValue string `json:"ext_value"`
	}
	ItemOrderInfo struct {
		SkuId     string         `json:"sku_id"`
		ItemId    string         `json:"item_id"`
		ItemName  string         `json:"item_name"`
		UnitPrice int32          `json:"unit_price"`
		Quantity  int32          `json:"quantity"`
		ExtInfo   []OrderExtInfo `json:"ext_info"`
	}

	OrderLogisticsInformationRequest struct {
		TrackingNo    string `json:"tracking_no"`
		LogisticsCode string `json:"logistics_code"`
	}

	MerchantOrderSyncBizModel struct {
		OutBizNo          string                             `json:"out_biz_no"`
		BuyerId           string                             `json:"buyer_id"`
		SellerId          string                             `json:"seller_id"`
		PartnerId         string                             `json:"partner_id"`
		Amount            int32                              `json:"amount"`
		TradeNo           string                             `json:"trade_no"`
		ItemOrderList     []ItemOrderInfo                    `json:"item_order_list"`
		LogisticsInfoList []OrderLogisticsInformationRequest `json:"logistics_info_list"`
		ExtInfo           []OrderExtInfo                     `json:"ext_info"`
	}
)

func (MerchantOrderSyncReq) GetMethod() string {
	return "alipay.merchant.order.sync"
}
