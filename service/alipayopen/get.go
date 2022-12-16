package alipayopen

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetAppRSA2() string {
	return c.config.appRSA2
}

func (c *Client) GetAlipayRSA2() string {
	return c.config.alipayRSA2
}

func (c *Client) GetAes() string {
	return c.config.aes
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
