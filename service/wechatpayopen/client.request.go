package wechatpayopen

import (
	"context"
	"fmt"
	"net/url"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) NewRequest(ctx context.Context, path string, param *gorequest.Params, method string, response any, errResponse any) error {

	// 认证 + 获取统一的 body bytes
	signResult, err := Sign(&SignParams{
		Method:              method,
		Body:                param.DeepGetAny(),
		Url:                 fmt.Sprintf("%s%s", c.config.baseURL, path),
		PrivateKey:          c.config.privateKey,
		SpMchId:             c.config.spMchId,
		CertificateSerialNo: c.config.certificateSerialNo,
	})

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(fmt.Sprintf("%s%s", c.config.baseURL, path))

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

	// 设置结果
	httpClient.SetResult(&response)

	// 设置错误结果
	httpClient.SetError(&errResponse)

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

func (c *Client) request(ctx context.Context, path string, param *gorequest.Params, method string, response any, errResponse any) error {

	// TODO: 请准备商户开发必要参数，参考：https://pay.weixin.qq.com/doc/v3/partner/4013080340
	config, err := CreateMchConfig(
		c.config.spMchId,             // 商户号，是由微信支付系统生成并分配给每个商户的唯一标识符，商户号获取方式参考 https://pay.weixin.qq.com/doc/v3/partner/4013080340
		c.config.certificateSerialNo, // 商户API证书序列号，如何获取请参考 https://pay.weixin.qq.com/doc/v3/partner/4013058924
		c.config.privateKey,          // 商户API证书私钥文件路径，本地文件路径
		c.config.publicKeyID,         // 微信支付公钥ID，如何获取请参考 https://pay.weixin.qq.com/doc/v3/partner/4013038589
		c.config.publicKey,           // 微信支付公钥文件路径，本地文件路径
	)
	if err != nil {
		return err
	}

	reqUrl, err := url.Parse(fmt.Sprintf("%s%s", c.config.baseURL, path))
	if err != nil {
		return err
	}

	authorization, err := BuildAuthorization(config.MchId(), config.CertificateSerialNo(), config.PrivateKey(), method, reqUrl.RequestURI(), nil)
	if err != nil {
		return err
	}

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(fmt.Sprintf("%s%s", c.config.baseURL, path))

	// 设置方式
	httpClient.SetMethod(method)

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 设置头部
	httpClient.SetContentType("application/json")
	httpClient.SetHeader("Accept", "application/json")
	httpClient.SetHeader("Wechatpay-Serial", c.config.publicKeyID)
	httpClient.SetHeader("Authorization", authorization)
	httpClient.SetHeader("Accept-Language", "zh-CN")

	// 设置结果
	httpClient.SetResult(&response)

	// 设置错误结果
	httpClient.SetError(&errResponse)

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
