package chengquan

import (
	"context"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL string // 接口地址
		appID   string
		appKey  string
		aesKey  string
		aesIv   string
		version string
	}

	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.baseURL = options.baseURL
	c.config.appID = options.appID
	c.config.appKey = options.appKey
	c.config.aesKey = options.aesKey
	c.config.aesIv = options.aesIv
	c.config.version = options.version

	// 创建请求客户端
	c.httpClient = resty.New()
	if options.restyClient != nil {
		c.httpClient = options.restyClient
	}

	// 设置基础 URL
	c.httpClient.SetBaseURL(c.config.baseURL)

	// 设置 Debug
	if options.debug {
		c.httpClient.EnableDebug()
	}

	return c, nil
}

// Close 关闭 请求客户端
func (c *Client) Close() (err error) {
	if c.httpClient != nil {
		err = c.httpClient.Close()
	}
	return
}
