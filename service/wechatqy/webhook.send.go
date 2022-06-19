package wechatqy

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
)

type WebhookSendResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type WebhookSendResult struct {
	Result WebhookSendResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func newWebhookSendResult(result WebhookSendResponse, body []byte, http gorequest.Response, err error) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Http: http, Err: err}
}

// WebhookSend 发送应用消息
// https://developer.work.weixin.qq.com/document/path/90372
func (c *Client) WebhookSend(notMustParams ...gorequest.Params) *WebhookSendResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/cgi-bin/webhook/send?key=%s&type=%s", c.GetKey(), "text"), params)
	// 定义
	var response WebhookSendResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWebhookSendResult(response, request.ResponseBody, request, err)
}
