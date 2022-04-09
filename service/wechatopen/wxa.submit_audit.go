package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaSubmitAuditResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxaSubmitAuditResult struct {
	Result WxaSubmitAuditResponse // 结果
	Body   []byte                 // 内容
	Err    error                  // 错误
}

func NewWxaSubmitAuditResult(result WxaSubmitAuditResponse, body []byte, err error) *WxaSubmitAuditResult {
	return &WxaSubmitAuditResult{Result: result, Body: body, Err: err}
}

// WxaSubmitAudit 提交审核
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/submit_audit.html
func (app *App) WxaSubmitAudit(notMustParams ...Params) *WxaSubmitAuditResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/submit_audit?access_token=%s", app.authorizerAccessToken), params, http.MethodPost)
	// 定义
	var response WxaSubmitAuditResponse
	err = json.Unmarshal(body, &response)
	return NewWxaSubmitAuditResult(response, body, err)
}
