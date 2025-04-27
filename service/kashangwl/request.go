package kashangwl

import (
	"context"
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 公共参数
	param.Set("timestamp", time.Now().UnixNano()/1e6)
	param.Set("customer_id", c.GetCustomerId())

	// 签名参数
	param.Set("sign", c.getSign(c.GetCustomerKey(), param))

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
