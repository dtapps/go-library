package chengquan

func (c *Client) GetURL() string {
	return c.config.baseURL
}

func (c *Client) SetURL(v string) *Client {
	c.config.baseURL = v
	return c
}

func (c *Client) GetAppID() string {
	return c.config.appID
}

func (c *Client) SetAppID(v string) *Client {
	c.config.appID = v
	return c
}

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) SetAppKey(v string) *Client {
	c.config.appKey = v
	return c
}

func (c *Client) GetAesKey() string {
	return c.config.aesKey
}

func (c *Client) SetAesKey(v string) *Client {
	c.config.aesKey = v
	return c
}

func (c *Client) GetAesIv() string {
	return c.config.aesIv
}

func (c *Client) SetAesIv(v string) *Client {
	c.config.aesIv = v
	return c
}
