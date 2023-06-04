package sendcloud

func (c *Client) GetApiUser() string {
	return c.config.apiUser
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}
