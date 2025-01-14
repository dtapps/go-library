package wechatpayapiv2

import (
	"go.dtapp.net/library/utils/gorequest"
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
	config struct {
		appId      string // 小程序或者公众号唯一凭证
		appSecret  string // 小程序或者公众号唯一凭证密钥
		mchId      string // 微信支付的商户id
		mchKey     string // 私钥
		certString string
		keyString  string
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.mchId = config.MchId
	c.config.mchKey = config.MchKey
	c.config.certString = config.CertString
	c.config.keyString = config.KeyString

	return c, nil
}
