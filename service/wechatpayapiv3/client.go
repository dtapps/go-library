package wechatpayapiv3

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId          string // 小程序或者公众号唯一凭证
	AppSecret      string // 小程序或者公众号唯一凭证密钥
	MchId          string // 微信支付的商户id
	AesKey         string // 私钥
	ApiV3          string // API v3密钥
	MchSslSerialNo string // pem 证书号
	MchSslKey      string // pem key 内容
}

// Client 实例
type Client struct {
	config struct {
		appId          string // 小程序或者公众号唯一凭证
		appSecret      string // 小程序或者公众号唯一凭证密钥
		mchId          string // 微信支付的商户id
		aesKey         string // 私钥
		apiV3          string // API v3密钥
		mchSslSerialNo string // pem 证书号
		mchSslKey      string // pem key 内容
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	trace      bool           // OpenTelemetry链路追踪
	span       trace.Span     // OpenTelemetry链路追踪
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.mchId = config.MchId
	c.config.aesKey = config.AesKey
	c.config.apiV3 = config.ApiV3
	c.config.mchSslSerialNo = config.MchSslSerialNo
	c.config.mchSslKey = config.MchSslKey

	c.trace = true
	return c, nil
}
