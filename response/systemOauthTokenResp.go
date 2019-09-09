package response

// 换取授权访问令牌
type SystemOauthTokenResp struct {
	baseResponse
	UserId       string `json:"user_id"`       // 支付宝用户的唯一userId
	AccessToken  string `json:"access_token"`  // 访问令牌。通过该令牌调用需要授权类接口
	ExpiresIn    int64  `json:"expires_in"`    // 访问令牌的有效时间，单位是秒。
	RefreshToken string `json:"refresh_token"` // 刷新令牌。通过该令牌可以刷新access_token
	ReExpiresIn  int64  `json:"re_expires_in"` // 刷新令牌的有效时间，单位是秒。
}
