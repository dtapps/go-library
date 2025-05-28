package wechatqy

import (
	"context"
	"fmt"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
	"resty.dev/v3"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) (err error) {

	if gorequest.IsHttpURL(url) == false {
		return fmt.Errorf("不是有效地址: %s", url)
	}

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置参数
	if method == http.MethodGet {
		httpClient.SetQueryParams(param.DeepGetString())
	} else {
		httpClient.SetBody(param.DeepGetAny())
	}

	// 设置结果
	httpClient.SetResult(&response)

	// 发起请求
	var resp = &resty.Response{}
	if method == http.MethodGet {
		resp, err = httpClient.Get(url)
		if err != nil {
			return err
		}
	} else {
		resp, err = httpClient.Post(url)
		if err != nil {
			return err
		}
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	return nil
}
