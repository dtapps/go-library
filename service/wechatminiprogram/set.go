package wechatminiprogram

import (
	"context"
)

func (c *Client) SetAccessToken(ctx context.Context, accessToken string) string {
	c.config.accessToken = accessToken
	c.config.selfAccessToken = true
	return c.config.accessToken
}
