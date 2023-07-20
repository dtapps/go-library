package kuaidi100

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Customer string // 授权码
	Key      string // 密钥
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		customer string // 授权码
		key      string // 密钥
	}
	zap struct {
		status bool             // 状态
		client *golog.ApiZapLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.customer = config.Customer

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
