package wechatpayapiv2

import (
	"context"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL string // 接口地址

		appId      string // 小程序或者公众号唯一凭证
		appSecret  string // 小程序或者公众号唯一凭证密钥
		mchId      string // 微信支付的商户id
		mchKey     string // 私钥
		certString string
		keyString  string
	}

	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.baseURL = "https://api.mch.weixin.qq.com"
	if options.baseURL != "" {
		c.config.baseURL = options.baseURL
	}
	c.config.appId = options.appId
	c.config.appSecret = options.appSecret
	c.config.mchId = options.mchId
	c.config.mchKey = options.mchKey
	c.config.certString = options.certString
	c.config.keyString = options.keyString

	// 创建请求客户端
	c.httpClient = resty.New()
	if options.restyClient != nil {
		c.httpClient = options.restyClient
	}

	// 设置基础 URL
	c.httpClient.SetBaseURL(c.config.baseURL)

	// 设置 Debug
	if options.debug {
		c.httpClient.EnableDebug()
	}

	return c, nil
}

// Close 关闭 请求客户端
func (c *Client) Close() (err error) {
	if c.httpClient != nil {
		err = c.httpClient.Close()
	}
	return
}
