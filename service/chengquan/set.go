package chengquan

func (c *Client) SetAppID(appID string) {
	c.config.appID = appID
}

func (c *Client) SetAppKey(appKey string) {
	c.config.appKey = appKey
}

func (c *Client) SetAesKey(aesKey string) {
	c.config.aesKey = aesKey
}

func (c *Client) SetAesIv(aesIv string) {
	c.config.aesIv = aesIv
}
