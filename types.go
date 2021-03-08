package kuaishou

// Code2AccessTokenResponse 授权token的获取
type Code2AccessTokenResponse struct {
	Result       int8     `json:"result"`
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int64    `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	OpenID       string   `json:"open_id"`
	Scopes       []string `json:"scopes"`
}

// UserInfo 用户信息
type UserInfo struct {
	Name    int8   `json:"name"`     //用户昵称
	Sex     string `json:"sex"`      //性别，M:男性，F:女性，其他：未知
	Fan     int    `json:"fan"`      // 粉丝数量
	Follow  int    `json:"follow"`   //关注数
	Head    string `json:"head"`     //头像
	BigHead string `json:"big_head"` //大头像（可能为空）
	City    string `json:"city"`     //城市
}
