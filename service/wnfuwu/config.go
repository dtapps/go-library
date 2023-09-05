package wnfuwu

import "github.com/dtapps/go-library/utils/golog"

// ConfigApp 配置
func (c *Client) ConfigApp(userId int64, apiKey string) *Client {
	c.config.userId = userId
	c.config.apiKey = apiKey
	return c
}

// ConfigSLogClientFun 日志配置
func (c *Client) ConfigSLogClientFun(apiSLogFun golog.ApiSLogFun) {
	apiSLog := apiSLogFun()
	if apiSLog != nil {
		c.slog.client = apiSLog
		c.slog.status = true
	}
}
