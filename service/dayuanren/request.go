package dayuanren

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 签名
	param.Set("sign", c.sign(param))

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(url)

	// 设置FORM格式
	client.SetContentTypeForm()

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
