package wikeyun

import (
	"fmt"
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
	client       *gorequest.App // 请求客户端
	clientIp     string         // Ip
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
	config       *ConfigClient
}

func NewClient(config *ConfigClient) *Client {

	c := &Client{config: config}
	c.config = config

	c.client = gorequest.NewHttp()
	if c.config.PgsqlDb != nil {
		c.logStatus = true
		c.logTableName = "wikeyun"
		c.log = golog.NewApi(&golog.ApiConfig{
			Db:        c.config.PgsqlDb,
			TableName: c.logTableName,
		})
	}
	xip := goip.GetOutsideIp()
	if xip != "" && xip != "0.0.0.0" {
		c.clientIp = xip
	}

	return c
}

// 请求接口
func (c *Client) request(url string, params map[string]interface{}) (resp gorequest.Response, err error) {

	// 签名
	sign := c.sign(params)

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(fmt.Sprintf("%s?app_key=%d&timestamp=%s&client=%s&format=%s&v=%s&sign=%s", url, c.config.AppKey, sign.Timestamp, sign.Client, sign.Format, sign.V, sign.Sign))

	// 设置FORM格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.logStatus == true {
		go c.postgresqlLog(request)
	}

	return request, err
}
