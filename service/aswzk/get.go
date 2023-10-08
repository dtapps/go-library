package aswzk

func (c *Client) GetUserID() string {
	return c.config.userID
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}
