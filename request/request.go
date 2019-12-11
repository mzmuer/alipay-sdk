package request

import "github.com/mzmuer/alipay-sdk/utils"

type Request interface {
	GetMethod() string
	GetNotifyUrl() string
	GetReturnUrl() string
	GetProdCode() string
	GetTerminalType() string
	GetTerminalInfo() string
	GetApiVersion() string
	GetNeedEncrypt() bool
	GetTextParams() map[string]string
	GetBizModel() interface{}
}

type UploadRequest interface {
	GetFileParams() map[string]*utils.FileItem
}

type (
	// 请求结构
	baseRequest struct {
		Method       string // 请求方法
		NotifyUrl    string
		ReturnUrl    string
		ProdCode     string
		TerminalType string
		TerminalInfo string
		Version      string // 1.0
		NeedEncrypt  bool
		BizModel     interface{}
		BizContent   string // BizContent 设置之后就不会使用BizModel
		UdfParams    map[string]string
	}
)

// 获取API名称。
func (*baseRequest) GetMethod() string {
	return "undefined methods"
}

// 返回通知地址
func (r *baseRequest) GetNotifyUrl() string {
	return r.NotifyUrl
}

// 返回回跳地址
func (r *baseRequest) GetReturnUrl() string {
	return r.ReturnUrl
}

// 获取产品码
func (r *baseRequest) GetProdCode() string {
	return r.ProdCode
}

// 获取终端类型
func (r *baseRequest) GetTerminalType() string {
	return r.TerminalType
}

// 获取终端信息
func (r *baseRequest) GetTerminalInfo() string {
	return r.TerminalInfo
}

func (r *baseRequest) GetApiVersion() string {
	return "1.0"
}

// 判断是否需要加密
func (r *baseRequest) GetNeedEncrypt() bool {
	return r.NeedEncrypt
}

// 获取所有的Key-Value形式的文本请求参数集合
func (r *baseRequest) GetTextParams() map[string]string {
	m := r.UdfParams
	if m == nil {
		m = map[string]string{}
	}
	m["biz_content"] = r.BizContent
	return m
}

func (r *baseRequest) GetBizModel() interface{} {
	return r.BizModel
}

///**
// * 得到当前API的响应结果类型
// *
// * @return 响应类型
// */
//public Class<T> getResponseClass();
