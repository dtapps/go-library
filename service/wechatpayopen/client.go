package wechatpayopen

import (
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

type ConfigClient struct {
	SpAppid        string   // 服务商应用ID
	SpMchId        string   // 服务商户号
	SubAppid       string   // 子商户应用ID
	SubMchId       string   // 子商户号
	ApiV2          string   // APIv2密钥
	ApiV3          string   // APIv3密钥
	SerialNo       string   // 序列号
	MchSslSerialNo string   // pem 证书号
	MchSslCer      string   // pem 内容
	MchSslKey      string   // pem key 内容
	PgsqlDb        *gorm.DB // pgsql数据库
}

// Client 微信支付服务
type Client struct {
	client       *gorequest.App // 请求客户端
	log          *golog.Api     // 日志服务
	logTableName string         // 日志表名
	logStatus    bool           // 日志状态
	config       *ConfigClient  // 配置
}

// NewClient 实例化
func NewClient(config *ConfigClient) *Client {

	c := &Client{config: config}

	c.client = gorequest.NewHttp()
	if c.config.PgsqlDb != nil {
		c.logStatus = true
		c.logTableName = "wechatpayopen"
		c.log = golog.NewApi(&golog.ApiConfig{
			Db:        c.config.PgsqlDb,
			TableName: c.logTableName,
		})
	}

	return c
}

// SubConfig 子商户配置
func (c *Client) SubConfig(subAppid, subMchId string) *Client {
	c.config.SpAppid = subAppid
	c.config.SubMchId = subMchId
	return c
}

func (c *Client) request(url string, params map[string]interface{}, method string) (resp gorequest.Response, err error) {

	// 认证
	authorization, err := c.authorization(method, params, url)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置JSON格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 设置头部
	client.SetHeader("Authorization", authorization)
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Accept-Language", "zh-CN")

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
