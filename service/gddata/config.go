package gddata

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) Config(token string) *Client {
	c.config.token = token
	return c
}

// ConfigApiClientFun 日志配置
func (c *Client) ConfigApiClientFun(apiClientFun golog.ApiClientFun) {
	apiClient := apiClientFun()
	if apiClient != nil {
		c.log.client = apiClient
		c.log.status = true
	}
}
