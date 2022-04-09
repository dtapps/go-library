package workwechat

import (
	"encoding/json"
	"fmt"
)

// AuthAccessTokenResult 返回参数
type AuthAccessTokenResult struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AuthAccessToken 获取access_token https://open.work.weixin.qq.com/api/doc/90000/90135/91039
func (app *App) AuthAccessToken() (result AuthAccessTokenResult, err error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", app.AppID, app.Secret)

	// request
	body, err := app.request(url, map[string]interface{}{}, "GET")
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
