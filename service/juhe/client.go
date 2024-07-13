package juhe

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
}

// Client 实例
type Client struct {
	config struct {
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient() (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.trace = true
	return c, nil
}
