package wechatopen

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// 缓存前缀
// wechat_open:component_verify_ticket:
// wechat_open:component_access_token:
// wechat_open:authorizer_access_token:
// wechat_open:pre_auth_code:
type redisCachePrefixFun func() (componentVerifyTicket, componentAccessToken, authorizerAccessToken, preAuthCode string)

// ClientConfig 实例配置
type ClientConfig struct {
	AuthorizerAppid     string `json:"authorizer_appid"` // 授权方 appid
	ComponentAppId      string `json:"component_app_id"` // 第三方平台 appid
	ComponentAppSecret  string // 第三方平台 app_secret
	MessageToken        string
	MessageKey          string
	RedisClient         *dorm.RedisClient   // 缓存数据库
	RedisCachePrefixFun redisCachePrefixFun // 缓存前缀
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		componentAccessToken   string // 第三方平台 access_token
		componentVerifyTicket  string // 微信后台推送的 ticket
		preAuthCode            string // 预授权码
		authorizerAccessToken  string // 接口调用令牌
		authorizerRefreshToken string // 刷新令牌
		authorizerAppid        string // 授权方 appid
		componentAppId         string // 第三方平台appid
		componentAppSecret     string // 第三方平台app_secret
		messageToken           string
		messageKey             string
	}
	cache struct {
		redisClient                 *dorm.RedisClient // 缓存数据库
		componentVerifyTicketPrefix string
		componentAccessTokenPrefix  string
		authorizerAccessTokenPrefix string
		preAuthCodePrefix           string
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.componentAppId = config.ComponentAppId
	c.config.componentAppSecret = config.ComponentAppSecret
	c.config.messageToken = config.MessageToken
	c.config.messageKey = config.MessageKey

	c.requestClient = gorequest.NewHttp()

	c.cache.redisClient = config.RedisClient

	c.cache.componentVerifyTicketPrefix, c.cache.componentAccessTokenPrefix, c.cache.authorizerAccessTokenPrefix, c.cache.preAuthCodePrefix = config.RedisCachePrefixFun()
	if c.cache.componentVerifyTicketPrefix == "" || c.cache.componentAccessTokenPrefix == "" || c.cache.authorizerAccessTokenPrefix == "" || c.cache.preAuthCodePrefix == "" {
		return nil, redisCachePrefixNoConfig
	}

	return c, nil
}
