package wechatopen

import (
	"context"
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string, response any) (resp gorequest.Response, err error) {

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	c.httpClient.SetHeader("authorizer_appid", c.GetAuthorizerAppid())

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	if request.HeaderIsImg() == false {
		err = json.Unmarshal(request.ResponseBody, &response)
	}

	return request, err
}
