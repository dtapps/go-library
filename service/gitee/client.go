package gitee

import (
	"github.com/dtapps/go-library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	ClientID     string
	ClientSecret string
	RedirectUri  string
	PgsqlDb      *gorm.DB // 日志数据库
}

type Client struct {
	client *gorequest.App // 请求客户端
	config *ConfigClient  // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.client = gorequest.NewHttp()

	return c, nil
}
