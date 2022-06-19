package jd

// GetAppKey 应用Key
func (c *Client) GetAppKey() string {
	return c.config.AppKey
}

// GetSecretKey 密钥
func (c *Client) GetSecretKey() string {
	return c.config.SecretKey
}

// GetSiteId 网站ID/APP ID
func (c *Client) GetSiteId() string {
	return c.config.SiteId
}

// GetPositionId 推广位id
func (c *Client) GetPositionId() string {
	return c.config.PositionId
}
