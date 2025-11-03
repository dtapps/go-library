package wechatpayopen

import (
	"context"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL        string // 接口地址
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

	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

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
	c.config.apiV2 = options.apiV2
	c.config.apiV3 = options.apiV3
	c.config.serialNo = options.serialNo
	c.config.mchSslSerialNo = options.mchSslSerialNo
	c.config.mchSslCer = options.mchSslCer
	c.config.mchSslKey = options.mchSslKey

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

	// 绑定日志钩子
	if options.restyLog != nil {
		// 请求中间件
		c.httpClient.SetRequestMiddlewares(
			options.restyLog.IntrusionRequest, // 自定义请求中间件，注入开始时间
			resty.PrepareRequestMiddleware,    // 官方请求中间件，创建RawRequest
			options.restyLog.BeforeRequest,    // 自定义请求中间件，记录开始时间和OTel
		)
		// 响应中间件
		c.httpClient.SetResponseMiddlewares(
			options.restyLog.CopyResponseBodyMiddleware, // 自定义请求中间件，将响应体拷贝到Context
			resty.AutoParseResponseMiddleware,           // 官方请求中间件，自动解析
			options.restyLog.AfterResponse,              // 自定义请求中间件，打印/保存
		)
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
