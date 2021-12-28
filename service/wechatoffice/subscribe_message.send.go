package wechatoffice

import (
	"encoding/json"
	"fmt"
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

// SubscribeMessageSendResult 返回参数
type SubscribeMessageSendResult struct {
	Errcode int    // 错误码
	Errmsg  string // 错误信息
}

func (app *App) SubscribeMessageSend(param SubscribeMessageSend) (result SubscribeMessageSendResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}

	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", app.AccessToken), params, "POST")
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
