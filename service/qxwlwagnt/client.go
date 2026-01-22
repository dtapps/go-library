package qxwlwagnt

import (
	"context"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL   string // 接口地址
		userName  string // userName
		appKey    string // appKey
		appSecret string // appSecret
	}

	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.baseURL = options.baseURL
	c.config.userName = options.userName
	c.config.appKey = options.appKey
	c.config.appSecret = options.appSecret

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

	// 请求中间件
	// c.httpClient.SetRequestMiddlewares(
	// PreRequestMiddleware(options.debug, options.userName, options.appKey, options.appSecret), // 自定义请求中间件，签名
	// resty.PrepareRequestMiddleware, // 官方请求中间件，创建RawRequest
	// )
	// 响应中间件
	c.httpClient.SetResponseMiddlewares(
		Ensure2xxResponseMiddleware,       // 自定义请求中间件，判断状态
		resty.AutoParseResponseMiddleware, // 官方请求中间件，自动解析
	)

	return c, nil
}

// R 返回一个自定义的 Request，以便我们可以调用 SetBodyMap() SetBodyStruct() 解决因 body 顺序不同导致 SHA256 不一样的问题
func (c *Client) R() *Request {
	return &Request{c.httpClient.R()}
}

// Close 关闭 请求客户端
func (c *Client) Close() (err error) {
	if c.httpClient != nil {
		err = c.httpClient.Close()
	}
	return
}
