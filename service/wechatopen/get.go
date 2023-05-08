package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/golog"
)

// GetComponentAppId 获取第三方平台appid
func (c *Client) GetComponentAppId(ctx context.Context) string {
	return c.config.componentAppId
}

// NewGetComponentAppId 获取第三方平台appid
func (c *Client) NewGetComponentAppId(ctx context.Context) string {
	return c.config.componentAppId
}

// GetComponentAppSecret 获取第三方平台app_secret
func (c *Client) GetComponentAppSecret(ctx context.Context) string {
	return c.config.componentAppSecret
}

// NewGetComponentAppSecret 获取第三方平台app_secret
func (c *Client) NewGetComponentAppSecret(ctx context.Context) string {
	return c.config.componentAppSecret
}

// GetMessageToken 获取第三方平台消息令牌
func (c *Client) GetMessageToken(ctx context.Context) string {
	return c.config.messageToken
}

// NewGetMessageToken 获取第三方平台消息令牌
func (c *Client) NewGetMessageToken(ctx context.Context) string {
	return c.config.messageToken
}

// GetMessageKey 获取第三方平台消息密钥
func (c *Client) GetMessageKey(ctx context.Context) string {
	return c.config.messageKey
}

// NewGetMessageKey 获取第三方平台消息密钥
func (c *Client) NewGetMessageKey(ctx context.Context) string {
	return c.config.messageKey
}

// NewGetComponentAccessToken 获取第三方平台access_token
func (c *Client) NewGetComponentAccessToken(ctx context.Context) string {
	return c.config.componentAccessToken
}

// GetComponentVerifyTicket 获取第三方平台推送ticket
func (c *Client) GetComponentVerifyTicket(ctx context.Context) string {
	return c.config.componentVerifyTicket
}

// NewGetComponentVerifyTicket 获取第三方平台推送ticket
func (c *Client) NewGetComponentVerifyTicket(ctx context.Context) string {
	return c.config.componentVerifyTicket
}

// NewGetPreAuthCode 获取第三方平台预授权码
func (c *Client) NewGetPreAuthCode(ctx context.Context) string {
	return c.config.preAuthCode
}

// GetAuthorizerAppid 获取授权方appid
func (c *Client) GetAuthorizerAppid(ctx context.Context) string {
	return c.config.authorizerAppid
}

// NewGetAuthorizerAppid 获取授权方appid
func (c *Client) NewGetAuthorizerAppid(ctx context.Context) string {
	return c.config.authorizerAppid
}

// NewGetAuthorizerAccessToken 获取授权方access_token
func (c *Client) NewGetAuthorizerAccessToken(ctx context.Context) string {
	return c.config.authorizerAccessToken
}

// NewGetAuthorizerRefreshToken 获取授权方refresh_token
func (c *Client) NewGetAuthorizerRefreshToken(ctx context.Context) string {
	return c.config.authorizerRefreshToken
}

func (c *Client) GetLog(ctx context.Context) *golog.ApiClient {
	return c.log.client
}
