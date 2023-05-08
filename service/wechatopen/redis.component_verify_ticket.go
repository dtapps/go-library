package wechatopen

import (
	"context"
	"time"
)

// GetComponentVerifyTicketCacheKeyName 第三方平台推送ticket 缓存名称
func GetComponentVerifyTicketCacheKeyName(ctx context.Context, c *Client) string {
	return c.cache.componentVerifyTicketPrefix + c.GetComponentAppId(ctx)
}

// GetComponentVerifyTicket 第三方平台推送ticket 获取
func GetComponentVerifyTicket(ctx context.Context, c *Client) string {
	if c.cache.redisClient.GetDb() == nil {
		return c.config.componentVerifyTicket
	}
	result, _ := c.cache.redisClient.Get(ctx, GetComponentVerifyTicketCacheKeyName(ctx, c)).Result()
	return result
}

// SetComponentVerifyTicket 第三方平台推送ticket 设置
func SetComponentVerifyTicket(ctx context.Context, c *Client, componentVerifyTicket string) string {
	if componentVerifyTicket == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, GetComponentVerifyTicketCacheKeyName(ctx, c), componentVerifyTicket, time.Hour*12)
	return GetComponentVerifyTicket(ctx, c)
}

// DelComponentVerifyTicket 第三方平台推送ticket 删除
func DelComponentVerifyTicket(ctx context.Context, c *Client) error {
	return c.cache.redisClient.Del(ctx, GetComponentVerifyTicketCacheKeyName(ctx, c)).Err()
}
