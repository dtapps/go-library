package wechatopen

import "context"

// NewSetComponentAccessToken 设置第三方平台access_token
func (c *Client) NewSetComponentAccessToken(ctx context.Context, componentAccessToken string) {
	c.config.componentAccessToken = componentAccessToken
}

// NewSetComponentVerifyTicket 设置第三方平台推送ticket
func (c *Client) NewSetComponentVerifyTicket(ctx context.Context, componentVerifyTicket string) {
	c.config.componentVerifyTicket = componentVerifyTicket
}

// NewSetPreAuthCode 设置第三方平台预授权码
func (c *Client) NewSetPreAuthCode(ctx context.Context, preAuthCode string) {
	c.config.preAuthCode = preAuthCode
}

// NewSetAuthorizerAppid 设置授权方appid
func (c *Client) NewSetAuthorizerAppid(ctx context.Context, authorizerAppid string) {
	c.config.authorizerAppid = authorizerAppid
}

// NewSetAuthorizerAccessToken 设置授权方access_token
func (c *Client) NewSetAuthorizerAccessToken(ctx context.Context, authorizerAccessToken string) {
	c.config.authorizerAccessToken = authorizerAccessToken
}

// NewSetAuthorizerRefreshToken 设置授权方refresh_token
func (c *Client) NewSetAuthorizerRefreshToken(ctx context.Context, authorizerRefreshToken string) {
	c.config.authorizerRefreshToken = authorizerRefreshToken
}
