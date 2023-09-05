package caiyunapp

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Token string
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		token string
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.token = config.Token

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
