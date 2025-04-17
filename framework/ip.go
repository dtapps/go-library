package framework

// ClientIP 获取请求客户端的IP
func (c *Context) ClientIP() string {
	if c.IsGin() {
		return c.ginCtx.ClientIP()
	}
	if c.IsHertz() {
		return c.hertzCtx.ClientIP()
	}
	return ""
}
