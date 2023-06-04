package pintoto

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}
