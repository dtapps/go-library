package kuaidi100

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string) (gorequest.Response, error) {

	newParams := gorequest.NewParams()

	// 公共参数
	newParams.Set("customer", c.GetCustomer())

	// 请求参数
	newParams.Set("param", gojson.JsonEncodeNoError(param))

	// 签名
	newParams.Set("sign", c.getSign(gojson.JsonEncodeNoError(param)))

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(newParams)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.Middleware(ctx, request)
	}

	return request, err
}
