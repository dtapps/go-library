package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageTemplateSendResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Msgid   int    `json:"msgid"`
}

type MessageTemplateSendResult struct {
	Result MessageTemplateSendResponse // 结果
	Body   []byte                      // 内容
	Err    error                       // 错误
}

func NewMessageTemplateSendResult(result MessageTemplateSendResponse, body []byte, err error) *MessageTemplateSendResult {
	return &MessageTemplateSendResult{Result: result, Body: body, Err: err}
}

// MessageTemplateSend 模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (app *App) MessageTemplateSend(notMustParams ...Params) *MessageTemplateSendResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response MessageTemplateSendResponse
	err = json.Unmarshal(body, &response)
	return NewMessageTemplateSendResult(response, body, err)
}
