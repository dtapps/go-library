package golog

// ConfigSLogClientFun 日志配置
func (sl *ApiSLog) ConfigSLogClientFun(sLogFun SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		sl.slog.client = sLog
		sl.slog.status = true
	}
}

// ConfigSLogResultClientFun 日志配置然后返回
func (sl *ApiSLog) ConfigSLogResultClientFun(sLogFun SLogFun) *ApiSLog {
	sLog := sLogFun()
	if sLog != nil {
		sl.slog.client = sLog
		sl.slog.status = true
	}
	return sl
}
