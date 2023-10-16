package wnfuwu

import (
	"errors"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string // 接口地址
	UserID int64  // 商户ID
	ApiKey string // 秘钥
}

// Client 实例
type Client struct {
	requestClient       *gorequest.App // 请求服务
	requestClientStatus bool           // 请求服务状态
	config              struct {
		apiURL string // 接口地址
		userID int64  // 商户ID
		apiKey string // 秘钥
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.apiURL = config.ApiURL
	c.config.userID = config.UserID
	c.config.apiKey = config.ApiKey

	if c.config.apiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	return c, nil
}
