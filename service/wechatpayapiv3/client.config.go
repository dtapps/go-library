package wechatpayapiv3

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetMchId() string {
	return c.config.mchId
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
