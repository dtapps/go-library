package wikeyun

import "github.com/dtapps/go-library/utils/golog"

// ConfigApp 配置
func (c *Client) ConfigApp(storeId, appKey int64, appSecret string) *Client {
	c.config.storeId = storeId
	c.config.appKey = appKey
	c.config.appSecret = appSecret
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
