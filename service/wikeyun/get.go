package wikeyun

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
