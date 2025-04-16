package framework

// ClientIP 获取请求客户端的IP
func (c *Context) ClientIP() string {
	if c.ginCtx != nil {
		return c.ginCtx.ClientIP()
	} else if c.hertzCtx != nil {
		return c.hertzCtx.ClientIP()
	}
	return ""
}
