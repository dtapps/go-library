package pinduoduo

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

// 请求接口
func (c *Client) request(ctx context.Context, param *gorequest.Params, response any) error {

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 签名
	c.Sign(param)

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 设置结果
	httpClient.SetResult(&response)

	// 发起请求
	resp, err := httpClient.Post(apiUrl)
	if err != nil {
		return err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	return nil
}

// 请求接口
func (c *Client) requestAndErr(ctx context.Context, param *gorequest.Params, response any, apiErr any) error {

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 签名
	c.Sign(param)

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 设置结果
	httpClient.SetResult(&response)

	// 设置错误
	httpClient.SetError(&apiErr)

	// 发起请求
	resp, err := httpClient.Post(apiUrl)
	if err != nil {
		return err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	return nil
}
