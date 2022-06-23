package kuaishou

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	MongoDb *dorm.MongoClient // 日志数据库
	PgsqlDb *gorm.DB          // 日志数据库
}

type Client struct {
	ua     string
	client *gorequest.App   // 请求客户端
	log    *golog.ApiClient // 日志服务
	config *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"

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
			golog.WithMongoCollectionClient(c.config.MongoDb),
			golog.WithTableName(logTable),
		)
		if err != nil {
			return nil, err
		}
	}

	//if c.config.PgsqlDb != nil {
	//	go c.log.GormMiddleware(request)
	//}
	//if c.config.MongoDb != nil {
	//	go c.log.MongoMiddleware(request)
	//}

	return c, nil
}
