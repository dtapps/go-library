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
	client    *gorequest.App   // 请求客户端
	log       *golog.ApiClient // 日志服务
	logStatus bool             // 日志状态
	config    *ConfigClient    // 配置
}

// NewClient 实例化
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

// SubConfig 子商户配置
func (c *Client) SubConfig(subAppid, subMchId string) *Client {
	c.config.SpAppid = subAppid
	c.config.SubMchId = subMchId
	return c
}
