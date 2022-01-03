package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SubscribeMessageSend 入参
type SubscribeMessageSend struct {
	Touser           string                 `json:"touser"`                      // 接收者（用户）的 openid
	TemplateId       string                 `json:"template_id"`                 // 所需下发的订阅模板id
	Page             string                 `json:"page,omitempty"`              // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]interface{} `json:"data"`                        // 模板内容，格式形如 { "key1": { "value": any }, "key2": { "value": any } }
	MiniprogramState string                 `json:"miniprogram_state,omitempty"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string                 `json:"lang,omitempty"`              // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

type SubscribeMessageSendResponse struct {
	Errcode int    // 错误码
	Errmsg  string // 错误信息
}

type SubscribeMessageSendResult struct {
	Result SubscribeMessageSendResponse // 结果
	Body   []byte                       // 内容
	Err    error                        // 错误
}

func NewSubscribeMessageSendResult(result SubscribeMessageSendResponse, body []byte, err error) *SubscribeMessageSendResult {
	return &SubscribeMessageSendResult{Result: result, Body: body, Err: err}
}

// SubscribeMessageSend 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (app *App) SubscribeMessageSend(notMustParams ...Params) *SubscribeMessageSendResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response SubscribeMessageSendResponse
	err = json.Unmarshal(body, &response)
	return NewSubscribeMessageSendResult(response, body, err)
}
