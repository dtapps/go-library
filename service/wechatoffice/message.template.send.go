package wechatoffice

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func NewMessageTemplateSendResult(result MessageTemplateSendResponse, body []byte, http gorequest.Response, err error) *MessageTemplateSendResult {
	return &MessageTemplateSendResult{Result: result, Body: body, Http: http, Err: err}
}

// MessageTemplateSend 模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (app *App) MessageTemplateSend(notMustParams ...Params) *MessageTemplateSendResult {
	app.accessToken = app.GetAccessToken()
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", app.accessToken), params, http.MethodPost)
	// 定义
	var response MessageTemplateSendResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewMessageTemplateSendResult(response, request.ResponseBody, request, err)
}
