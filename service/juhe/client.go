package juhe

import (
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
}

// Client 实例
type Client struct {
	config struct {
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient() (*Client, error) {
	c := &Client{}
	c.httpClient = gorequest.NewHttp()

	return c, nil
}
