package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) DoRequest(ctx context.Context, path string, param *gorequest.Params, method string, response any, errResponse any) (err error) {

	// 判断path前面有没有/
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	urlStr := fmt.Sprintf("%s%s", c.config.baseURL, path)

	// 认证 + 获取统一的 body bytes
	signResult, err := Sign(&SignParams{
		Method:              method,
		Body:                param.DeepGetAny(),
		Url:                 urlStr,
		PrivateKey:          c.config.privateKey,
		MchId:               c.config.mchId,
		CertificateSerialNo: c.config.certificateSerialNo,
	})

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(urlStr)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置参数
	httpClient.SetBody(signResult.BodyBytes)

	// 设置头部
	httpClient.SetContentType("application/json")
	httpClient.SetHeader("Accept", "application/json")
	httpClient.SetHeader("Wechatpay-Serial", c.config.publicKeyID)
	httpClient.SetHeader("Authorization", signResult.Authorization)
	httpClient.SetHeader("Accept-Language", "zh-CN")

	// 设置错误结果
	httpClient.SetError(&errResponse)

	// 发起请求
	resp, err := httpClient.Send()
	if err != nil {
		return err
	}

	// 解析结果
	err = json.Unmarshal(resp.Bytes(), &response)
	if err != nil {
		return err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	return err
}
