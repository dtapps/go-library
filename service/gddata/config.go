package gddata

func (c *Client) Config(token string) *Client {
	c.config.token = token
	return c
}
