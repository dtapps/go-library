package taobao

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetAdzoneId() int64 {
	return c.config.adzoneId
}
