package aswzk

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

// 请求接口
func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) error {

	// 获取时间戳
	xTimestamp := fmt.Sprintf("%v", gotime.Current().Timestamp())

	// 签名
	xSign := sign(param, c.GetApiKey(), xTimestamp)

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(c.GetUrl() + url)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置格式
	httpClient.SetContentType("application/json")

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 添加请求头
	httpClient.SetHeader("X-Timestamp", xTimestamp)
	httpClient.SetHeader("X-UserId", c.GetUserID())
	httpClient.SetHeader("X-Sign", xSign)

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
