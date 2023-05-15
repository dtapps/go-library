package dingdanxia

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiKey string // 密钥
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		apiKey string // 密钥
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
	zap struct {
		status bool             // 状态
		client *golog.ApiZapLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.apiKey = config.ApiKey

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
