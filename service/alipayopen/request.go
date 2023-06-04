package alipayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
)

func (c *Client) request(ctx context.Context, param map[string]interface{}) (gorequest.Response, error) {

	// 签名
	params := c.sign(ctx, param)

	// 创建请求
	client := c.requestClient

	// 设置参数
	client.SetParams(params)

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.zap.status {
		go c.zap.client.MiddlewareCustom(ctx, gostring.ToString(params["method"]), request)
	}

	return request, err
}
