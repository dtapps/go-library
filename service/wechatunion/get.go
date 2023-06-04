package wechatunion

import (
	"context"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetPid() string {
	return c.config.pid
}

func (c *Client) getAccessToken(ctx context.Context) string {
	c.config.accessToken = c.GetAccessToken(ctx)
	return c.config.accessToken
}
