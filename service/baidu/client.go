package baidu

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Ak string
}
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		ak string
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.ak = config.Ak

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
