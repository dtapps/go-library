package cloudflare

import (
	"context"
	"fmt"

	"resty.dev/v3"
)

type Client struct {
	options *Options
	*resty.Client
}

// NewClient 创建请求客户端
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {

	// 判断配置
	options := NewOptions(opts)
	if options.apiKey == "" {
		return nil, fmt.Errorf("check api_key")
	}
	if options.baseURL == "" {
		options.baseURL = "https://api.cloudflare.com/client/v4"
	}

	// 创建请求客户端
	client := resty.New()

	// 设置基础 URL
	client.SetBaseURL(options.baseURL)

	// 设置 Debug
	if options.debug {
		client.SetDebug(true)
	}

	// 设置令牌
	client.SetAuthToken(options.apiKey)

	// 响应中间件
	client.SetResponseMiddlewares(
		Ensure2xxResponseMiddleware,       // 自定义请求中间件，判断状态
		resty.MiddlewareResponseAutoParse, // 官方请求中间件，自动解析
	)

	return &Client{
		options: options,
		Client:  client,
	}, nil
}

// WithDebug 开启调试模式
func (c *Client) WithDebug() *Client {
	c.SetDebug(true)
	return c
}

// R 返回一个自定义的 Request，以便我们可以调用 SetBodyMap() SetBodyStruct() 解决因 body 顺序不同导致 SHA256 不一样的问题
func (c *Client) R() *Request {
	return &Request{c.Client.R().SetContentType("application/json")}
}
