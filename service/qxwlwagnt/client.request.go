package qxwlwagnt

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) Request(ctx context.Context, url string, param *gorequest.Params, method string, response any) error {

	// 参数
	newParams := gorequest.NewParams()

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(url)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置参数
	httpClient.SetContentType("application/json")
	httpClient.SetBody(newParams.DeepGetString())

	// 设置结果
	httpClient.SetResult(&response)

	// 发起请求
	resp, err := httpClient.Send()
	if err != nil {
		return err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	return nil
}
