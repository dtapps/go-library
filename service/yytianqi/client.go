package yytianqi

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

type Client struct {
	key        string
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

func NewClient(key string) (*Client, error) {
	return &Client{key: key, httpClient: gorequest.NewHttp()}, nil
}
