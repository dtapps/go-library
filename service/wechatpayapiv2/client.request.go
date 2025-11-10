package wechatpayapiv2

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
	"resty.dev/v3"
)

func (c *Client) request(ctx context.Context, path string, param *gorequest.Params, cert *tls.Certificate, response any) error {

	// 判断path前面有没有/
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	urlStr := fmt.Sprintf("%s%s", c.config.baseURL, path)

	var httpClient *resty.Request
	if cert != nil {
		// 创建请求客户带证书
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{*cert},
		}
		httpClient = c.httpClient.SetTLSClientConfig(tlsConfig).R().SetContext(ctx)
	} else {
		// 创建请求客户端
		httpClient = c.httpClient.R().SetContext(ctx)
	}

	// 设置请求地址
	httpClient.SetURL(urlStr)

	// 设置方式
	httpClient.SetMethod(http.MethodPost)

	// 设置格式
	httpClient.SetContentType("application/xml")

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

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
