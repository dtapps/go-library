package wechatminiprogram

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

const (
	WECHAT_API_URL = "https://api.weixin.qq.com"
	WECHAT_MP_URL  = "https://mp.weixin.qq.com"
	CGIUrl         = WECHAT_API_URL + "/cgi-bin"
	UnionUrl       = WECHAT_API_URL + "/union"
)

type ConfigClient struct {
	AppId       string            // 小程序唯一凭证，即 appId
	AppSecret   string            // 小程序唯一凭证密钥，即 appSecret
	AccessToken string            // 接口调用凭证
	JsapiTicket string            // 签名凭证
	RedisClient *dorm.RedisClient // 缓存数据库
	TokenDb     *gorm.DB          // 令牌数据库
	PgsqlDb     *gorm.DB          // pgsql数据库

}

// Client 微信小程序服务
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
		c.logTableName = "wechatminiprogram"
		c.log = golog.NewApi(&golog.ApiConfig{
			Db:        c.config.PgsqlDb,
			TableName: c.logTableName,
		})
	}

	return c
}

// ConfigApp 配置
func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.AppId = appId
	c.config.AppSecret = appSecret
	return c
}

// 请求接口
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

func (c *Client) GetAppId() string {
	return c.config.AppId
}
