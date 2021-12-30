package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Oauth2AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	Openid       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

type Oauth2AccessTokenResult struct {
	Result Oauth2AccessTokenResponse // 结果
	Body   []byte                    // 内容
	Err    error                     // 错误
}

func NewOauth2AccessTokenResult(result Oauth2AccessTokenResponse, body []byte, err error) *Oauth2AccessTokenResult {
	return &Oauth2AccessTokenResult{Result: result, Body: body, Err: err}
}

// Oauth2AccessToken 通过code换取网页授权access_token
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (app *App) Oauth2AccessToken(code string) *Oauth2AccessTokenResult {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", app.AppId, app.AppSecret, code), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response Oauth2AccessTokenResponse
	err = json.Unmarshal(body, &response)
	return NewOauth2AccessTokenResult(response, body, err)
}
