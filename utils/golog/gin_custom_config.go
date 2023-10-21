package golog

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/goip"
	"runtime"
)

// ConfigSLogClientFun 日志配置
func (c *GinCustomClient) ConfigSLogClientFun(sLogFun SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
}

func (c *GinCustomClient) setConfig(ctx context.Context) {

	info := getSystem()

	c.config.systemHostname = info.SystemHostname
	c.config.systemOs = info.SystemOs
	c.config.systemKernel = info.SystemKernel

	c.config.systemInsideIp = goip.GetInsideIp(ctx)

	c.config.sdkVersion = go_library.Version()
	c.config.goVersion = runtime.Version()
}
