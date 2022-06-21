package wechatminiprogram

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	AppId       string            // 小程序唯一凭证，即 appId
	AppSecret   string            // 小程序唯一凭证密钥，即 appSecret
	AccessToken string            // 接口调用凭证
	JsapiTicket string            // 签名凭证
	RedisClient *dorm.RedisClient // 缓存数据库
	TokenDb     *gorm.DB          // 令牌数据库
	MongoDb     *dorm.MongoClient // 日志数据库
	PgsqlDb     *gorm.DB          // 日志数据库

}

// Client 微信小程序服务
type Client struct {
	client *gorequest.App   // 请求客户端
	log    *golog.ApiClient // 日志服务
	config *ConfigClient    // 配置
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
			golog.WithMongoCollectionClient(c.config.MongoDb),
			golog.WithTableName(logTable),
		)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// ConfigApp 配置
func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.AppId = appId
	c.config.AppSecret = appSecret
	return c
}
