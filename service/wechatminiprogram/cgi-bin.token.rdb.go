package wechatminiprogram

import (
	"context"
	"fmt"
	"time"
)

func (c *Client) GetAccessToken(ctx context.Context) string {
	if c.config.RedisClient.Db == nil {
		return c.config.AccessToken
	}
	newCache := c.config.RedisClient.NewSimpleStringCache(c.config.RedisClient.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := c.CgiBinToken(ctx)
		return token.Result.AccessToken
	}
	return newCache.GetCache(c.getAccessTokenCacheKeyName())
}

func (c *Client) getAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_access_token:%v", c.config.AppId)
}
