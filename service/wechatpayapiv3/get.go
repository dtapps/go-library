package wechatpayapiv3

func (c *Client) GetAppId() string {
	return c.config.AppId
}

func (c *Client) GetMchId() string {
	return c.config.MchId
}

func (c *Client) GetMchSslKey() string {
	return c.config.MchSslKey
}

func (c *Client) GetMchSslSerialNo() string {
	return c.config.MchSslSerialNo
}
