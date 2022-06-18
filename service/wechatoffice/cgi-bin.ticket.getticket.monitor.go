package wechatoffice

import (
	"context"
	"errors"
	"time"
)

// GetJsapiTicketMonitor 监控api_ticket
func (c *Client) GetJsapiTicketMonitor() (string, error) {
	if c.config.RedisClient.Db == nil {
		return "", errors.New("驱动没有初始化")
	}
	result := c.DebugCgiBinTicketCheck()
	if result.Result.Errcode == 0 {
		return c.config.JsapiTicket, nil
	}
	c.config.AccessToken = c.GetAccessToken()
	token := c.CgiBinTicketGetTicket("jsapi")
	c.config.RedisClient.Db.Set(context.Background(), c.getJsapiTicketCacheKeyName(), token.Result.Ticket, time.Second*7000)
	return token.Result.Ticket, nil
}
