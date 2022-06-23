package wikeyun

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/goip"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	StoreId      int               // 店铺ID
	AppKey       int               // key
	AppSecret    string            // secret
	MongoDb      *dorm.MongoClient // 日志数据库
	PgsqlDb      *gorm.DB          // 日志数据库
	DatabaseName string            // 库名
}

type Client struct {
	client   *gorequest.App   // 请求客户端
	clientIp string           // Ip
	log      *golog.ApiClient // 日志服务
	config   *ConfigClient
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.client = gorequest.NewHttp()

	if c.config.PgsqlDb != nil {
		c.log, err = golog.NewApiClient(
			golog.WithGormClient(c.config.PgsqlDb),
			golog.WithTableName(logTable),
		)
		if err != nil {
			return nil, err
		}
	}
	if c.config.MongoDb != nil {
		c.log, err = golog.NewApiClient(
			golog.WithMongoClient(c.config.MongoDb),
			golog.WithDatabaseName(c.config.DatabaseName),
			golog.WithCollectionName(logTable),
		)
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
