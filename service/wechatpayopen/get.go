package wechatpayopen

func (c *Client) GetSpAppid() string {
	return c.config.SpAppid
}

func (c *Client) GetSpMchId() string {
	return c.config.SpMchId
}

func (c *Client) GetSubAppid() string {
	return c.config.SubAppid
}

func (c *Client) GetSubMchId() string {
	return c.config.SubMchId
}
