package chengquan

func (c *Client) GetAppID() string {
	return c.config.appID
}

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) GetAesKey() string {
	return c.config.aesKey
}

func (c *Client) GetAesIv() string {
	return c.config.aesIv
}
