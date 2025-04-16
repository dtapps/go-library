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

// GetHeader 设置响应头部信息
func (c *Context) GetHeader(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.GetHeader(key)
	}
	if c.hertzCtx != nil {
		return string(c.hertzCtx.GetHeader(key))
	}
	return ""
}

// Host 获取请求域名
func (c *Context) Host() string {
	if c.ginCtx != nil {
		return c.ginCtx.Request.Host
	}
	if c.hertzCtx != nil {
		return string(c.hertzCtx.Host())
	}
	return ""
}

// Method 获取请求方法
func (c *Context) Method() string {
	if c.ginCtx != nil {
		return c.ginCtx.Request.Method
	}
	if c.hertzCtx != nil {
		return string(c.hertzCtx.Method())
	}
	return ""
}

// UserAgent 获取请求客户端的UserAgent
func (c *Context) UserAgent() string {
	if c.ginCtx != nil {
		return c.ginCtx.Request.UserAgent()
	}
	if c.hertzCtx != nil {
		return string(c.hertzCtx.UserAgent())
	}
	return ""
}

// ContentType 获取请求客户端的ContentType
func (c *Context) ContentType() string {
	if c.ginCtx != nil {
		return c.ginCtx.ContentType()
	}
	if c.hertzCtx != nil {
		return string(c.hertzCtx.ContentType())
	}
	return ""
}
