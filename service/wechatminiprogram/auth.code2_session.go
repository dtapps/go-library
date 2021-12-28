package wechatminiprogram

import (
	"encoding/json"
	"fmt"
)

// AuthCode2Session 请求参数
type AuthCode2Session struct {
	JsCode string `json:"js_code"`
}

// AuthCode2SessionResult 返回参数
type AuthCode2SessionResult struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	Unionid    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回
	Errcode    string `json:"errcode"`     // 错误码
	Errmsg     string `json:"errmsg"`      // 错误信息
}

func (app *App) AuthCode2Session(param AuthCode2Session) (result AuthCode2SessionResult, err error) {
	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", app.AppId, app.AppSecret, param.JsCode), map[string]interface{}{}, "GET")
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
