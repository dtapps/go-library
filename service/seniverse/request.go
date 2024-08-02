package seniverse

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

func (c *V3Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrlV3 + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置用户代理
	c.httpClient.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	param.Set("key", c.key)
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	return request, err
}

func (c *V4Client) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrlV4 + url

	// 设置参数
	param.Set("ts", fmt.Sprintf("%d", time.Now().Unix()))
	param.Set("ttl", "600")
	param.Set("public_key", c.publicKey)

	// 签名并返回请求地址
	urlStr := c.sign(uri, param)

	// 设置请求地址
	c.httpClient.SetUri(urlStr)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	return request, err
}
