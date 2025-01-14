package cma

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

type Client struct {
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

func NewClient() (*Client, error) {
	return &Client{httpClient: gorequest.NewHttp()}, nil
}
