package pinduoduo

func (c *Client) GetClientId() string {
	return c.config.clientId
}

func (c *Client) GetClientSecret() string {
	return c.config.clientSecret
}

func (c *Client) GetMediaId() string {
	return c.config.mediaId
}

func (c *Client) GetPid() string {
	return c.config.pid
}
