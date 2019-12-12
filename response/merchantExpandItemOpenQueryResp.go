package response

import "time"

// 创建商品接口
type (
	CmdItemSkuInfo struct {
		SkuId         string    `json:"sku_id"`
		ItemId        string    `json:"item_id"`
		Price         int32     `json:"price"`
		OriginalPrice int32     `json:"original_price"`
		Status        string    `json:"status"`
		GmtCreate     time.Time `json:"gmt_create"`
		GmtModified   time.Time `json:"gmt_modified"`
		Inventory     int32     `json:"inventory"`
	}

	MaterialInfo struct {
		MaterialId string `json:"material_id"`
		Type       string `json:"type"`
		Content    string `json:"content"`
	}
	ItemExtInfo struct {
		ExtKey   string `json:"ext_key"`
		ExtValue string `json:"ext_value"`
	}

	ItemPropertyInfo struct {
		PropertyKey       string   `json:"property_key"`
		PropertyValueList []string `json:"property_value_list"`
	}
	CmdItemInfo struct {
		ItemId              string   `json:"item_id"`
		Type                string   `json:"type"`
		Status              string   `json:"status"`
		FrontCategoryIdList []string `json:"front_category_id_list"`
		StandardCategoryId  string   `json:"standard_category_id"`
		TargetId            string   `json:"target_id"`
		Description         string   `json:"description"`
		TargetType          string   `json:"target_type"`
		Name                string
		GmtCreate           time.Time          `json:"gmt_create"`
		GmtModified         time.Time          `json:"gmt_modified"`
		SkuList             []CmdItemSkuInfo   `json:"sku_list"`
		MaterialList        []MaterialInfo     `json:"material_list"`
		ExtInfo             []ItemExtInfo      `json:"ext_info"`
		PropertyList        []ItemPropertyInfo `json:"property_list"`
	}
	MerchantExpandItemOpenQueryResp struct {
		BaseResponse
		ItemList CmdItemInfo `json:"item_list"`
	}
)
