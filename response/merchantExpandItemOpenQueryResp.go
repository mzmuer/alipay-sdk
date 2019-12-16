package response

import (
	"github.com/mzmuer/alipay-sdk/utils"
)

// 查询商品接口
type (
	CmdItemSkuInfo struct {
		OriginalPrice   int        `json:"original_price"`
		Price           int        `json:"price"`
		GmtCreate       utils.Time `json:"gmt_create"`
		GmtModified     utils.Time `json:"gmt_modified"`
		Inventory       int        `json:"inventory"`
		RemainInventory int        `json:"remain_inventory"`
		SkuID           string     `json:"sku_id"`
		Status          string     `json:"status"`
	}

	MaterialInfo struct {
		GmtModified utils.Time `json:"gmt_modified"`
		GmtCreate   utils.Time `json:"gmt_create"`
		MaterialID  string     `json:"material_id"`
		Type        string     `json:"type"`
		Content     string     `json:"content"`
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
		ItemID              string   `json:"item_id"`
		Type                string   `json:"type"`
		Status              string   `json:"status"`
		FrontCategoryIdList []string `json:"front_category_id_list"`
		StandardCategoryId  string   `json:"standard_category_id"`
		TargetID            string   `json:"target_id"`
		Description         string   `json:"description"`
		TargetType          string   `json:"target_type"`
		Name                string
		SkuList             []CmdItemSkuInfo   `json:"sku_list"`
		MaterialList        []MaterialInfo     `json:"material_list"`
		ExtInfo             []ItemExtInfo      `json:"ext_info"`
		PropertyList        []ItemPropertyInfo `json:"property_list"`
	}
	MerchantExpandItemOpenQueryResp struct {
		BaseResponse
		ItemList []CmdItemInfo `json:"item_list"`
	}
)
