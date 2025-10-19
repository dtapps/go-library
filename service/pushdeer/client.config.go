package pushdeer

func (c *Client) SetUrl(baseURL string) *Client {
	if baseURL != "" {
		c.config.baseURL = baseURL
	}
	return c
}

func (c *Client) SetPushKey(pushKey string) *Client {
	c.config.pushKey = pushKey
	return c
}
