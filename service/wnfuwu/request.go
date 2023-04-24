package wnfuwu

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, params map[string]interface{}) (gorequest.Response, error) {

	// 签名
	params["sign"] = c.sign(params)

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置FORM格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 传入SDk版本
	client.AfferentSdkUserVersion(go_library.Version())

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.log.status {
		go c.log.client.Middleware(ctx, request, go_library.Version())
	}

	return request, err
}
