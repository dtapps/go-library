package wechatopen

import (
	"errors"
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ComponentAppId     string `json:"component_app_id"`     // 第三方平台appid
	ComponentAppSecret string `json:"component_app_secret"` // 第三方平台app_secret
	MessageToken       string `json:"message_token"`        // 第三方平台消息令牌
	MessageKey         string `json:"message_key"`          // 第三方平台消息密钥
}

// Client 实例
type Client struct {
	config struct {
		componentAppId         string // 第三方平台appid
		componentAppSecret     string // 第三方平台app_secret
		messageToken           string // 第三方平台消息令牌
		messageKey             string // 第三方平台消息密钥
		componentAccessToken   string // 第三方平台access_token
		componentVerifyTicket  string // 第三方平台推送ticket
		componentPreAuthCode   string // 第三方平台预授权码
		authorizerAppid        string // 授权方appid
		authorizerAccessToken  string // 授权方access_token
		authorizerRefreshToken string // 授权方refresh_token
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	c.httpClient = gorequest.NewHttp()

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

	return c, nil
}
