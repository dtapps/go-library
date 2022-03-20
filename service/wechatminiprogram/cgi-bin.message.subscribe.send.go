package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gohttp"
	"net/http"
)

type SubscribeMessageSendResponse struct {
	Errcode int    // 错误码
	Errmsg  string // 错误信息
}

type SubscribeMessageSendResult struct {
	Result SubscribeMessageSendResponse // 结果
	Body   []byte                       // 内容
	Http   gohttp.Response              // 请求
	Err    error                        // 错误
}

func NewSubscribeMessageSendResult(result SubscribeMessageSendResponse, body []byte, http gohttp.Response, err error) *SubscribeMessageSendResult {
	return &SubscribeMessageSendResult{Result: result, Body: body, Http: http, Err: err}
}

// SubscribeMessageSend 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (app *App) SubscribeMessageSend(notMustParams ...Params) *SubscribeMessageSendResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response SubscribeMessageSendResponse
	err = json.Unmarshal(request.Body, &response)
	return NewSubscribeMessageSendResult(response, request.Body, request, err)
}
