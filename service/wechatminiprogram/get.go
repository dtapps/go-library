package wechatminiprogram

import (
	"context"
	"github.com/dtapps/go-library/utils/golog"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) getAccessToken(ctx context.Context) string {
	c.config.accessToken = c.GetAccessToken(ctx)
	return c.config.accessToken
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
