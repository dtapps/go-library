package aswzk

import (
	"errors"
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiUrl string // 接口地址
	UserID string // 用户编号
	ApiKey string // 秘钥
}

// Client 实例
type Client struct {
	config struct {
		apiUrl string // 接口地址
		userID string // 用户编号
		apiKey string // 秘钥
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	if config.ApiUrl == "" {
		return nil, errors.New("ApiUrl is empty")
	}

	c.httpClient = gorequest.NewHttp()

	c.config.apiUrl = config.ApiUrl
	c.config.userID = config.UserID
	c.config.apiKey = config.ApiKey

	return c, nil
}
