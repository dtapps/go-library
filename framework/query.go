package framework

// GetRawQuery 查询字符串
func (c *Context) GetRawQuery() string {
	if c.IsGin() {
		return c.ginCtx.Request.URL.RawQuery
	}
	if c.IsHertz() {
		return string(c.hertzCtx.Request.QueryString())
	}
	if c.IsEcho() {
		return c.echoCtx.QueryParams().Encode()
	}
	return ""
}

func (c *Context) Query(key string) string {
	if c.IsGin() {
		return c.ginCtx.Query(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.Query(key)
	}
	if c.IsEcho() {
		return c.echoCtx.QueryParam(key)
	}
	return ""
}

func (c *Context) DefaultQuery(key, defaultValue string) string {
	if c.IsGin() {
		return c.ginCtx.DefaultQuery(key, defaultValue)
	}
	if c.IsHertz() {
		return c.hertzCtx.DefaultQuery(key, defaultValue)
	}
	if c.IsEcho() {
		value := c.echoCtx.QueryParam(key)
		if value == "" {
			return defaultValue
		}
		return value
	}
	return ""
}

func (c *Context) GetQuery(key string) (string, bool) {
	if c.IsGin() {
		return c.ginCtx.GetQuery(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.GetQuery(key)
	}
	if c.IsEcho() {
		value := c.echoCtx.QueryParam(key)
		return value, value != ""
	}
	return "", false
}
