package pinduoduo

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

	// 日志
	if c.log.status {
		go c.log.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", params["type"]), request)
	}
	if c.zap.status {
		go c.zap.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", params["type"]), request)
	}

	return request, err
}
