package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) WithUrlComponentAccessToken(url string) string {
	return url + "?access_token=" + c.GetComponentAccessToken()
}

func (c *Client) WithUrlAuthorizerAccessToken(url string) string {
	return url + "?access_token=" + c.GetAuthorizerAccessToken()
}

// requestImage 请求
func (c *Client) request(ctx context.Context, path string, param *gorequest.Params, method string, response any) (err error) {

	// 判断path前面有没有/
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	urlStr := fmt.Sprintf("%s%s", c.config.baseURL, path)

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(urlStr)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置格式
	httpClient.SetContentType("application/json")

	// 设置参数
	if method == http.MethodGet {
		httpClient.SetQueryParams(param.DeepGetString())
	} else {
		httpClient.SetBody(param.DeepCopy())
	}

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

// requestImage 请求图片
func (c *Client) requestImage(ctx context.Context, path string, param *gorequest.Params, method string, response any) (body []byte, err error) {

	// 判断path前面有没有/
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	urlStr := fmt.Sprintf("%s%s", c.config.baseURL, path)

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置请求地址
	httpClient.SetURL(urlStr)

	// 设置方式
	httpClient.SetMethod(method)

	// 设置格式
	httpClient.SetContentType("application/json")

	// 设置参数
	if method == http.MethodGet {
		httpClient.SetQueryParams(param.DeepGetString())
	} else {
		httpClient.SetBody(param.DeepCopy())
	}

	// 发起请求
	resp, err := httpClient.Send()
	if err != nil {
		return nil, err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return nil, fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	// 尝试判断是否为 JSON 错误（微信失败时返回 JSON，成功时返回 image/png）
	contentType := resp.Header().Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		return nil, json.Unmarshal(resp.Bytes(), &response)
	}

	// 否则认为是图片二进制数据
	return resp.Bytes(), nil
}
