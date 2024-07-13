package wechatpayapiv2

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, certStatus bool, cert *tls.Certificate, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置格式
	c.httpClient.SetContentTypeXml()

	// 设置参数
	c.httpClient.SetParams(param)

	// 设置证书
	if certStatus {
		c.httpClient.SetP12Cert(cert)
	}

	// OpenTelemetry链路追踪
	c.TraceSetAttributes(attribute.String("http.url", uri))
	c.TraceSetAttributes(attribute.String("http.params", gojson.JsonEncodeNoError(param)))

	// 发起请求
	request, err := c.httpClient.Post(ctx)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return gorequest.Response{}, err
	}

	// 解析响应
	err = xml.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}

	return request, err
}
