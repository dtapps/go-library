package wechatpayopen

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any, errResponse any) error {

	// 认证
	authorization, err := c.authorization(method, param.DeepGetAny(), c.config.baseURL+url)
	if err != nil {
		return err
	}

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(c.config.baseURL + url)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置JSON格式
	httpClient.SetContentType("application/json")

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 设置头部
	httpClient.SetHeader("Authorization", authorization)
	httpClient.SetHeader("Accept", "application/json")
	httpClient.SetHeader("Accept-Language", "zh-CN")

	// 设置结果
	httpClient.SetResult(&response)

	// 设置错误结果
	httpClient.SetError(&errResponse)

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
