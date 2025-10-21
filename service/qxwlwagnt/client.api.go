package qxwlwagnt

import (
	"context"
	"time"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) Request(ctx context.Context, url string, param *gorequest.Params, method string, response any) error {

	// 参数
	newParams := gorequest.NewParams()

	// 公共参数
	timestamp := time.Now().Format("20060102150405")
	newParams.Set("appKey", c.config.appKey)
	newParams.Set("userName", c.config.userName)
	newParams.Set("timeStamp", timestamp)
	param.Set("sign", c.getSign(newParams.DeepGetString()))

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(url)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置类型
	// httpClient.SetContentType("application/json")

	// 设置参数
	httpClient.SetQueryParams(newParams.DeepGetString())
	httpClient.SetQueryParams(param.DeepGetString())

	// 设置结果
	httpClient.SetResult(&response)

	// 发起请求
	_, err := httpClient.Send()
	if err != nil {
		return err
	}

	return nil
}
