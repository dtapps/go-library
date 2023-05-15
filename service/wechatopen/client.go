package wechatopen

import (
	"errors"
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
	ComponentAppId      string              `json:"component_app_id"`       // 第三方平台appid
	ComponentAppSecret  string              `json:"component_app_secret"`   // 第三方平台app_secret
	MessageToken        string              `json:"message_token"`          // 第三方平台消息令牌
	MessageKey          string              `json:"message_key"`            // 第三方平台消息密钥
	RedisClient         *dorm.RedisClient   `json:"redis_client"`           // 缓存数据库
	RedisCachePrefixFun redisCachePrefixFun `json:"redis_cache_prefix_fun"` // 缓存前缀
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		componentAppId         string // 第三方平台appid
		componentAppSecret     string // 第三方平台app_secret
		messageToken           string // 第三方平台消息令牌
		messageKey             string // 第三方平台消息密钥
		componentAccessToken   string // 第三方平台access_token
		componentVerifyTicket  string // 第三方平台推送ticket
		preAuthCode            string // 第三方平台预授权码
		authorizerAppid        string // 授权方appid
		authorizerAccessToken  string // 授权方access_token
		authorizerRefreshToken string // 授权方refresh_token
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
	zap struct {
		status bool             // 状态
		client *golog.ApiZapLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	if config.ComponentAppId == "" {
		return nil, errors.New("请配置 ComponentAppId")
	}
	c.config.componentAppId = config.ComponentAppId

	if config.ComponentAppSecret == "" {
		return nil, errors.New("请配置 ComponentAppSecret")
	}
	c.config.componentAppSecret = config.ComponentAppSecret

	if config.MessageToken == "" {
		return nil, errors.New("请配置 MessageToken")
	}
	c.config.messageToken = config.MessageToken

	if config.MessageKey == "" {
		return nil, errors.New("请配置 MessageKey")
	}
	c.config.messageKey = config.MessageKey

	c.requestClient = gorequest.NewHttp()

	c.cache.redisClient = config.RedisClient

	c.cache.componentVerifyTicketPrefix, c.cache.componentAccessTokenPrefix, c.cache.authorizerAccessTokenPrefix, c.cache.preAuthCodePrefix = config.RedisCachePrefixFun()
	if c.cache.componentVerifyTicketPrefix == "" || c.cache.componentAccessTokenPrefix == "" || c.cache.authorizerAccessTokenPrefix == "" || c.cache.preAuthCodePrefix == "" {
		return nil, errors.New("请配置 RedisCachePrefixFun")
	}

	return c, nil
}
