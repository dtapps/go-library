package kuaidi100

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Customer string // 授权码
	Key      string // 密钥
}

// Client 实例
type Client struct {
	config struct {
		customer string // 授权码
		key      string // 密钥
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.customer = config.Customer
	c.config.key = config.Key

	c.trace = true
	return c, nil
}
