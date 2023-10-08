package praise_goodness

import (
	"errors"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string
	MchID  int64
	Key    string
}

// Client 实例
type Client struct {
	requestClient       *gorequest.App // 请求服务
	requestClientStatus bool           // 请求服务状态
	config              struct {
		apiURL string
		mchID  int64
		Key    string
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
	c.config.mchID = config.MchID
	c.config.Key = config.Key

	if c.config.apiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	return c, nil
}
