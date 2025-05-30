package wechatqy

import (
	"path/filepath"

	"go.dtapp.net/library/utils/resty_extend"
	"resty.dev/v3"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId       string
	AgentId     int
	Secret      string
	RedirectUri string

	Debug       bool   // 调试
	GlcStatus   bool   // 远程日志
	LogPath     string // 日志地址
	ServiceName string // 服务名称
}

// Client 实例
type Client struct {
	config struct {
		appId       string
		agentId     int
		secret      string
		redirectUri string
	}
	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = resty.New().SetDebug(config.Debug)
	if config.GlcStatus {
		c.httpClient.SetLogger(&resty_extend.GlcLogger{})
	} else {
		if config.LogPath != "" {
			c.httpClient.SetLogger(resty_extend.NewLog(filepath.Join(config.LogPath), config.ServiceName))
		}
	}

	c.config.appId = config.AppId
	c.config.agentId = config.AgentId
	c.config.secret = config.Secret
	c.config.redirectUri = config.RedirectUri

	return c, nil
}

// Close 关闭 请求客户端
func (c *Client) Close() {
	if c.httpClient != nil {
		c.httpClient.Close()
	}
}
