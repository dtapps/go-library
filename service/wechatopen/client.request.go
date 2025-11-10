package wechatopen

import (
	"context"
	"fmt"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, path string, param *gorequest.Params, method string, response any) (err error) {

	// 判断path前面有没有/
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	urlStr := fmt.Sprintf("%s%s", c.config.baseURL, path)

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(urlStr)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置格式
	httpClient.SetContentType("application/json")

	// 设置参数
	httpClient.SetBody(param.DeepCopy())

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
