package wechatpayopen

func (c *Client) GetSpAppid() string {
	return c.config.spAppid
}
func (c *Client) SetSpAppid(v string) *Client {
	c.config.spAppid = v
	return c
}

func (c *Client) GetSpMchId() string {
	return c.config.spMchId
}
func (c *Client) SetSpMchId(v string) *Client {
	c.config.spMchId = v
	return c
}

func (c *Client) GetSubAppid() string {
	return c.config.subAppid
}

func (c *Client) SetSubAppid(v string) *Client {
	c.config.subAppid = v
	return c
}

func (c *Client) GetSubMchId() string {
	return c.config.subMchId
}
func (c *Client) SetSubMchId(v string) *Client {
	c.config.subMchId = v
	return c
}

func (c *Client) GetApiV3() string {
	return c.config.apiV3
}
func (c *Client) SetApiV3(v string) *Client {
	c.config.apiV3 = v
	return c
}

func (c *Client) GetCertificateSerialNo() string {
	return c.config.certificateSerialNo
}
func (c *Client) SetCertificateSerialNo(v string) *Client {
	c.config.certificateSerialNo = v
	return c
}
