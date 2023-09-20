package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string) (gorequest.Response, error) {

	// 认证
	authorization, err := c.authorization(method, param, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	client.SetHeader("sp_appid", c.GetSpAppid())
	client.SetHeader("sp_mch_id", c.GetSpMchId())
	client.SetHeader("sub_appid", c.GetSubAppid())
	client.SetHeader("sub_mch_id", c.GetSubMchId())

	// 设置方式
	client.SetMethod(method)

	// 设置JSON格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(param)

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置头部
	client.SetHeader("Authorization", authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")

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
