package dingdanxia

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	ApiKey  string   // 密钥
	PgsqlDb *gorm.DB // pgsql数据库
}

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
		c.logTableName = "dingdanxia"
		c.log = golog.NewApi(&golog.ApiConfig{
			Db:        c.config.PgsqlDb,
			TableName: c.logTableName,
		})
	}

	return c
}

// 请求接口
func (c *Client) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 公共参数
	params["apikey"] = c.config.ApiKey

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

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
