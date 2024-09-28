package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (c *Client) request(ctx context.Context, span trace.Span, param gorequest.Params, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl

	// 签名
	c.Sign(param)

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置参数
	c.httpClient.SetParams(param)

	// OpenTelemetry链路追踪
	span.SetAttributes(attribute.String("http.url", uri))
	span.SetAttributes(attribute.String("http.params", gojson.JsonEncodeNoError(param)))

	// 发起请求
	request, err := c.httpClient.Get(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return gorequest.Response{}, err
	}

	// 解析响应
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return request, err
}
