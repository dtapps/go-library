package wechatminiprogram

import "context"

func (c *Client) getAppId() string {
	return c.config.AppId
}

func (c *Client) getAppSecret() string {
	return c.config.AppSecret
}

func (c *Client) getAccessToken(ctx context.Context) string {
	c.config.AccessToken = c.GetAccessToken(ctx)
	return c.config.AccessToken
}

func (c *Client) GetAppId() string {
	return c.config.AppId
}
