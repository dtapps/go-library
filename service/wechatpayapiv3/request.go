package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, commonParams bool, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 公共参数
	if method == http.MethodPost {
		if commonParams == true {
			param.Set("appid", c.GetAppId())
			param.Set("mchid", c.GetMchId())
		}
	}

	// 认证
	authorization, err := c.authorization(method, param.DeepGetAny(), uri)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置JSON格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// 设置头部
	c.httpClient.SetHeader("Authorization", "WECHATPAY2-SHA256-RSA2048 "+authorization)
	c.httpClient.SetHeader("Accept", "application/json")
	c.httpClient.SetHeader("Accept-Language", "zh-CN")
	if url == "https://api.mch.weixin.qq.com/v3/merchant-service/complaints-v2" {
		c.httpClient.SetHeader("Wechatpay-Serial", c.GetMchSslSerialNo())
	}

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
