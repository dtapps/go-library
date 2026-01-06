package feishu

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, response any) error {

	if gorequest.IsHttpURL(url) == false {
		return fmt.Errorf("不是有效地址: %s", url)
	}

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 设置结果
	httpClient.SetResult(&response)

	// 发起请求
	resp, err := httpClient.Post(url)
	if err != nil {
		return err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	return nil
}
