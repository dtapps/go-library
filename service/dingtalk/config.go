package dingtalk

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) Config(secret, accessToken string) *Client {
	c.config.secret = secret
	c.config.accessToken = accessToken
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
