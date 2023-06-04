package wechatpayopen

func (c *Client) GetSpAppid() string {
	return c.config.spAppid
}

func (c *Client) GetSpMchId() string {
	return c.config.spMchId
}

func (c *Client) GetSubAppid() string {
	return c.config.subAppid
}

func (c *Client) GetSubMchId() string {
	return c.config.subMchId
}

func (c *Client) GetMchSslKey() string {
	return c.config.mchSslKey
}

func (c *Client) GetMchSslSerialNo() string {
	return c.config.mchSslSerialNo
}

func (c *Client) GetApiV3() string {
	return c.config.apiV3
}
