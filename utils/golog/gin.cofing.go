package golog

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/goip"
	"runtime"
)

// ConfigSLogClientFun 日志配置
func (c *GinClient) ConfigSLogClientFun(sLogFun SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
}

func (c *GinClient) setConfig(ctx context.Context) {

	info := getSystem()

	c.config.systemHostname = info.SystemHostname
	c.config.systemOs = info.SystemOs
	c.config.systemVersion = info.SystemVersion
	c.config.systemKernel = info.SystemKernel
	c.config.systemKernelVersion = info.SystemKernelVersion
	c.config.systemBootTime = info.SystemBootTime
	c.config.cpuCores = info.CpuCores
	c.config.cpuModelName = info.CpuModelName
	c.config.cpuMhz = info.CpuMhz

	c.config.systemInsideIp = goip.GetInsideIp(ctx)

	c.config.sdkVersion = go_library.Version()
	c.config.goVersion = runtime.Version()
}
