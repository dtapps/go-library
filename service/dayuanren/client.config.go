package dayuanren

func (c *Client) GetApiURL() string {
	return c.config.apiURL
}

func (c *Client) SetApiURL(v string) *Client {
	c.config.apiURL = v
	return c
}

func (c *Client) GetUserID() int64 {
	return c.config.userID
}

func (c *Client) SetUserID(v int64) *Client {
	c.config.userID = v
	return c
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}

func (c *Client) SetApiKey(v string) *Client {
	c.config.apiKey = v
	return c
}
