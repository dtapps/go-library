package wechatopen

import (
	"context"
	"fmt"
	"time"
)

// 微信后台推送的ticke
func (c *Client) getComponentVerifyTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_open:component_verify_ticket:%v", c.config.ComponentAppId)
}

// SetComponentVerifyTicket 设置微信后台推送的ticke
func (c *Client) SetComponentVerifyTicket(componentVerifyTicket string) string {
	if componentVerifyTicket == "" {
		return ""
	}
	c.config.RedisClient.Db.Set(context.Background(), c.getComponentVerifyTicketCacheKeyName(), componentVerifyTicket, time.Hour*12)
	return c.GetComponentVerifyTicket()
}

// GetComponentVerifyTicket 获取微信后台推送的ticke
func (c *Client) GetComponentVerifyTicket() string {
	if c.config.RedisClient.Db == nil {
		return c.config.ComponentVerifyTicket
	}
	result, _ := c.config.RedisClient.Db.Get(context.Background(), c.getComponentVerifyTicketCacheKeyName()).Result()
	return result
}

// 令牌
func (c *Client) getComponentAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_open:component_access_token:%v", c.config.ComponentAppId)
}

// SetComponentAccessToken 设置令牌
func (c *Client) SetComponentAccessToken(componentAccessToken string) string {
	if componentAccessToken == "" {
		return ""
	}
	c.config.RedisClient.Db.Set(context.Background(), c.getComponentAccessTokenCacheKeyName(), componentAccessToken, time.Second*7200)
	return c.GetComponentAccessToken()
}

// GetComponentAccessToken 获取令牌
func (c *Client) GetComponentAccessToken() string {
	if c.config.RedisClient.Db == nil {
		return c.config.ComponentAccessToken
	}
	result, _ := c.config.RedisClient.Db.Get(context.Background(), c.getComponentAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorComponentAccessToken 监控令牌
func (c *Client) MonitorComponentAccessToken() string {
	// 查询
	componentAccessToken := c.GetComponentAccessToken()
	// 判断
	result := c.CgiBinGetApiDomainIp(componentAccessToken)
	if len(result.Result.IpList) > 0 {
		return componentAccessToken
	}
	// 重新获取
	return c.SetComponentAccessToken(c.CgiBinComponentApiComponentToken().Result.ComponentAccessToken)
}

// 授权方令牌
func (c *Client) getAuthorizerAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_open:authorizer_access_token:%v:%v", c.config.ComponentAppId, c.config.AuthorizerAppid)
}

// SetAuthorizerAccessToken 设置授权方令牌
func (c *Client) SetAuthorizerAccessToken(authorizerAccessToken string) string {
	if authorizerAccessToken == "" {
		return ""
	}
	c.config.RedisClient.Db.Set(context.Background(), c.getAuthorizerAccessTokenCacheKeyName(), authorizerAccessToken, time.Hour*2)
	return c.GetComponentAccessToken()
}

// GetAuthorizerAccessToken 获取授权方令牌
func (c *Client) GetAuthorizerAccessToken() string {
	if c.config.RedisClient.Db == nil {
		return c.config.AuthorizerAccessToken
	}
	result, _ := c.config.RedisClient.Db.Get(context.Background(), c.getAuthorizerAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorAuthorizerAccessToken 监控授权方令牌
func (c *Client) MonitorAuthorizerAccessToken(authorizerRefreshToken string) string {
	// 查询
	authorizerAccessToken := c.GetAuthorizerAccessToken()
	// 判断
	if authorizerAccessToken != "" {
		return authorizerAccessToken
	}
	// 重新获取
	return c.SetAuthorizerAccessToken(c.CgiBinComponentApiAuthorizerToken(authorizerRefreshToken).Result.AuthorizerAccessToken)
}

// 预授权码
func (c *Client) getPreAuthCodeCacheKeyName() string {
	return fmt.Sprintf("wechat_open:pre_auth_code:%v", c.config.ComponentAppId)
}

// SetPreAuthCode 设置预授权码
func (c *Client) SetPreAuthCode(preAuthCode string) string {
	if preAuthCode == "" {
		return ""
	}
	c.config.RedisClient.Db.Set(context.Background(), c.getPreAuthCodeCacheKeyName(), preAuthCode, time.Second*1700)
	return c.GetComponentAccessToken()
}

// GetPreAuthCode 获取预授权码
func (c *Client) GetPreAuthCode() string {
	if c.config.RedisClient.Db == nil {
		return c.config.AuthorizerAccessToken
	}
	result, _ := c.config.RedisClient.Db.Get(context.Background(), c.getPreAuthCodeCacheKeyName()).Result()
	return result
}

// DelPreAuthCode 删除预授权码
func (c *Client) DelPreAuthCode() error {
	return c.config.RedisClient.Db.Del(context.Background(), c.getPreAuthCodeCacheKeyName()).Err()
}

// MonitorPreAuthCode 监控预授权码
func (c *Client) MonitorPreAuthCode() string {
	// 查询
	preAuthCode := c.GetPreAuthCode()
	// 判断
	if preAuthCode != "" {
		return preAuthCode
	}
	// 重新获取
	return c.SetPreAuthCode(c.CgiBinComponentApiCreatePreAuthCoden().Result.PreAuthCode)
}
