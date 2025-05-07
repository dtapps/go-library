package kuaidi100

import (
	"context"
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 参数
	newParams := gorequest.NewParams()

	// 公共参数
	newParams.Set("customer", c.GetCustomer())

	// 请求参数
	newParams.Set("param", gorequest.JsonEncodeNoError(param.DeepGetAny()))

	// 签名
	newParams.Set("sign", c.getSign(gorequest.JsonEncodeNoError(param.DeepGetAny())))

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(newParams)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
