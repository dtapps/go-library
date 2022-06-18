package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type SubscribeMessageSendResponse struct {
	Errcode int    // 错误码
	Errmsg  string // 错误信息
}

type SubscribeMessageSendResult struct {
	Result SubscribeMessageSendResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func NewSubscribeMessageSendResult(result SubscribeMessageSendResponse, body []byte, http gorequest.Response, err error) *SubscribeMessageSendResult {
	return &SubscribeMessageSendResult{Result: result, Body: body, Http: http, Err: err}
}

// SubscribeMessageSend 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (c *Client) SubscribeMessageSend(notMustParams ...Params) *SubscribeMessageSendResult {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response SubscribeMessageSendResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewSubscribeMessageSendResult(response, request.ResponseBody, request, err)
}
