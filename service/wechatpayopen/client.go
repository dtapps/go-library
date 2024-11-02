package wechatpayopen

import (
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	SpAppid        string // 服务商应用ID
	SpMchId        string // 服务商户号
	ApiV2          string // APIv2密钥
	ApiV3          string // APIv3密钥
	SerialNo       string // 序列号
	MchSslSerialNo string // pem 证书号
	MchSslCer      string // pem 内容
	MchSslKey      string // pem key 内容
}

// Client 实例
type Client struct {
	config struct {
		spAppid        string // 服务商应用ID
		spMchId        string // 服务商户号
		subAppid       string // 子商户应用ID
		subMchId       string // 子商户号
		apiV2          string // APIv2密钥
		apiV3          string // APIv3密钥
		serialNo       string // 序列号
		mchSslSerialNo string // pem 证书号
		mchSslCer      string // pem 内容
		mchSslKey      string // pem key 内容
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.spAppid = config.SpAppid
	c.config.spMchId = config.SpMchId
	c.config.apiV2 = config.ApiV2
	c.config.apiV3 = config.ApiV3
	c.config.serialNo = config.SerialNo
	c.config.mchSslSerialNo = config.MchSslSerialNo
	c.config.mchSslCer = config.MchSslCer
	c.config.mchSslKey = config.MchSslKey

	return c, nil
}
