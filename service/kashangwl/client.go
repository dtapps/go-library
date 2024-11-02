package kashangwl

import (
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	CustomerId  int64  // 商家编号
	CustomerKey string // 商家密钥
}

// Client 实例
type Client struct {
	config struct {
		customerId  int64  // 商家编号
		customerKey string // 商家密钥
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.customerId = config.CustomerId
	c.config.customerKey = config.CustomerKey

	return c, nil
}
