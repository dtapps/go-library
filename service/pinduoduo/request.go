package pinduoduo

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, param *gorequest.Params) (gorequest.Response, error) {

	// 签名
	c.Sign(param)

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.slog.status {
		go c.slog.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", param.Get("type")), request)
	}

	return request, err
}
