package kuaidi100

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) Config(customer string) *Client {
	c.config.customer = customer
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