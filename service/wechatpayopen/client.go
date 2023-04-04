package wechatpayopen

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	SpAppid        string `json:"sp_appid"`  // 服务商应用ID
	SpMchId        string `json:"sp_mch_id"` // 服务商户号
	ApiV2          string // APIv2密钥
	ApiV3          string // APIv3密钥
	SerialNo       string // 序列号
	MchSslSerialNo string // pem 证书号
	MchSslCer      string // pem 内容
	MchSslKey      string // pem key 内容
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
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
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.spAppid = config.SpAppid
	c.config.spMchId = config.SpMchId
	c.config.apiV2 = config.ApiV2
	c.config.apiV3 = config.ApiV3
	c.config.serialNo = config.SerialNo
	c.config.mchSslSerialNo = config.MchSslSerialNo
	c.config.mchSslCer = config.MchSslCer
	c.config.mchSslKey = config.MchSslKey

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
