package jd

import "github.com/dtapps/go-library/utils/golog"

// GetAppKey 应用Key
func (c *Client) GetAppKey() string {
	return c.config.appKey
}

// GetSecretKey 密钥
func (c *Client) GetSecretKey() string {
	return c.config.secretKey
}

// GetSiteId 网站ID/APP ID
func (c *Client) GetSiteId() string {
	return c.config.siteId
}

// GetPositionId 推广位id
func (c *Client) GetPositionId() string {
	return c.config.positionId
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
