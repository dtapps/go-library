package aswzk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"net/url"
)

type NotifyUrlParams struct {
	NotifyUrl string `json:"notify_url"` // 回调地址
	UserID    string `json:"user_id"`    // 用户编号
	ApiKey    string `json:"api_key"`    // 秘钥
}

// NotifyUrl 通知回调地址
func (c *Client) NotifyUrl(ctx context.Context, params NotifyUrlParams, param gorequest.Params) error {

	// 验证回调地址
	_, err := url.ParseRequestURI(params.NotifyUrl)
	if err != nil {
		return err
	}

	// 检查密钥
	if params.ApiKey == "" {
		return errors.New("api_key cannot be empty")
	}

	// 获取时间戳
	xTimestamp := fmt.Sprintf("%v", gotime.Current().Timestamp())

	// 签名
	xSign := sign(param, params.ApiKey, xTimestamp)

	// 设置请求地址
	c.httpClient.SetUri(params.NotifyUrl)

	// 设置格式
	c.httpClient.SetContentTypeJson()

	// 设置参数
	c.httpClient.SetParams(param)

	// 添加请求头
	c.httpClient.SetHeader("X-Timestamp", xTimestamp)
	c.httpClient.SetHeader("X-Sign", xSign)

	// 发起请求
	request, err := c.httpClient.Post(ctx)
	if err != nil {
		return err
	}

	// 定义
	var response struct {
		Code int `json:"code"` // 状态码
	}
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return err
	}

	if response.Code == CodeSuccess {
		return nil
	}

	return errors.New(fmt.Sprintf("code: %v", response.Code))
}
