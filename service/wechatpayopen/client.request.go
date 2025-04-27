package wechatpayopen

import (
	"context"
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any, errResponse any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 认证
	authorization, err := c.authorization(method, param.DeepGetAny(), uri)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置JSON格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	c.httpClient.SetHeader("sub_appid", c.GetSubAppid())

	// 设置头部
	c.httpClient.SetHeader("Authorization", authorization)
	c.httpClient.SetHeader("Accept", "application/json")
	c.httpClient.SetHeader("Accept-Language", "zh-CN")

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	// 解析错误响应
	err = json.Unmarshal(request.ResponseBody, &errResponse)

	return request, err
}
