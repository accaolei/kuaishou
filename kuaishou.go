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
	AuthURL(state string) string //生成网站授权的url
	// 换取token
	Code2AccessToken(ctx context.Context, code string) (*Code2AccessTokenResponse, error)
	// 刷新token
	RefreshAccessToken(ctx context.Context, refreshToken string) (*Code2AccessTokenResponse, error)

	GetUserInfo(ctx context.Context, accessToken string) (*UserInfo, error)
}

type ks struct {
	appID       string
	appSecret   string
	client      HTTPClient
	callBackURL string
}

func (k *ks) AuthURL(state string) string {
	url := fmt.Sprintf("%s/oauth2/connect?app_id=%s&scope=user_info&response_type=code&redirect_uri=%s&state=%s", baseURL, k.appID, k.callBackURL, state)
	return url
}

func (k *ks) Code2AccessToken(ctx context.Context, code string) (*Code2AccessTokenResponse, error) {
	resp, err := k.client.Get(ctx, fmt.Sprintf("%s/oauth2/access_token&app_id=%s&app_secret=%s&code=%s&grant_type=authorization_code", baseURL, k.appID, k.appSecret, code))
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
	resp, err := k.client.Get(ctx, fmt.Sprintf("%s/oauth2/refresh_token&app_id=%s&app_secret=%s&refresh_token=%s&grant_type=authorization_code", baseURL, k.appID, k.appSecret, refreshToken))
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
	// /openapi/user_info
	return nil, nil
}

// New new
func New(appID, appSecret string) Kuaishou {
	return &ks{
		appID:     appID,
		appSecret: appSecret,
	}
}
