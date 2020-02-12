package request

// 创建商品接口
// https://docs.open.alipay.com/api_4/ant.merchant.expand.item.open.create/
type (
	MerchantExpandItemOpenCreateReq struct {
		BaseRequest
	}

	ItemSkuPropertyInfo struct {
		PropertyKey   string `json:"property_key"`
		PropertyValue string `json:"property_value"`
	}

	SkuCreateInfo struct {
		Price         int32                 `json:"price"`
		OriginalPrice int32                 `json:"original_price"`
		Inventory     int32                 `json:"inventory"`
		MaterialList  []MaterialCreateInfo  `json:"material_list"`
		PropertyList  []ItemSkuPropertyInfo `json:"property_list"`
	}

	ItemExtInfo struct {
		ExtKey   string `json:"ext_key"`
		ExtValue string `json:"ext_value"`
	}

	MaterialCreateInfo struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}

	ItemPropertyList struct {
		PropertyKey       string   `json:"property_key"`
		PropertyValueList []string `json:"property_value_list"`
	}

	MerchantExpandItemOpenCreateBizModel struct {
		Scene              string               `json:"scene"`
		TargetId           string               `json:"target_id"`
		TargetType         string               `json:"target_type"` //8-小程序
		StandardCategoryId string               `json:"standard_category_id"`
		Name               string               `json:"name"`
		Description        string               `json:"description"`
		Type               string               `json:"type"` //STANDARD_GOODS-标品
		ExtInfo            []ItemExtInfo        `json:"ext_info"`
		SkuList            []SkuCreateInfo      `json:"sku_list"`
		MaterialList       []MaterialCreateInfo `json:"material_list"`
		PropertyList       []ItemPropertyList   `json:"property_list"`
	}
)

func (MerchantExpandItemOpenCreateReq) GetMethod() string {
	return "ant.merchant.expand.item.open.create"
}
