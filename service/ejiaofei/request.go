package ejiaofei

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) requestXml(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 签名
	param.Set("userkey", c.xmlSign(url, param))

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = xml.Unmarshal(request.ResponseBody, &response)

	return request, err
}

func (c *Client) requestJson(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 签名
	param.Set("sign", c.jsonSign(param))

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)
	return request, err
}
