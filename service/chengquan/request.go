package chengquan

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, method string) (gorequest.Response, error) {

	// 公共参数
	param.Set("timestamp", gotime.Current().TimestampWithMillisecond()) // 时间戳，以毫秒为单位。校验开发者与橙券的时间差，橙券允许开发者请求最大时间误差为3分钟 (3*60*1000)
	param.Set("app_id", c.GetAppID())                                   // 商户账号，由橙券提供，如：13105356515

	// 签名
	param.Set("sign", c.sign(ctx, param))

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(c.config.apiURL + url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(param)

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.MiddlewareXml(ctx, request)
	}

	return request, err
}
