package ejiaofei

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	UserId string
	Pwd    string
	Key    string
}

// Client 实例
type Client struct {
	config struct {
		userId  string
		pwd     string
		key     string
		signStr string // 需要签名的字符串
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

	c.config.userId = config.UserId
	c.config.pwd = config.Pwd
	c.config.key = config.Key

	c.trace = true
	return c, nil
}
