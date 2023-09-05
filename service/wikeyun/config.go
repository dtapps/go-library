package wikeyun

import "github.com/dtapps/go-library/utils/golog"

// ConfigApp 配置
func (c *Client) ConfigApp(storeId, appKey int64, appSecret string) *Client {
	c.config.storeId = storeId
	c.config.appKey = appKey
	c.config.appSecret = appSecret
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
