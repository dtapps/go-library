package wechatopen

import (
	"encoding/json"
	"net/http"
)

type CgiBinComponentApiComponentTokenResponse struct {
	ComponentAccessToken string `json:"component_access_token"` // 第三方平台 access_token
	ExpiresIn            int64  `json:"expires_in"`             // 有效期，单位：秒
}

type CgiBinComponentApiComponentTokenResult struct {
	Result CgiBinComponentApiComponentTokenResponse // 结果
	Body   []byte                                   // 内容
	Err    error                                    // 错误
}

func NewCgiBinComponentApiComponentTokenResult(result CgiBinComponentApiComponentTokenResponse, body []byte, err error) *CgiBinComponentApiComponentTokenResult {
	return &CgiBinComponentApiComponentTokenResult{Result: result, Body: body, Err: err}
}

// CgiBinComponentApiComponentToken 令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func (app *App) CgiBinComponentApiComponentToken() *CgiBinComponentApiComponentTokenResult {
	app.componentVerifyTicket = app.GetComponentVerifyTicket()
	// 参数
	param := NewParams()
	param["component_appid"] = app.ComponentAppId                // 第三方平台 appid
	param["component_appsecret"] = app.ComponentAppSecret        // 第三方平台 appsecret
	param["component_verify_ticket"] = app.componentVerifyTicket // 微信后台推送的 ticket
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://api.weixin.qq.com/cgi-bin/component/api_component_token", params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiComponentTokenResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinComponentApiComponentTokenResult(response, body, err)
}
