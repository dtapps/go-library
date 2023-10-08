package kashangwl

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
	"time"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 公共参数
	param.Set("timestamp", time.Now().UnixNano()/1e6)
	param.Set("customer_id", c.GetCustomerId())

	// 签名参数
	param.Set("sign", c.getSign(c.GetCustomerKey(), param))

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeJson()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.slog.status {
		go c.slog.client.Middleware(ctx, request)
	}

	return request, err
}
