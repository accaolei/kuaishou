package kuaishou

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

const baseURL = "https://open.kuaishou.com"

// Kuaishou 快手开放平台的接口
type Kuaishou interface {
	AuthURL(scope, state string) string //生成网站授权的url
	// 换取token
	Code2AccessToken(ctx context.Context, code string) (*Code2AccessTokenResponse, error)
	// 刷新token
	RefreshAccessToken(ctx context.Context, refreshToken string) (*Code2AccessTokenResponse, error)
	// 获取用户信息
	GetUserInfo(ctx context.Context, accessToken string) (*UserInfo, error)

	GetUserVideoInfo(ctx context.Context, accessToken, cursor string, count int) ([]VideoInfo, error)
	GetUserVideoCount(ctx context.Context, accessToken string) (*VideoCount, error)
}

type ks struct {
	appID       string
	appSecret   string
	client      HTTPClient
	callBackURL string
}

func (k *ks) AuthURL(scope, state string) string {
	if scope == "" {
		scope = "user_info,user_video_info,user_video_delete,user_video_publish"
	}
	url := fmt.Sprintf("%s/oauth2/connect?app_id=%s&scope=%s&response_type=code&redirect_uri=%s&state=%s", baseURL, k.appID, scope, k.callBackURL, state)
	return url
}

func (k *ks) Code2AccessToken(ctx context.Context, code string) (*Code2AccessTokenResponse, error) {
	fmt.Println(baseURL)
	url := fmt.Sprintf("%s/oauth2/access_token?app_id=%s&app_secret=%s&code=%s&grant_type=authorization_code", baseURL, k.appID, k.appSecret, code)
	fmt.Println(url)
	resp, err := k.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(resp)
	if result := r.Get("result").Int(); result != 1 {
		return nil, fmt.Errorf("%d|%s", result, r.Get("error_msg"))
	}
	token := new(Code2AccessTokenResponse)
	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}
	return token, nil
}

func (k *ks) RefreshAccessToken(ctx context.Context, refreshToken string) (*Code2AccessTokenResponse, error) {
	resp, err := k.client.Get(ctx, fmt.Sprintf("%s/oauth2/refresh_token?app_id=%s&app_secret=%s&refresh_token=%s&grant_type=refresh_token", baseURL, k.appID, k.appSecret, refreshToken))
	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(resp)
	if result := r.Get("result").Int(); result != 1 {
		return nil, fmt.Errorf("%d|%s", result, r.Get("error_msg"))
	}
	token := new(Code2AccessTokenResponse)
	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}
	return token, nil
}
func (k *ks) GetUserInfo(ctx context.Context, accessToken string) (*UserInfo, error) {
	resp, err := k.client.Get(ctx, fmt.Sprintf("%s/openapi/user_info?app_id=%s&access_token=%s", baseURL, k.appID, accessToken))
	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(resp)
	if code := r.Get("result").Int(); code != 1 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("error_msg"))
	}
	fmt.Println(string(resp))
	type respData struct {
		Result   int8     `json:"result"`
		UserInfo UserInfo `json:"user_info"`
	}
	userInfo := new(respData)
	if err = json.Unmarshal(resp, userInfo); err != nil {
		return nil, err
	}
	return &userInfo.UserInfo, nil
}

// GetUserVideoInfo 作品列表
// cursor 非必填,游标，用于分页，值为作品id。分页查询时，传上一页create_time最小的photo_id。第一页不传此参数
// count 非必填,数量，默认为20
func (k *ks) GetUserVideoInfo(ctx context.Context, accessToken, cursor string, count int) ([]VideoInfo, error) {
	// user_video_info
	url := fmt.Sprintf("%s/openapi/photo/list?app_id=%s&access_token=%s", baseURL, k.appID, accessToken)
	if cursor != "" {
		url = fmt.Sprintf("%s&cursor=%s", url, cursor)
	}
	if count > 0 {
		url = fmt.Sprintf("%s&count=%d", url, count)
	}
	resp, err := k.client.Get(ctx, url)

	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(resp)
	if code := r.Get("result").Int(); code != 1 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("error_msg"))
	}
	type videos struct {
		VideoList []VideoInfo `json:"video_list"`
	}
	list := new(videos)
	if err = json.Unmarshal(resp, &list); err != nil {
		return nil, err
	}
	return list.VideoList, nil
}

// GetUserVideoCount 获取视频总数
func (k *ks) GetUserVideoCount(ctx context.Context, accessToken string) (*VideoCount, error) {
	resp, err := k.client.Get(ctx, fmt.Sprintf("%s/openapi/photo/count?app_id=%s&access_token=%s", baseURL, k.appID, accessToken))
	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(resp)
	if code := r.Get("result").Int(); code != 1 {
		return nil, fmt.Errorf("%d|%s", code, r.Get("error_msg"))
	}
	count := new(VideoCount)
	if err = json.Unmarshal(resp, &count); err != nil {
		return nil, err
	}
	return count, nil
}

// New new
func New(appID, appSecret, callBackURL string) Kuaishou {
	return &ks{
		appID:       appID,
		appSecret:   appSecret,
		callBackURL: callBackURL,
		client:      NewHTTPClient(),
	}
}
