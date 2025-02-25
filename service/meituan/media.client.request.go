package meituan

import (
	"context"
	"encoding/json"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

const (
	mediaApiUrl   = "https://media.meituan.com/"
	MediaLogTable = "meituan_media"
)

func (c *MediaClient) request(ctx context.Context, url string, method string, bodyParam *gorequest.Params, response any) (gorequest.Response, error) {

	// 请求地址
	path := mediaApiUrl + url

	// 设置请求地址
	c.httpClient.SetUri(path)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置参数
	c.httpClient.SetParams(bodyParam)

	// Create a SignUtil instance
	signUtil := NewSignUtil(c.GetAppKey(), c.GetAppSecret())
	signHeaders := signUtil.GetSignHeaders(map[string]any{
		"method": method,
		"data":   bodyParam.DeepGet(),
		"url":    path,
	})
	fmt.Printf("%+v\n", signHeaders)
	c.httpClient.SetHeader("S-Ca-App", signHeaders.SCaApp)
	c.httpClient.SetHeader("S-Ca-Timestamp", signHeaders.SCaTimestamp)
	c.httpClient.SetHeader("S-Ca-Signature", signHeaders.SCaSignature)
	c.httpClient.SetHeader("S-Ca-Signature-Headers", signHeaders.SCaSignatureHeaders)
	c.httpClient.SetHeader("Content-MD5", signHeaders.ContentMD5)
	c.httpClient.SetHeader("Content-Type", "application/json")

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
