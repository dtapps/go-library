package amap

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
