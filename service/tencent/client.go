package tencent

import (
	"context"
	"fmt"
	"net/url"

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
	if options.endpoint == "" {
		return nil, fmt.Errorf("check endpoint")
	}
	if _, err := url.Parse(options.endpoint); err != nil {
		return nil, fmt.Errorf("check endpoint: %w", err)
	}
	if options.secretID == "" {
		return nil, fmt.Errorf("check secret_id")
	}
	if options.secretKey == "" {
		return nil, fmt.Errorf("check secret_key")
	}

	// 创建请求客户端
	client := resty.New()

	// 设置基础 URL
	client.SetBaseURL(options.endpoint)

	// 设置 Debug
	if options.debug {
		client.SetDebug(true)
	}

	// 请求中间件
	client.SetRequestMiddlewares(
		resty.MiddlewareRequestCreate,                                              // 官方请求中间件，创建RawRequest
		PreRequestMiddleware(options.endpoint, options.secretID, options.secretKey), // 自定义请求中间件，签名
	)
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
	return &Request{c.Client.R()}
}
