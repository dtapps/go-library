package wechatopen

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	ComponentAccessToken   string // 第三方平台 access_token
	ComponentVerifyTicket  string // 微信后台推送的 ticket
	PreAuthCode            string // 预授权码
	AuthorizerAccessToken  string // 接口调用令牌
	AuthorizerRefreshToken string // 刷新令牌
	AuthorizerAppid        string // 授权方 appid
	ComponentAppId         string // 第三方平台 appid
	ComponentAppSecret     string // 第三方平台 app_secret
	MessageToken           string
	MessageKey             string
	RedisClient            *dorm.RedisClient // 缓存数据库
	PgsqlDb                *gorm.DB          // pgsql数据库
}

// Client 微信公众号服务
type Client struct {
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
	config       *ConfigClient  // 配置
}

func NewClient(config *ConfigClient) *Client {

	c := &Client{config: config}

	c.client = gorequest.NewHttp()
	if c.config.PgsqlDb != nil {
		c.logStatus = true
		c.logTableName = "wechatopen"
		c.log = golog.NewApi(&golog.ApiConfig{
			Db:        c.config.PgsqlDb,
			TableName: c.logTableName,
		})
	}

	return c
}

// ConfigComponent 配置
func (c *Client) ConfigComponent(componentAppId, componentAppSecret string) *Client {
	c.config.ComponentAppId = componentAppId
	c.config.ComponentAppSecret = componentAppSecret
	return c
}

// ConfigAuthorizer 配置第三方
func (c *Client) ConfigAuthorizer(authorizerAppid string) *Client {
	c.config.AuthorizerAppid = authorizerAppid
	return c
}

func (c *Client) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Request()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.logStatus == true {
		go c.postgresqlLog(request)
	}

	return request, err
}
