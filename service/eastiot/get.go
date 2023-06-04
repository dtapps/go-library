package eastiot

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}
