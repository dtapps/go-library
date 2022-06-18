package wechatunion

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
	Pid         string            // 推广位PID
	RedisClient *dorm.RedisClient // 缓存数据库
	PgsqlDb     *gorm.DB          // pgsql数据库
}

// Client 微信小程序联盟
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
