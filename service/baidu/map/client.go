package _map

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

type Client struct {
	ak         string
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(ctx context.Context, ak string) (*Client, error) {
	return &Client{ak: ak, httpClient: gorequest.NewHttp()}, nil
}
