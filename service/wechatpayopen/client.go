package wechatpayopen

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	SpAppid        string            // 服务商应用ID
	SpMchId        string            // 服务商户号
	SubAppid       string            // 子商户应用ID
	SubMchId       string            // 子商户号
	ApiV2          string            // APIv2密钥
	ApiV3          string            // APIv3密钥
	SerialNo       string            // 序列号
	MchSslSerialNo string            // pem 证书号
	MchSslCer      string            // pem 内容
	MchSslKey      string            // pem key 内容
	MongoDb        *dorm.MongoClient // 日志数据库
	PgsqlDb        *gorm.DB          // 日志数据库
	DatabaseName   string            // 库名
}

// Client 微信支付服务
type Client struct {
	client *gorequest.App   // 请求客户端
	log    *golog.ApiClient // 日志服务
	config *ConfigClient    // 配置
}

// NewClient 实例化
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

	return c, nil
}
