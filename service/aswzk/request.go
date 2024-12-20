package aswzk

import (
	"context"
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 请求地址
	uri := c.GetApiUrl() + url

	// 获取时间戳
	xTimestamp := fmt.Sprintf("%v", gotime.Current().Timestamp())

	// 签名
	xSign := sign(param, c.GetApiKey(), xTimestamp)

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// 添加请求头
	c.httpClient.SetHeader("X-Timestamp", xTimestamp)
	c.httpClient.SetHeader("X-UserId", c.GetUserID())
	c.httpClient.SetHeader("X-Sign", xSign)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
