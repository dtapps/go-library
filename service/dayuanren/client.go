package dayuanren

import (
	"context"

	"resty.dev/v3"
)

const (
	Version = "1.0.31"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string // 接口地址
	UserID int64  // 商户ID
	ApiKey string // 秘钥

	Debug       bool   // 调试
	GlcStatus   bool   // 远程日志
	LogPath     string // 日志地址
	ServiceName string // 服务名称
}

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
	client := resty.New()
	if options.restyClient != nil {
		client = options.restyClient
	}

	// 设置基础 URL
	client.SetBaseURL(options.baseURL)

	// 设置 Debug
	if options.debug {
		client.EnableDebug()
	}

	// 绑定日志钩子
	if options.restyLog != nil {
		// 请求中间件
		c.httpClient.SetRequestMiddlewares(
			resty.PrepareRequestMiddleware, // 必须放第一，用于生成原始 http.Request（RawRequest），
			options.restyLog.BeforeRequest, // 自定义请求中间件，记录请求开始时间、可做日志记录或其他请求预处理
		)
		// 响应中间件
		c.httpClient.SetResponseMiddlewares(
			options.restyLog.CopyResponseBodyMiddleware, // 放在 AutoParse 前，备份 Body
			resty.AutoParseResponseMiddleware,           // Resty 自动解析 JSON
			options.restyLog.AfterResponse,              // 最后打印 / 保存
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
