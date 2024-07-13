package wikeyun

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param gorequest.Params, response any) (gorequest.Response, error) {

	// 签名
	sign := c.sign(param)

	// 拼接url
	uri := fmt.Sprintf("%s%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", c.GetApiUrl(), url, c.GetAppKey(), sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign)

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置FORM格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(param)

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
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}

	return request, err
}
