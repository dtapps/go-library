package framework

// Param 获取路径参数
func (c *Context) Param(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Param(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Param(key)
	}
	return ""
}
