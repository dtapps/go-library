package chengquan

import (
	"errors"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string
	AppID  string
	AppKey string
	AesKey string
	AesIv  string
}

// Client 实例
type Client struct {
	config struct {
		apiURL  string
		appID   string
		appKey  string
		aesKey  string
		aesIv   string
		version string
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
	c.config.appID = config.AppID
	c.config.appKey = config.AppKey
	c.config.aesKey = config.AesKey
	c.config.aesIv = config.AesKey
	c.config.version = "1.0.0"

	c.trace = true
	return c, nil
}
