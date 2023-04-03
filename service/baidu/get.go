package baidu

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetAk() string {
	return c.config.ak
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
