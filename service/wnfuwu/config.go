package wnfuwu

import "github.com/dtapps/go-library/utils/golog"

// ConfigApp 配置
func (c *Client) ConfigApp(userId int64, apiKey string) *Client {
	c.config.userId = userId
	c.config.apiKey = apiKey
	return c
}

// ConfigZapClientFun 日志配置
func (c *Client) ConfigZapClientFun(apiZapLogFun golog.ApiZapLogFun) {
	apiZapLog := apiZapLogFun()
	if apiZapLog != nil {
		c.zap.client = apiZapLog
		c.zap.status = true
	}
}
