package wechatoffice

import (
	"context"
	"errors"
	"time"
)

// GetJsapiTicketMonitor 监控api_ticket
func (c *Client) GetJsapiTicketMonitor(ctx context.Context) (string, error) {
	if c.cache.redisClient.GetDb() == nil {
		return "", errors.New("驱动没有初始化")
	}
	result := c.DebugCgiBinTicketCheck(ctx)
	if result.Result.Errcode == 0 {
		return c.config.jsapiTicket, nil
	}
	c.config.accessToken = c.GetAccessToken(ctx)
	token := c.CgiBinTicketGetTicket(ctx, "jsapi")
	c.cache.redisClient.Set(ctx, c.getJsapiTicketCacheKeyName(), token.Result.Ticket, time.Second*7000)
	return token.Result.Ticket, nil
}
