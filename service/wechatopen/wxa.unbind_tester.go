package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaUnbindTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaUnbindTesterResult struct {
	Result WxaUnbindTesterResponse // 结果
	Body   []byte                  // 内容
	Err    error                   // 错误
}

func NewWxaUnbindTesterResult(result WxaUnbindTesterResponse, body []byte, err error) *WxaUnbindTesterResult {
	return &WxaUnbindTesterResult{Result: result, Body: body, Err: err}
}

// WxaUnbindTester 解除绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/unbind_tester.html
func (app *App) WxaUnbindTester(wechatid, userstr string) *WxaUnbindTesterResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 参数
	params := NewParams()
	if wechatid != "" {
		params["wechatid"] = wechatid
	}
	params["userstr"] = userstr
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/unbind_tester?access_token=%s", app.authorizerAccessToken), params, http.MethodPost)
	// 定义
	var response WxaUnbindTesterResponse
	err = json.Unmarshal(body, &response)
	return NewWxaUnbindTesterResult(response, body, err)
}
