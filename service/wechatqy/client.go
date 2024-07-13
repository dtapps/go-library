package wechatqy

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId       string
	AgentId     int
	Secret      string
	RedirectUri string
}

// Client 实例
type Client struct {
	config struct {
		appId       string
		agentId     int
		secret      string
		redirectUri string
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

	c.config.appId = config.AppId
	c.config.agentId = config.AgentId
	c.config.secret = config.Secret
	c.config.redirectUri = config.RedirectUri

	c.trace = true
	return c, nil
}
