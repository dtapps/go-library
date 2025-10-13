package dayuanren

func (c *Client) GetURL() string {
	return c.config.baseURL
}

func (c *Client) SetURL(v string) *Client {
	c.config.baseURL = v
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
