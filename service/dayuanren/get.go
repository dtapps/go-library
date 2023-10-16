package dayuanren

func (c *Client) GetApiURL() string {
	return c.config.apiURL
}
func (c *Client) GetUserID() int64 {
	return c.config.userID
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}
