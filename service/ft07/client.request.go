package ft07

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	if gorequest.IsHttpURL(url) == false {
		return gorequest.Response{}, fmt.Errorf("不是有效地址: %s", url)
	}

	// 设置请求地址
	c.httpClient.SetUri(url)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = gojson.Unmarshal(request.ResponseBody, &response)

	return request, err
}
