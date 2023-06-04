package pinduoduo

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) ConfigPid(pid string) *Client {
	c.config.pid = pid
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
