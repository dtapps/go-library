package weishi

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string) (gorequest.Response, error) {

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置用户代理
	client.SetUserAgent(c.config.ua)

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.Middleware(ctx, request)
	}

	return request, err
}
