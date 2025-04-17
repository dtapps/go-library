package framework

// GetRawQuery 查询字符串
func (c *Context) GetRawQuery() string {
	if c.ginCtx != nil {
		return c.ginCtx.Request.URL.RawQuery
	}
	if c.hertzCtx != nil {
		return string(c.hertzCtx.Request.QueryString())
	}
	return ""
}

func (c *Context) Query(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Query(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Query(key)
	}
	return ""
}

func (c *Context) DefaultQuery(key, defaultValue string) string {
	if c.ginCtx != nil {
		return c.ginCtx.DefaultQuery(key, defaultValue)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.DefaultQuery(key, defaultValue)
	}
	return ""
}

func (c *Context) GetQuery(key string) (string, bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetQuery(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetQuery(key)
	}
	return "", false
}
