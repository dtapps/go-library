package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type CgiBinComponentApiCreatePreAuthCodenResponse struct {
	PreAuthCode string `json:"pre_auth_code"` // 预授权码
	ExpiresIn   int64  `json:"expires_in"`    // 有效期，单位：秒
}

type CgiBinComponentApiCreatePreAuthCodenResult struct {
	Result CgiBinComponentApiCreatePreAuthCodenResponse // 结果
	Body   []byte                                       // 内容
	Http   gorequest.Response                           // 请求
	Err    error                                        // 错误
}

func NewCgiBinComponentApiCreatePreAuthCodenResult(result CgiBinComponentApiCreatePreAuthCodenResponse, body []byte, http gorequest.Response, err error) *CgiBinComponentApiCreatePreAuthCodenResult {
	return &CgiBinComponentApiCreatePreAuthCodenResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinComponentApiCreatePreAuthCoden 预授权码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
func (app *App) CgiBinComponentApiCreatePreAuthCoden() *CgiBinComponentApiCreatePreAuthCodenResult {
	// 参数
	param := NewParams()
	param["component_appid"] = app.componentAppId // 第三方平台 appid
	params := app.NewParamsWith(param)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%v", app.GetComponentAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiCreatePreAuthCodenResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinComponentApiCreatePreAuthCodenResult(response, request.ResponseBody, request, err)
}
