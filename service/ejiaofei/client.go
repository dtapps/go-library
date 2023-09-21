package ejiaofei

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	UserId string
	Pwd    string
	Key    string
}

// Client 实例
type Client struct {
	requestClient       *gorequest.App // 请求服务
	requestClientStatus bool           // 请求服务状态
	config              struct {
		userId  string
		pwd     string
		key     string
		signStr string // 需要签名的字符串
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.userId = config.UserId
	c.config.pwd = config.Pwd
	c.config.key = config.Key

	return c, nil
}
