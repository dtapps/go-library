package praise_goodness

import (
	"errors"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string // 接口地址
	MchID  int64
	Key    string
}

// Client 实例
type Client struct {
	config struct {
		apiURL string
		mchID  int64
		Key    string
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	if config.ApiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	c.httpClient = gorequest.NewHttp()

	c.config.apiURL = config.ApiURL
	c.config.mchID = config.MchID
	c.config.Key = config.Key

	c.trace = true
	return c, nil
}
