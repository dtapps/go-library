package wechatminiprogram

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/golog"
)

// ConfigApp 配置
func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.appId = appId
	c.config.appSecret = appSecret
	return c
}

// ConfigRedisClient 缓存数据库
func (c *Client) ConfigRedisClient(client *dorm.RedisClient) {
	c.cache.redisClient = client
}

// ConfigRedisCachePrefixFunWechatAccessToken 缓存前缀
func (c *Client) ConfigRedisCachePrefixFunWechatAccessToken(config string) error {
	c.cache.wechatAccessTokenPrefix = config
	if c.cache.wechatAccessTokenPrefix == "" {
		return redisCachePrefixNoConfig
	}
	return nil
}

// ConfigZapClientFun 日志配置
func (c *Client) ConfigZapClientFun(apiZapLogFun golog.ApiZapLogFun) {
	apiZapLog := apiZapLogFun()
	if apiZapLog != nil {
		c.zap.client = apiZapLog
		c.zap.status = true
	}
}
