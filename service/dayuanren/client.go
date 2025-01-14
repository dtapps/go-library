package dayuanren

import (
	"errors"
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string // 接口地址
	UserID int64  // 商户ID
	ApiKey string // 秘钥
}

// Client 实例
type Client struct {
	config struct {
		apiURL string // 接口地址
		userID int64  // 商户ID
		apiKey string // 秘钥
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	if config.ApiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	c.httpClient = gorequest.NewHttp()

	c.config.apiURL = config.ApiURL
	c.config.userID = config.UserID
	c.config.apiKey = config.ApiKey

	return c, nil
}
