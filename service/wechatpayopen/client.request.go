package wechatpayopen

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string, response any, errResponse any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 认证
	authorization, err := c.authorization(method, param, uri)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
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

	// OpenTelemetry链路追踪
	c.TraceSetAttributes(attribute.String("http.url", uri))
	c.TraceSetAttributes(attribute.String("http.method", method))
	c.TraceSetAttributes(attribute.String("http.params", gojson.JsonEncodeNoError(param)))

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return gorequest.Response{}, err
	}

	// 解析响应
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}

	// 解析错误响应
	err = gojson.Unmarshal(request.ResponseBody, &errResponse)

	return request, err
}
