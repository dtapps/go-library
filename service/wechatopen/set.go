package wechatopen

import "context"

// SetComponentAccessToken 设置第三方平台access_token
func (c *Client) SetComponentAccessToken(ctx context.Context, componentAccessToken string) {
	c.config.componentAccessToken = componentAccessToken
}

// SetComponentVerifyTicket 设置第三方平台推送ticket
func (c *Client) SetComponentVerifyTicket(ctx context.Context, componentVerifyTicket string) {
	c.config.componentVerifyTicket = componentVerifyTicket
}

// SetPreAuthCode 设置第三方平台预授权码
func (c *Client) SetPreAuthCode(ctx context.Context, preAuthCode string) {
	c.config.preAuthCode = preAuthCode
}

// SetAuthorizerAppid 设置授权方appid
func (c *Client) SetAuthorizerAppid(ctx context.Context, authorizerAppid string) {
	c.config.authorizerAppid = authorizerAppid
}

// SetAuthorizerAccessToken 设置授权方access_token
func (c *Client) SetAuthorizerAccessToken(ctx context.Context, authorizerAccessToken string) {
	c.config.authorizerAccessToken = authorizerAccessToken
}

// SetAuthorizerRefreshToken 设置授权方refresh_token
func (c *Client) SetAuthorizerRefreshToken(ctx context.Context, authorizerRefreshToken string) {
	c.config.authorizerRefreshToken = authorizerRefreshToken
}
