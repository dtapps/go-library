package pinduoduo

import (
	"context"
	"encoding/json"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, param *gorequest.Params, response any) (gorequest.Response, error) {

	if c.clientIP == "" {
		requestIP, isIP := gorequest.GetRequestIPStr(ctx)
		if isIP {
			c.httpClient.SetClientIP(requestIP)
		}
	}

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
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
