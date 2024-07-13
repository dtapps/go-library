package kashangwl

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	CustomerId  int64  // 商家编号
	CustomerKey string // 商家密钥
}

// Client 实例
type Client struct {
	config struct {
		customerId  int64  // 商家编号
		customerKey string // 商家密钥
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

	c.config.customerId = config.CustomerId
	c.config.customerKey = config.CustomerKey

	c.trace = true
	return c, nil
}
