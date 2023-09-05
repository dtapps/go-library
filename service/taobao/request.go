package taobao

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, params map[string]interface{}) (gorequest.Response, error) {

	// 签名
	c.Sign(params)

	// 创建请求
	client := c.requestClient

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", params["method"]), request)
	}

	return request, err
}
