package gorequest

import (
	"github.com/dtapps/go-library"
	"github.com/shirou/gopsutil/host"
	"log"
	"runtime"
)

type systemResult struct {
	SystemOs     string // 系统类型
	SystemKernel string // 系统内核
}

// 获取系统信息
func getSystem() (result systemResult) {

	hInfo, err := host.Info()
	if err != nil {
		log.Printf("getSystem.host.Info：%s\n", err)
	}

	result.SystemOs = hInfo.OS
	result.SystemKernel = hInfo.KernelArch

	return result
}

// 设置配置信息
func (app *App) setConfig() {

	info := getSystem()

	app.config.systemOs = info.SystemOs
	app.config.systemKernel = info.SystemKernel

	app.config.sdkVersion = go_library.Version()
	app.config.goVersion = runtime.Version()

}
