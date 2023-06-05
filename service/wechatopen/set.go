package wechatopen

import "context"

// SetAuthorizerAccessToken 授权方access_token
func (c *Client) SetAuthorizerAccessToken(ctx context.Context, authorizerAccessToken string) string {
	if authorizerAccessToken == "" {
		return ""
	}
	c.config.authorizerAccessToken = authorizerAccessToken
	return c.config.authorizerAccessToken
}

// SetAuthorizerRefreshToken 授权方refresh_token
func (c *Client) SetAuthorizerRefreshToken(ctx context.Context, authorizerRefreshToken string) string {
	if authorizerRefreshToken == "" {
		return ""
	}
	c.config.authorizerRefreshToken = authorizerRefreshToken
	return c.config.authorizerRefreshToken
}
