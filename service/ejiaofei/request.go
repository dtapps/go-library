package ejiaofei

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, method string) (gorequest.Response, error) {

	// 公共参数
	params["userid"] = c.GetUserId()
	params["pwd"] = c.GetPwd()

	// 签名
	params["userkey"] = gomd5.ToUpper(fmt.Sprintf("%s%s", c.config.signStr, c.GetKey()))

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
	if c.log.status {
		go c.log.client.MiddlewareXml(ctx, request, go_library.Version())
	}

	return request, err
}
