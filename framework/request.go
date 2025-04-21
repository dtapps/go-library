package framework

type RequestWrapper struct {
	c *Context
}

// Request 返回请求相关的封装方法
func (c *Context) Request() *RequestWrapper {
	return &RequestWrapper{c: c}
}
