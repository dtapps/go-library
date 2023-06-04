package eastiot

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId  string
	ApiKey string
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		appId  string
		apiKey string
	}
	zap struct {
		status bool             // 状态
		client *golog.ApiZapLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appId = config.AppId
	c.config.apiKey = config.ApiKey

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
