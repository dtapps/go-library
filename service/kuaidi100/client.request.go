package kuaidi100

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) error {

	// 参数
	newParams := gorequest.NewParams()

	// 公共参数
	newParams.Set("customer", c.GetCustomer())

	// 请求参数
	newParams.Set("param", gorequest.JsonEncodeNoError(param.DeepGetAny()))

	// 签名
	newParams.Set("sign", c.getSign(gorequest.JsonEncodeNoError(param.DeepGetAny())))

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(url)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置参数
	httpClient.SetFormData(newParams.DeepGetString())

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
