package golog

// ConfigSLogClientFun 日志配置
func (c *GinCustomClient) ConfigSLogClientFun(sLogFun SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
}

// ConfigSLogResultClientFun 日志配置然后返回
func (c *GinCustomClient) ConfigSLogResultClientFun(sLogFun SLogFun) *GinCustomClient {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
	return c
}
