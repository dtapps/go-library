package wechatminiprogram

import (
	"encoding/json"
	"fmt"
)

type AuthGetAccessTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

type AuthGetAccessTokenResult struct {
	Result AuthGetAccessTokenResponse // 结果
	Byte   []byte                     // 内容
	Err    error                      // 错误
}

func NewAuthGetAccessTokenResult(result AuthGetAccessTokenResponse, byte []byte, err error) *AuthGetAccessTokenResult {
	return &AuthGetAccessTokenResult{Result: result, Byte: byte, Err: err}
}

// AuthGetAccessToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (app *App) AuthGetAccessToken() *AuthGetAccessTokenResult {
	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app.AppId, app.AppSecret), map[string]interface{}{}, "GET")
	// 定义
	var response AuthGetAccessTokenResponse
	err = json.Unmarshal(body, &response)
	return NewAuthGetAccessTokenResult(response, body, err)
}
