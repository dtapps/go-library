package wechatqy

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId       string
	AgentId     int
	Secret      string
	RedirectUri string
	Key         string // key
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		appId       string
		agentId     int
		secret      string
		redirectUri string
		key         string
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
	c.config.agentId = config.AgentId
	c.config.secret = config.Secret
	c.config.redirectUri = config.RedirectUri
	c.config.key = config.Key

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
