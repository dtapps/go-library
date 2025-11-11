package wechatopen

import (
	"context"
	"time"

	"resty.dev/v3"
)

// Client 实例
type Client struct {
	config struct {
		baseURL string // 接口地址

		componentAppId                  string    // 第三方平台appid
		componentAppSecret              string    // 第三方平台app_secret
		messageToken                    string    // 第三方平台消息令牌
		messageKey                      string    // 第三方平台消息密钥
		componentAccessToken            string    // 第三方平台access_token
		componentAccessTokenUpdateTime  time.Time // 第三方平台access_token 更新时间
		componentVerifyTicket           string    // 第三方平台推送ticket
		componentVerifyTicketUpdateTime time.Time // 第三方平台推送ticket 更新时间
		componentPreAuthCode            string    // 第三方平台预授权码
		componentPreAuthCodeUpdateTime  time.Time // 第三方平台预授权码 更新时间

		authorizerAppid                 string    // 授权方appid
		authorizerAccessToken           string    // 授权方access_token
		authorizerAccessTokenUpdateTime time.Time // 授权方access_token 更新时间
		authorizerRefreshToken          string    // 授权方refresh_token
		authorizerReleaseVersion        string    // 授权方release_version
	}

	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := NewOptions(opts)

	c := &Client{}
	c.config.baseURL = "https://api.weixin.qq.com/"
	if options.baseURL != "" {
		c.config.baseURL = options.baseURL
	}
	c.config.componentAppId = options.componentAppId
	c.config.componentAppSecret = options.componentAppSecret
	c.config.messageToken = options.messageToken
	c.config.messageKey = options.messageKey
	c.config.componentAccessToken = options.componentAccessToken
	c.config.componentAccessTokenUpdateTime = options.componentAccessTokenUpdateTime
	c.config.componentVerifyTicket = options.componentVerifyTicket
	c.config.componentVerifyTicketUpdateTime = options.componentVerifyTicketUpdateTime
	c.config.componentPreAuthCode = options.componentPreAuthCode
	c.config.componentPreAuthCodeUpdateTime = options.componentPreAuthCodeUpdateTime

	c.config.authorizerAppid = options.authorizerAppid
	c.config.authorizerAccessToken = options.authorizerAccessToken
	c.config.authorizerAccessTokenUpdateTime = options.authorizerAccessTokenUpdateTime
	c.config.authorizerRefreshToken = options.authorizerRefreshToken

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
