package request

// 换取授权访问令牌
type SystemOauthTokenReq struct {
	BaseRequest
	GrantType    string `json:"grant_type"`    // 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code"`          // 授权码，用户对应用授权后得到。
	RefreshToken string `json:"refresh_token"` // 刷刷新令牌，上次换取访问令牌时得到。见出参的refresh_token字段
}

func (SystemOauthTokenReq) GetMethod() string {
	return "alipay.system.oauth.token"
}

func (r *SystemOauthTokenReq) GetTextParams() map[string]string {
	m := r.UdfParams
	if m == nil {
		m = map[string]string{}
	}

	m["grant_type"] = r.GrantType
	m["code"] = r.Code
	m["refresh_token"] = r.RefreshToken
	return m
}
