package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SnsOauth2AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	Openid       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

type SnsOauth2AccessTokenResult struct {
	Result SnsOauth2AccessTokenResponse // 结果
	Body   []byte                       // 内容
	Err    error                        // 错误
}

func NewSnsOauth2AccessTokenResult(result SnsOauth2AccessTokenResponse, body []byte, err error) *SnsOauth2AccessTokenResult {
	return &SnsOauth2AccessTokenResult{Result: result, Body: body, Err: err}
}

// SnsOauth2AccessToken 通过code换取网页授权access_token
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html#0
func (app *App) SnsOauth2AccessToken(code string) *SnsOauth2AccessTokenResult {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", app.AppId, app.AppSecret, code), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response SnsOauth2AccessTokenResponse
	err = json.Unmarshal(body, &response)
	return NewSnsOauth2AccessTokenResult(response, body, err)
}
