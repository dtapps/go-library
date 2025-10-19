package pushdeer

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type MessagePushResponse struct {
	Code    int64  `json:"code"`
	Content any    `json:"content,omitempty"`
	Error   string `json:"error,omitempty"`
}

// MessagePush 推送消息
// https://www.pushdeer.com/dev.html
func (c *Client) MessagePush(ctx context.Context, text string, notMustParams ...*gorequest.Params) (response MessagePushResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if c.config.pushKey != "" {
		params.Set("pushkey", c.config.pushKey)
	}
	params.Set("text", text) // 推送消息内容

	// 请求
	err = c.request(ctx, "message/push", params, http.MethodPost, &response)
	return
}
