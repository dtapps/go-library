package wechatpayopen

import (
	"context"
	"crypto/rsa"
	"crypto/x509"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL string // 接口地址

		spAppid  string // 服务商应用ID
		spMchId  string // 服务商户号
		subAppid string // 子商户应用ID
		subMchId string // 子商户号
		apiV3    string // APIv3密钥

		certificateSerialNo string            // 证书序列号
		certificate         *x509.Certificate // pem 证书
		privateKey          *rsa.PrivateKey   // pem 私钥
		publicKeyID         string            // 公钥ID
		publicKey           *rsa.PublicKey    // pem 公钥
	}

	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)
	if options.err != nil {
		return nil, options.err
	}

	c := &Client{}
	c.config.baseURL = "https://api.mch.weixin.qq.com"
	// c.config.baseURL = "https://api2.mch.weixin.qq.com"
	if options.baseURL != "" {
		c.config.baseURL = options.baseURL
	}
	c.config.spAppid = options.spAppid
	c.config.spMchId = options.spMchId
	c.config.subAppid = options.subAppid
	c.config.subMchId = options.subMchId
	c.config.apiV3 = options.apiV3
	c.config.certificateSerialNo = options.certificateSerialNo

	c.config.certificate = options.certificate
	c.config.privateKey = options.privateKey
	c.config.publicKeyID = options.publicKeyID
	c.config.publicKey = options.publicKey

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
