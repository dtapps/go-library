package wechatminiprogram

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
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

// ConfigSLogClientFun 日志配置
func (c *Client) ConfigSLogClientFun(sLogFun golog.SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
}

// SetHttp 配置请求
func (c *Client) SetHttp(app *gorequest.App) {
	c.requestClient = app
	c.requestClientStatus = true
	c.requestClient.Uri = apiUrl
}

// DefaultHttp 默认请求
func (c *Client) DefaultHttp() {
	c.requestClient = gorequest.NewHttp()
	c.requestClientStatus = true
	c.requestClient.Uri = apiUrl
}
