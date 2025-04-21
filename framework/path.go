package framework

// Param 获取路径参数
func (c *Context) Param(key string) string {
	if c.IsGin() {
		return c.ginCtx.Param(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.Param(key)
	}
	if c.IsEcho() {
		return c.echoCtx.Param(key)
	}
	return ""
}
