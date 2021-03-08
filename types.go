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

// VideoInfo 作品信息
type VideoInfo struct {
	PhotoID      string `json:"photo_id"` //作品id
	Caption      string `json:"caption"`  //作品标题
	Cover        string `json:"cover"`    //封面
	PlayURL      string `json:"play_url"` //播放链接
	CreateTime   int64  `json:"create_time"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	ViewCount    int64  `json:"view_count"`
	Pending      bool   `json:"pending"`
}

// VideoCount 视频数量
type VideoCount struct {
	PublicCount  int64 `json:"public_count"`  //公开视频数量
	FriendCount  int64 `json:"friend_count"`  //仅好友可见视频数量
	PrivateCount int64 `json:"private_count"` //仅自己可见视频数量
	AllCount     int64 `json:"all_count"`     //视频总量
}
