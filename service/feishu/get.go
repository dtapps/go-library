package feishu

func (c *Client) GetKey() string {
	return c.config.key
}
