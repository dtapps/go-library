package wechatminiprogram

import (
	"context"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) getAccessToken(ctx context.Context) string {
	if c.config.selfAccessToken {
		return c.config.accessToken
	}
	c.config.accessToken = c.GetAccessToken(ctx)
	return c.config.accessToken
}
