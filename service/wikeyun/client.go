package wikeyun

import (
	"go.dtapp.net/library/utils/goip"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	StoreId   int      // 店铺ID
	AppKey    int      // key
	AppSecret string   // secret
	PgsqlDb   *gorm.DB // pgsql数据库
}

type Client struct {
	client    *gorequest.App   // 请求客户端
	clientIp  string           // Ip
	log       *golog.ApiClient // 日志服务
	logStatus bool             // 日志状态
	config    *ConfigClient
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.client = gorequest.NewHttp()
	if c.config.PgsqlDb != nil {
		c.logStatus = true
		c.log, err = golog.NewApiClient(&golog.ConfigApiClient{
			Db:        c.config.PgsqlDb,
			TableName: logTable,
		})
		if err != nil {
			return nil, err
		}
	}
	xip := goip.GetOutsideIp()
	if xip != "" && xip != "0.0.0.0" {
		c.clientIp = xip
	}

	return c, nil
}
