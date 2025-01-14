package cloud

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string, contentType string) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	switch contentType {
	case "JSON":
		c.httpClient.SetContentTypeJson()
	case "FORM":
		c.httpClient.SetContentTypeForm()
	}

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	return request, err
}
