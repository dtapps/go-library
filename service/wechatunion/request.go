package wechatunion

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// 请求
func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string) (gorequest.Response, error) {

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	client.SetHeader("app_id", c.GetAppId())

	// 设置请求方式
	client.SetMethod(method)

	// 设置FORM格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.Middleware(ctx, request)
	}

	return request, err
}
