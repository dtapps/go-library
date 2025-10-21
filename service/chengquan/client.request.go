package chengquan

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) error {

	// 公共参数
	param.Set("timestamp", gotime.Current().TimestampWithMillisecond()) // 时间戳，以毫秒为单位。校验开发者与橙券的时间差，橙券允许开发者请求最大时间误差为3分钟 (3*60*1000)
	param.Set("app_id", c.GetAppID())                                   // 商户账号，由橙券提供，如：13105356515

	// 签名
	param.Set("sign", c.sign(ctx, param))

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(url)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置格式
	httpClient.SetContentType("application/x-www-form-urlencoded")

	// 设置参数
	httpClient.SetFormData(param.DeepGetString())

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
