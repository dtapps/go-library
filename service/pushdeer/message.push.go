package pushdeer

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type MessagePushResponse struct {
	Code    int64  `json:"code"`
	Content any    `json:"content,omitempty"`
	Error   string `json:"error,omitempty"`
}

type MessagePushResult struct {
	Result MessagePushResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newMessagePushResult(result MessagePushResponse, body []byte, http gorequest.Response) *MessagePushResult {
	return &MessagePushResult{Result: result, Body: body, Http: http}
}

// MessagePush 推送消息
// https://www.pushdeer.com/dev.html
func (c *Client) MessagePush(ctx context.Context, text string, notMustParams ...gorequest.Params) (*MessagePushResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if c.config.pushKey != "" {
		params.Set("pushkey", c.config.pushKey)
	}
	params.Set("text", text) // 推送消息内容

	// 请求
	var response MessagePushResponse
	request, err := c.request(ctx, c.config.pushKey+"message/push", params, http.MethodPost, &response)
	return newMessagePushResult(response, request.ResponseBody, request), err
}
