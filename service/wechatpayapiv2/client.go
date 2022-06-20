package wechatpayapiv2

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	AppId      string // 小程序或者公众号唯一凭证
	AppSecret  string // 小程序或者公众号唯一凭证密钥
	MchId      string // 微信支付的商户id
	MchKey     string // 私钥
	CertString string
	KeyString  string
	MongoDb    *dorm.MongoClient // 日志数据库
	PgsqlDb    *gorm.DB          // pgsql数据库
}

// Client 微信支付服务
type Client struct {
	client    *gorequest.App   // 请求客户端
	log       *golog.ApiClient // 日志服务
	logStatus bool             // 日志状态
	config    *ConfigClient    // 配置
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

	return c, nil
}
