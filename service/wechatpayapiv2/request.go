package wechatpayapiv2

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, certStatus bool, cert *tls.Certificate, response any) (gorequest.Response, error) {

	// 请求地址
	uri := apiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置格式
	c.httpClient.SetContentTypeXml()

	// 设置参数
	c.httpClient.SetParams(param)

	// 设置证书
	if certStatus {
		c.httpClient.SetP12Cert(cert)
	}

	// 发起请求
	request, err := c.httpClient.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = xml.Unmarshal(request.ResponseBody, &response)

	return request, err
}
