package golog

import (
	"context"
	go_library "github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/goip"
	"runtime"
)

// ConfigSLogClientFun 日志配置
func (sl *ApiSLog) ConfigSLogClientFun(sLogFun SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		sl.slog.client = sLog
		sl.slog.status = true
	}
}

func (sl *ApiSLog) setConfig(ctx context.Context) {

	info := getSystem()

	sl.systemConfig.systemHostname = info.SystemHostname
	sl.systemConfig.systemOs = info.SystemOs
	sl.systemConfig.systemKernel = info.SystemKernel

	sl.systemConfig.systemInsideIp = goip.GetInsideIp(ctx)

	sl.systemConfig.sdkVersion = go_library.Version()
	sl.systemConfig.goVersion = runtime.Version()

}
