package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CgiBinComponentApiCreatePreAuthCodenResponse struct {
	PreAuthCode string `json:"pre_auth_code"` // 预授权码
	ExpiresIn   int64  `json:"expires_in"`    // 有效期，单位：秒
}

type CgiBinComponentApiCreatePreAuthCodenResult struct {
	Result CgiBinComponentApiCreatePreAuthCodenResponse // 结果
	Body   []byte                                       // 内容
	Err    error                                        // 错误
}

func NewCgiBinComponentApiCreatePreAuthCodenResult(result CgiBinComponentApiCreatePreAuthCodenResponse, body []byte, err error) *CgiBinComponentApiCreatePreAuthCodenResult {
	return &CgiBinComponentApiCreatePreAuthCodenResult{Result: result, Body: body, Err: err}
}

// CgiBinComponentApiCreatePreAuthCoden 预授权码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
func (app *App) CgiBinComponentApiCreatePreAuthCoden() *CgiBinComponentApiCreatePreAuthCodenResult {
	app.componentAccessToken = app.GetComponentAccessToken()
	// 参数
	param := NewParams()
	param["component_appid"] = app.ComponentAppId // 第三方平台 appid
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%v", app.componentAccessToken), params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiCreatePreAuthCodenResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinComponentApiCreatePreAuthCodenResult(response, body, err)
}
