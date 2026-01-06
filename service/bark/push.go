package bark

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PushResponse struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// Push 推送消息
// https://bark.day.app/#/tutorial?id=请求方式
func (c *Client) Push(ctx context.Context, notMustParams ...*gorequest.Params) (response PushResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if c.config.pushKey != "" {
		params.Set("device_key", c.config.pushKey) // 设备key
	}

	// 请求
	err = c.request(ctx, "push", params, http.MethodPost, &response)
	return
}
