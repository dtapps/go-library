package caiyunapp

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// Client 实例
type Client struct {
	token      string
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(token string) (*Client, error) {
	return &Client{token: token, httpClient: gorequest.NewHttp()}, nil
}
