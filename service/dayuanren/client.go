package dayuanren

import (
	"context"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL string // 接口地址
		userID  int64  // 商户ID
		apiKey  string // 秘钥
	}
	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.baseURL = options.baseURL
	c.config.userID = options.userID
	c.config.apiKey = options.apiKey

	// 创建请求客户端
	c.httpClient = resty.New()
	if options.restyClient != nil {
		c.httpClient = options.restyClient
	}

	// 设置基础 URL
	c.httpClient.SetBaseURL(options.baseURL)

	// 设置 Debug
	if options.debug {
		c.httpClient.EnableDebug()
	}

	// 绑定日志钩子
	if options.restyLog != nil {
		// 请求中间件
		c.httpClient.SetRequestMiddlewares(
			options.restyLog.IntrusionRequest, // 自定义请求中间件，注入开始时间
			resty.PrepareRequestMiddleware,    // 官方请求中间件，创建RawRequest
			options.restyLog.BeforeRequest,    // 自定义请求中间件，记录开始时间和OTel
		)
		// 响应中间件
		c.httpClient.SetResponseMiddlewares(
			options.restyLog.CopyResponseBodyMiddleware, // 自定义请求中间件，将响应体拷贝到Context
			resty.AutoParseResponseMiddleware,           // 官方请求中间件，自动解析
			options.restyLog.AfterResponse,              // 自定义请求中间件，打印/保存
		)
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
