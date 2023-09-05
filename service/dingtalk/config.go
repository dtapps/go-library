package dingtalk

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) Config(secret, accessToken string) *Client {
	c.config.secret = secret
	c.config.accessToken = accessToken
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
