package framework

// Header 设置响应头部信息
func (c *Context) Header(key, value string) {
	if c.IsGin() {
		c.ginCtx.Header(key, value)
	}
	if c.IsHertz() {
		c.hertzCtx.Header(key, value)
	}
	if c.IsEcho() {
		c.hertzCtx.Header(key, value)
	}
}

// GetHeader 设置响应头部信息
func (c *Context) GetHeader(key string) string {
	if c.IsGin() {
		return c.ginCtx.GetHeader(key)
	}
	if c.IsHertz() {
		return string(c.hertzCtx.GetHeader(key))
	}
	if c.IsEcho() {
	}
	return ""
}

// Host 获取请求域名
func (c *Context) Host() string {
	if c.IsGin() {
		return c.ginCtx.Request.Host
	}
	if c.IsHertz() {
		return string(c.hertzCtx.Host())
	}
	if c.IsEcho() {
	}
	return ""
}

// Method 获取请求方法
func (c *Context) Method() string {
	if c.IsGin() {
		return c.ginCtx.Request.Method
	}
	if c.IsHertz() {
		return string(c.hertzCtx.Method())
	}
	if c.IsEcho() {
	}
	return ""
}

// UserAgent 获取请求客户端的UserAgent
func (c *Context) UserAgent() string {
	if c.IsGin() {
		return c.ginCtx.Request.UserAgent()
	}
	if c.IsHertz() {
		return string(c.hertzCtx.UserAgent())
	}
	if c.IsEcho() {
	}
	return ""
}

// ContentType 获取请求客户端的ContentType
func (c *Context) ContentType() string {
	if c.IsGin() {
		return c.ginCtx.ContentType()
	}
	if c.IsHertz() {
		return string(c.hertzCtx.ContentType())
	}
	if c.IsEcho() {
	}
	return ""
}

// StatusCode 获取请求客户端的StatusCode
func (c *Context) StatusCode() int {
	if c.IsGin() {
		return c.ginCtx.Writer.Status()
	}
	if c.IsHertz() {
		return c.hertzCtx.Response.StatusCode()
	}
	if c.IsEcho() {
	}
	return 0
}

// FullPath 返回当前请求匹配的完整路由模板路径，例如 /user/:id。
// 常用于权限校验、路由分组识别等场景。
// 注意：此方法返回的是框架注册时的模板路径，而不是用户实际访问的 URL。
func (c *Context) FullPath() string {
	if c.IsGin() {
		return c.ginCtx.FullPath()
	}
	if c.IsHertz() {
		return c.hertzCtx.FullPath()
	}
	if c.IsEcho() {
	}
	return ""
}

// Path 返回当前请求的实际访问路径，例如 /user/123。
// 常用于业务逻辑处理、日志记录等。
// 注意：该路径已被框架解码并标准化，不包含查询参数。
func (c *Context) Path() string {
	if c.IsGin() {
		return c.ginCtx.Request.URL.Path
	}
	if c.IsHertz() {
		return string(c.hertzCtx.Path())
	}
	if c.IsEcho() {
	}
	return ""
}
