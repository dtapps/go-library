package jd

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppKey     string // 应用Key
	SecretKey  string // 密钥
	SiteId     string // 网站ID/APP ID
	PositionId string // 推广位id
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		appKey     string // 应用Key
		secretKey  string // 密钥
		siteId     string // 网站ID/APP ID
		positionId string // 推广位id
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appKey = config.AppKey
	c.config.secretKey = config.SecretKey
	c.config.siteId = config.SiteId
	c.config.positionId = config.PositionId

	c.requestClient = gorequest.NewHttp()
	c.requestClient.Uri = apiUrl

	return c, nil
}
