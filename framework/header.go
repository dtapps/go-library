package framework

// Header 设置响应头部信息
func (c *Context) Header(key, value string) {
	if c.ginCtx != nil {
		c.ginCtx.Header(key, value)
	}
	if c.hertzCtx != nil {
		c.hertzCtx.Header(key, value)
	}
}
