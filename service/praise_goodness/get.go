package praise_goodness

func (c *Client) GetMchID() int64 {
	return c.config.mchID
}

func (c *Client) GetKey() string {
	return c.config.Key
}
