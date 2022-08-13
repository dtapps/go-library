package wechatopen

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddleware(ctx, request, go_library.Version())
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddleware(ctx, request, go_library.Version())
	}

	return request, err
}
