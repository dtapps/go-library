package wechatopen

import (
	"context"
)

// GetComponentAppId 第三方平台appid
func (c *Client) GetComponentAppId(ctx context.Context) string {
	return c.config.componentAppId
}

// GetComponentAppSecret 第三方平台app_secret
func (c *Client) GetComponentAppSecret(ctx context.Context) string {
	return c.config.componentAppSecret
}

// GetMessageToken 第三方平台消息令牌
func (c *Client) GetMessageToken(ctx context.Context) string {
	return c.config.messageToken
}

// GetMessageKey 第三方平台消息密钥
func (c *Client) GetMessageKey(ctx context.Context) string {
	return c.config.messageKey
}

// GetAuthorizerAppid 授权方appid
func (c *Client) GetAuthorizerAppid(ctx context.Context) string {
	return c.config.authorizerAppid
}
