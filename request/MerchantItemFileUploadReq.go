package request

import (
	"github.com/mzmuer/alipay-sdk/utils"
)

// 商品文件上传接口
type MerchantItemFileUploadReq struct {
	baseRequest
	Scene       string          `json:"scene"`        // 业务场景描述，比如订单信息同步场景对应SYNC_ORDER
	FileContent *utils.FileItem `json:"file_content"` // 文件二进制字节流，最大为4M
}

func (*MerchantItemFileUploadReq) GetMethod() string {
	return "alipay.merchant.item.file.upload"
}

func (r *MerchantItemFileUploadReq) GetTextParams() map[string]string {
	m := r.UdfParams
	if m == nil {
		m = map[string]string{}
	}

	m["scene"] = r.Scene
	return m
}

func (r *MerchantItemFileUploadReq) GetFileParams() map[string]*utils.FileItem {
	return map[string]*utils.FileItem{
		"file_content": r.FileContent,
	}
}
