package wechatqy

import (
	"context"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		appId       string
		agentId     int
		secret      string
		redirectUri string
	}
	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.appId = options.appId
	c.config.agentId = options.agentId
	c.config.secret = options.secret
	c.config.secret = options.secret
	c.config.redirectUri = options.redirectUri

	// 创建请求客户端
	c.httpClient = resty.New()
	if options.restyClient != nil {
		c.httpClient = options.restyClient
	}

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
