package wechatoffice

import (
	"context"
	"time"
)

// GetJsapiTicket 获取api_ticket
func (c *Client) GetJsapiTicket(ctx context.Context) string {
	if c.cache.redisClient.GetDb() == nil {
		return c.config.jsapiTicket
	}
	newCache := c.cache.redisClient.NewSimpleStringCache(c.cache.redisClient.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := c.CgiBinTicketGetTicket(ctx, "jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(ctx, c.getJsapiTicketCacheKeyName())
}

func (c *Client) getJsapiTicketCacheKeyName() string {
	return c.cache.wechatJsapiTicketPrefix + c.GetAppId()
}
