package meituan

import (
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Secret string // 秘钥
	AppKey string // 渠道标记
}

// Client 实例
type Client struct {
	config struct {
		secret string // 秘钥
		appKey string // 渠道标记
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.secret = config.Secret
	c.config.appKey = config.AppKey

	return c, nil
}
