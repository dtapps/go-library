package kuaidi100

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string) (gorequest.Response, error) {

	newParams := gorequest.NewParams()

	// 公共参数
	newParams["customer"] = c.GetCustomer()

	// 请求参数
	newParams["param"] = gojson.JsonEncodeNoError(params)

	// 签名
	newParams["sign"] = c.getSign(gojson.JsonEncodeNoError(params))

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.zap.status {
		go c.zap.client.Middleware(ctx, request)
	}

	return request, err
}
