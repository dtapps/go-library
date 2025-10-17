package pinduoduo

import (
	"context"

	"resty.dev/v3"
)

const (
	baseURL = "https://gw-api.pinduoduo.com/api/router"

	Version = "1.0.64"
)

// Client 实例
type Client struct {
	config struct {
		clientId         string   // POP分配给应用的client_id
		clientSecret     string   // POP分配给应用的client_secret
		mediaId          string   // 媒体ID
		pid              string   // 推广位
		accessToken      string   // 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
		accessTokenScope []string // 授权范围
	}
	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.clientId = options.clientId
	c.config.clientSecret = options.clientSecret
	c.config.mediaId = options.mediaId
	c.config.pid = options.pid
	c.config.accessToken = options.accessToken
	c.config.accessTokenScope = options.accessTokenScope

	// 创建请求客户端
	c.httpClient = resty.New()
	if options.restyClient != nil {
		c.httpClient = options.restyClient
	}

	// 设置基础 URL
	c.httpClient.SetBaseURL(baseURL)

	// 设置 Debug
	if options.debug {
		c.httpClient.EnableDebug()
	}

	// 绑定日志钩子
	if options.restyLog != nil {
		// 请求中间件
		c.httpClient.SetRequestMiddlewares(
			options.restyLog.IntrusionRequest, // 自定义请求中间件，注入开始时间
			resty.PrepareRequestMiddleware,    // 官方请求中间件，创建 RawRequest
			options.restyLog.BeforeRequest,    // 自定义请求中间件，记录开始时间和OTel
		)
		// 响应中间件
		c.httpClient.SetResponseMiddlewares(
			options.restyLog.CopyResponseBodyMiddleware, // 自定义请求中间件，将响应体拷贝到 Context
			resty.AutoParseResponseMiddleware,           // 官方请求中间件，自动解析
			options.restyLog.AfterResponse,              // 自定义请求中间件，打印/保存
		)
	}

	return c, nil
}

// Close 关闭 请求客户端
func (c *Client) Close() {
	if c.httpClient != nil {
		c.httpClient.Close()
	}
}

type ErrResp struct {
	ErrorResponse struct {
		ErrorMsg  string `json:"error_msg"`
		SubMsg    string `json:"sub_msg"`
		SubCode   any    `json:"sub_code"`
		ErrorCode int    `json:"error_code"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}

type CustomParametersResult struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}
