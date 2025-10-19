package aswzk

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

type NotifyUrlParams struct {
	NotifyUrl string `json:"notify_url"` // 回调地址
	UserID    string `json:"user_id"`    // 用户编号
	ApiKey    string `json:"api_key"`    // 秘钥
}

// NotifyUrl 通知回调地址
func (c *Client) NotifyUrl(ctx context.Context, params NotifyUrlParams, param *gorequest.Params) error {

	// 验证回调地址
	_, err := url.ParseRequestURI(params.NotifyUrl)
	if err != nil {
		return err
	}

	// 检查密钥
	if params.ApiKey == "" {
		return errors.New("api_key cannot be empty")
	}

	// 定义
	var response struct {
		Code int `json:"code"` // 状态码
	}

	// 获取时间戳
	xTimestamp := fmt.Sprintf("%v", gotime.Current().Timestamp())

	// 签名
	xSign := sign(param, params.ApiKey, xTimestamp)

	// 创建请求客户端
	httpClient := c.httpClient.R().SetContext(ctx)

	// 设置格式
	httpClient.SetContentType("application/json")

	// 设置参数
	httpClient.SetBody(param.DeepGetAny())

	// 添加请求头
	httpClient.SetHeader("X-Timestamp", xTimestamp)
	httpClient.SetHeader("X-Sign", xSign)

	// 设置结果
	httpClient.SetResult(&response)

	// 发起请求
	resp, err := httpClient.Post(params.NotifyUrl)
	if err != nil {
		return err
	}

	// 检查 HTTP 状态码
	if resp.IsError() {
		return fmt.Errorf("请求失败，HTTP 状态码: %d", resp.StatusCode())
	}

	if response.Code == CodeSuccess {
		return nil
	}

	return fmt.Errorf("code: %v", response.Code)
}
