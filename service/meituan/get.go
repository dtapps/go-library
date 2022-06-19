package meituan

func (c *Client) GetAppKey() string {
	return c.config.AppKey
}
