package meituan

import (
	"go.dtapp.net/library/utils/gorequest"
)

// MediaClientConfig 实例配置
type MediaClientConfig struct {
	AppKey    string // 分配的AppKey
	AppSecret string // 分配的AppSecret
}

// MediaClient 实例
type MediaClient struct {
	config struct {
		appKey    string // 分配的AppKey
		appSecret string // 分配的AppSecret
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewMediaClient 创建实例化
func NewMediaClient(config *MediaClientConfig) *MediaClient {
	c := &MediaClient{}

	c.httpClient = gorequest.NewHttp()

	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	return c
}
