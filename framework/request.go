package framework

type RequestWrapper struct {
	c *Context
}

// Request 返回请求相关的封装方法
func (c *Context) Request() *RequestWrapper {
	return &RequestWrapper{c: c}
}

// GetHeader 获取请求的Header
func (cr *RequestWrapper) GetHeader(key string) string {
	if cr.c.IsGin() {
		return cr.c.ginCtx.GetHeader(key)
	}
	if cr.c.IsHertz() {
		return string(cr.c.hertzCtx.Request.Header.Peek(key))
	}
	if cr.c.IsEcho() {
		return cr.c.echoCtx.Request().Header.Get(key)
	}
	return ""
}

// Host 获取请求域名
func (cr *RequestWrapper) Host() string {
	if cr.c.IsGin() {
		return cr.c.ginCtx.Request.Host
	}
	if cr.c.IsHertz() {
		return string(cr.c.hertzCtx.URI().Host())
	}
	if cr.c.IsEcho() {
		return cr.c.echoCtx.Request().Host
	}
	return ""
}

// Method 获取请求方法
func (cr *RequestWrapper) Method() string {
	if cr.c.IsGin() {
		return cr.c.ginCtx.Request.Method
	}
	if cr.c.IsHertz() {
		return string(cr.c.hertzCtx.Request.Header.Method())
	}
	if cr.c.IsEcho() {
		return cr.c.echoCtx.Request().Method
	}
	return ""
}

// UserAgent 获取请求客户端的UserAgent
func (cr *RequestWrapper) UserAgent() string {
	if cr.c.IsGin() {
		return cr.c.ginCtx.Request.UserAgent()
	}
	if cr.c.IsHertz() {
		return string(cr.c.hertzCtx.Request.Header.UserAgent())
	}
	if cr.c.IsEcho() {
		return cr.c.echoCtx.Request().UserAgent()
	}
	return ""
}

// ContentType 获取请求客户端的ContentType
func (cr *RequestWrapper) ContentType() string {
	if cr.c.IsGin() {
		return cr.c.ginCtx.ContentType()
	}
	if cr.c.IsHertz() {
		return string(cr.c.hertzCtx.Request.Header.ContentType())
	}
	if cr.c.IsEcho() {
		return cr.c.echoCtx.Request().Header.Get("Content-Type")
	}
	return ""
}
