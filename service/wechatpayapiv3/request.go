package wechatpayapiv3

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, commonParams bool) (gorequest.Response, error) {

	// 公共参数
	if method == http.MethodPost {
		if commonParams == true {
			param.Set("appid", c.GetAppId())
			param.Set("mchid", c.GetMchId())
		}
	}

	// 认证
	authorization, err := c.authorization(method, param, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 创建请求
	client := c.requestClient
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(url)

	client.SetHeader("app_id", c.GetAppId())
	client.SetHeader("mch_id", c.GetMchId())

	// 设置方式
	client.SetMethod(method)

	// 设置JSON格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(param.ToMap())

	// 设置头部
	client.SetHeader("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")
	if url == "https://api.mch.weixin.qq.com/v3/merchant-service/complaints-v2" {
		client.SetHeader("Wechatpay-Serial", c.GetMchSslSerialNo())
	}

	// 发起请求
	request, err := client.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.slog.status {
		go c.slog.client.Middleware(ctx, request)
	}

	return request, err
}
