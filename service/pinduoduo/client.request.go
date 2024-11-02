package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, param gorequest.Params, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl

	// 签名
	c.Sign(param)

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = gojson.Unmarshal(request.ResponseBody, &response)

	return request, err
}
