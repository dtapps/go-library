package wechatpayapiv2

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId      string `json:"app_id"` // 小程序或者公众号唯一凭证
	AppSecret  string // 小程序或者公众号唯一凭证密钥
	MchId      string `json:"mch_id"` // 微信支付的商户id
	MchKey     string // 私钥
	CertString string
	KeyString  string
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		appId      string // 小程序或者公众号唯一凭证
		appSecret  string // 小程序或者公众号唯一凭证密钥
		mchId      string // 微信支付的商户id
		mchKey     string // 私钥
		certString string
		keyString  string
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.mchId = config.MchId
	c.config.mchKey = config.MchKey
	c.config.certString = config.CertString
	c.config.keyString = config.KeyString

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
