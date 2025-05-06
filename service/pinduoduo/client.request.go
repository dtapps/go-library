package pinduoduo

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, param *gorequest.Params, response any) (gorequest.Response, error) {
	err := c.restyRequestV3(ctx, param, response)
	return gorequest.Response{}, err
}

func (c *Client) restyRequestV3(ctx context.Context, param *gorequest.Params, response any) error {

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)
	defer c.httpClient.Close()

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

	return err
}
