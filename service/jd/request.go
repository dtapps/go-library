package jd

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, param *gorequest.Params) (gorequest.Response, error) {

	// 签名
	c.Sign(param)

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(param.ToMap())

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", param.Get("method")), request)
	}

	return request, err
}
