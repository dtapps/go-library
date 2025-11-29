package tencent

import (
	"context"
	"fmt"
	"net/url"

	"resty.dev/v3"
)

type Client struct {
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
		client.EnableDebug()
	}

	// 绑定日志钩子
	if options.restyLog != nil {
		// 请求中间件
		client.SetRequestMiddlewares(
			options.restyLog.IntrusionRequest,                                           // 自定义请求中间件，注入开始时间
			resty.PrepareRequestMiddleware,                                              // 官方请求中间件，创建RawRequest
			PreRequestMiddleware(options.endpoint, options.secretID, options.secretKey), // 自定义请求中间件，签名
			options.restyLog.BeforeRequest,                                              // 自定义请求中间件，记录开始时间和OTel
		)
		// 响应中间件
		client.SetResponseMiddlewares(
			options.restyLog.CopyResponseBodyMiddleware, // 自定义请求中间件，将响应体拷贝到Context
			Ensure2xxResponseMiddleware,                 // 自定义请求中间件，判断状态
			resty.AutoParseResponseMiddleware,           // 官方请求中间件，自动解析
			options.restyLog.AfterResponse,              // 自定义请求中间件，打印/保存
		)
	} else {
		client.SetRequestMiddlewares(
			resty.PrepareRequestMiddleware,                                              // 先调用，创建 RawRequest
			PreRequestMiddleware(options.endpoint, options.secretID, options.secretKey), // 自定义请求中间件，签名
		)
		// 响应中间件
		client.SetResponseMiddlewares(
			Ensure2xxResponseMiddleware,       // 自定义请求中间件，判断状态
			resty.AutoParseResponseMiddleware, // 官方请求中间件，自动解析
		)
	}

	return &Client{client}, nil
}

// WithDebug 开启调试模式
func (c *Client) WithDebug() *Client {
	c.EnableDebug()
	return c
}

// R 返回一个自定义的 Request，以便我们可以调用 SetBodyMap() SetBodyStruct() 解决因 body 顺序不同导致 SHA256 不一样的问题
func (c *Client) R() *Request {
	return &Request{c.Client.R()}
}
