package wechatpayopen

// SubConfig 子商户配置
func (c *Client) SubConfig(subAppid, subMchId string) *Client {
	c.config.SpAppid = subAppid
	c.config.SubMchId = subMchId
	return c
}
