# alipay-sdk
支付宝sdk的go版本

## 安装
使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/mzmuer/alipay-sdk
```
## 快速开始

### 使用密钥
```go
package main

import (
	"fmt"

	"github.com/mzmuer/alipay-sdk/request"
	"github.com/mzmuer/alipay-sdk/response"

	"github.com/mzmuer/alipay-sdk"
)

func main() {
	p, err := alipay.NewClient("your appId", "public key", "priv_key", false)
	if err != nil {
		panic(err)
	}
	
	req := request.TradeCreateReq{}
	req.NotifyUrl = "notify url"
	req.BizModel = request.TradeCreateBizModel{
		Body:           "body",
		Subject:        "subject",
		OutTradeNo:     "orderid",
		TimeoutExpress: "15m",
		TotalAmount:    "0.01",
		BuyerId:        "buyrtid",
	}

	result := response.TradeCreateResp{}
	_, err = p.Execute(&req, &result)

	fmt.Println(result, err)
}
```
### 使用证书
```go
package main

import (
	"fmt"

	"github.com/mzmuer/alipay-sdk/request"
	"github.com/mzmuer/alipay-sdk/response"

	"github.com/mzmuer/alipay-sdk"
)

func main() {
	c, err := alipay.NewCertClient("your appId", "privateKey", `appPubCert`, "alipayRootCert", "alipayPubCert", false)
	if err != nil {
		panic(err)
	}

	req := request.SystemOauthTokenReq{}
	req.GrantType = "authorization_code"
	req.Code = "cc6c559845a64762b24e2cd63c4fZX47"

	result := response.SystemOauthTokenResp{}
	_, err = c.Execute(&req, &result)

	fmt.Println(result, err)
}
```
### 如何使用未支持的接口
   ```go
package main

import (
	"fmt"

	"github.com/mzmuer/alipay-sdk/request"
	"github.com/mzmuer/alipay-sdk/response"

	"github.com/mzmuer/alipay-sdk"
)

type UimplementedReq struct {
	request.BaseRequest
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	// ....
}

// 如果接口参数是在biz_model中，参考（request.TradeCreateReq）
type UimplementedModel struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
    // ....
}

type UimplementedResp struct {
	response.BaseResponse
	Res1 string `json:"res1"`
	Res2 string `json:"res2"`
	// .....
}

// 必须实现该接口
func (*UimplementedReq) GetMethod() string {
	return "method.name"
}

// 如果除了request.BaseRequest还有额外参数，必须实现该接口（参考request.SystemOauthTokenReq）
func (r *UimplementedReq) GetTextParams() map[string]string {
	m := r.UdfParams
	if m == nil {
		m = map[string]string{}
	}

	m["field1"] = r.Field1
	m["field2"] = r.Field1
	return m
}


func main() {
	c, err := alipay.NewCertClient("your appId", "privateKey", `appPubCert`, "alipayRootCert", "alipayPubCert", false)
	if err != nil {
		panic(err)
	}

	req := UimplementedReq{}
	req.Field1 = ""
	req.Field2 = ""
	req.BizModel = UimplementedModel{
		Field1: "",
		Field2: "",
	}

	result := UimplementedResp{}
	_, err = c.Execute(&req, &result)

	fmt.Println(result, err)
}
```

## 目前支持接口
* [alipay.trade.refund](https://docs.open.alipay.com/api_1/alipay.trade.refund)(统一收单交易退款接口)
* [alipay.trade.wap.pay](https://docs.open.alipay.com/api_1/alipay.trade.wap.pay)(手机网站支付)
* [alipay.trade.app.pay](https://docs.open.alipay.com/api_1/alipay.trade.app.pay )(app支付接口2.0)
* [alipay.trade.fastpay.refund.query](https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query)(统一收单交易退款查询)
* [alipay.trade.create](https://docs.open.alipay.com/api_1/alipay.trade.create)(统一收单交易创建接口)
* [alipay.system.oauth.token](https://docs.open.alipay.com/api_9/alipay.system.oauth.token)(换取授权访问令牌)
* [alipay.user.info.share](https://docs.open.alipay.com/api_2/alipay.user.info.share)(支付宝会员授权信息查询接口)
* [alipay.open.app.mini.templatemessage.send](https://docs.open.alipay.com/api_5/alipay.open.app.mini.templatemessage.send)(小程序发送模板消息)
* [alipay.pass.template.update](https://docs.open.alipay.com/api_24/alipay.pass.template.update)(卡券模板更新接口)
* [alipay.pass.template.add](https://docs.open.alipay.com/api_24/alipay.pass.template.add)(卡券模板创建接口接口)
* [alipay.pass.instance.update](https://docs.open.alipay.com/api_24/alipay.pass.instance.update)(卡券实例更新接口)
* [alipay.pass.instance.add](https://docs.open.alipay.com/api_24/alipay.pass.instance.add)(卡券实例发放接口)
* [alipay.merchant.order.sync](https://docs.open.alipay.com/api_4/alipay.merchant.order.sync)(订单数据同步接口)
* [alipay.merchant.item.file.upload](https://docs.open.alipay.com/api_4/alipay.merchant.item.file.upload)(商品文件上传接口)
* [ant.merchant.expand.item.open.query](https://docs.open.alipay.com/api_4/ant.merchant.expand.item.open.query)(查询商品接口)
* [ant.merchant.expand.item.open.create](https://docs.open.alipay.com/api_4/ant.merchant.expand.item.open.create)(创建商品接口)
* [alipay.fund.trans.toaccount.transfer](https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer)(单笔转账到支付宝账户接口)
* [alipay.fund.trans.order.query](https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query)(查询转账订单接口)
* 持续跟新...