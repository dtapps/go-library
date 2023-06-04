package meituan

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetSecret() string {
	return c.config.secret
}
