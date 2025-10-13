package pinduoduo

import (
	"context"

	"resty.dev/v3"
)

const (
	Version = "1.0.62"
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
	client := resty.New()
	if options.restyClient != nil {
		client = options.restyClient
	}

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
