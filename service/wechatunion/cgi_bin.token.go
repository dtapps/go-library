package wechatunion

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type authGetAccessTokenResult struct {
	Byte   []byte                     // 内容
	Result authGetAccessTokenResponse // 结果
	Err    error                      // 错误
}

// NewAuthGetAccessTokenResult 实例化
func NewAuthGetAccessTokenResult(byte []byte, result authGetAccessTokenResponse, err error) *authGetAccessTokenResult {
	return &authGetAccessTokenResult{
		Byte:   byte,
		Result: result,
		Err:    err,
	}
}

// AuthGetAccessToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (app *App) AuthGetAccessToken() *authGetAccessTokenResult {
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", app.AppId, app.AppSecret), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response authGetAccessTokenResponse
	err = json.Unmarshal(body, &response)
	return NewAuthGetAccessTokenResult(body, response, err)
}

// 返回参数
type authGetAccessTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}
