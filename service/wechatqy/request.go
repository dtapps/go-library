package wechatqy

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string) (gorequest.Response, error) {
	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置FORM格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 传入SDk版本
	client.AfferentSdkUserVersion(go_library.Version())

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.log.status {
		go c.log.client.Middleware(ctx, request, go_library.Version())
	}

	return request, err
}
