package cloudflare

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	AccountID    string
	ZoneID       string
	GlobalAPIKey string
	MongoDb      *dorm.MongoClient // 日志数据库
	PgsqlDb      *gorm.DB          // 日志数据库
	DatabaseName string            // 库名
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
