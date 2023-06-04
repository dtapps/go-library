package wnfuwu

func (c *Client) GetUserId() int64 {
	return c.config.userId
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}
