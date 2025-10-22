package qxwlwagnt

import "context"

type NewConfig struct {
	BaseURL   string // 接口地址
	UserName  string // userName
	AppKey    string // appKey
	AppSecret string // appSecret
}

// GetNewClient 返回一个全新的 *Client 实例，使用新配置，但复用当前的中间件、日志等行为
func (c *Client) GetNewClient(ctx context.Context, config NewConfig) *Client {
	// 1. 克隆 httpClient（复制中间件、调试、超时等，但不共享状态）
	newHTTPClient := c.httpClient.Clone(ctx)

	// 2. 创建新 Client
	newClient := &Client{
		config: struct {
			baseURL   string
			userName  string
			appKey    string
			appSecret string
		}{
			baseURL:   config.BaseURL,
			userName:  config.UserName,
			appKey:    config.AppKey,
			appSecret: config.AppSecret,
		},
		httpClient: newHTTPClient,
	}

	// 3. 设置新 baseURL
	if config.BaseURL != "" {
		newClient.httpClient.SetBaseURL(config.BaseURL)
	}

	return newClient
}
