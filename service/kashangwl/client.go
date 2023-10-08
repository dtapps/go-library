package kashangwl

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	CustomerId     int64  // 商家编号
	CustomerKey    string // 商家密钥
	CacheLogStatus bool   // 缓存日志状态
}

// Client 实例
type Client struct {
	requestClient       *gorequest.App // 请求服务
	requestClientStatus bool           // 请求服务状态
	config              struct {
		customerId  int64  // 商家编号
		customerKey string // 商家密钥
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.customerId = config.CustomerId
	c.config.customerKey = config.CustomerKey

	return c, nil
}
