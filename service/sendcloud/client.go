package sendcloud

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

type ConfigClient struct {
	ApiUser string // API_USER
	ApiKey  string // API_KEY
}

type Client struct {
	client *gorequest.App // 请求服务
	config struct {
		apiUser string // API_USER
		apiKey  string // API_KEY
	}
	zap struct {
		status bool             // 状态
		client *golog.ApiZapLog // 日志服务
	}
}

func NewClient(config *ConfigClient) (*Client, error) {

	c := &Client{}

	c.config.apiUser = config.ApiUser
	c.config.apiKey = config.ApiKey

	c.client = gorequest.NewHttp()

	return c, nil
}
