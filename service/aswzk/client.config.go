package aswzk

func (c *Client) GetUrl() string {
	return c.config.baseURL
}

func (c *Client) SetUrl(v string) *Client {
	c.config.baseURL = v
	return c
}

func (c *Client) GetUserID() string {
	return c.config.userID
}

func (c *Client) SetUserID(v string) *Client {
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
