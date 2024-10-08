package wechatqy

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
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
}

func newWebhookSendResult(result WebhookSendResponse, body []byte, http gorequest.Response) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Http: http}
}

// WebhookSend 群机器人 发送消息
// https://developer.work.weixin.qq.com/document/path/99110
func (c *Client) WebhookSend(ctx context.Context, key string, Type string, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "cgi-bin/webhook/send")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WebhookSendResponse
	request, err := c.request(ctx, span, apiUrl+fmt.Sprintf("cgi-bin/webhook/send?key=%s&type=%s", key, Type), params, http.MethodPost, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}

// WebhookSendURL 群机器人 发送消息
// https://developer.work.weixin.qq.com/document/path/99110
func (c *Client) WebhookSendURL(ctx context.Context, url string, Type string, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, url)
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WebhookSendResponse
	request, err := c.request(ctx, span, fmt.Sprintf("%s&type=%s", url, Type), params, http.MethodPost, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}
