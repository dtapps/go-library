package wechatopen

// SetAuthorizerAppid 设置代理商小程序
func (c *Client) SetAuthorizerAppid(authorizerAppid string) {
	c.config.AuthorizerAppid = authorizerAppid
	return
}
