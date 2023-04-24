package wikeyun

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetStoreId() int64 {
	return c.config.storeId
}

func (c *Client) GetAppKey() int64 {
	return c.config.appKey
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetClientIp() string {
	return c.config.clientIp
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
