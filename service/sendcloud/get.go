package sendcloud

func (c *Client) GetApiUser() string {
	return c.config.ApiUser
}

func (c *Client) GetApiKey() string {
	return c.config.ApiKey
}
