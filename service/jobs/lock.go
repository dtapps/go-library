package jobs

import (
	"dtapps/dta/library/utils/goredis"
	"fmt"
	"time"
)

// Lock 上锁
func (app *App) Lock(info Task, id any) string {
	cacheName := fmt.Sprintf("cron:%v:%v", info.Type, id)
	judgeCache := app.Redis.NewStringOperation().Get(cacheName).UnwrapOr("")
	if judgeCache != "" {
		return judgeCache
	}
	app.Redis.NewStringOperation().Set(cacheName, fmt.Sprintf("已在%v机器上锁成功", app.OutsideIp), goredis.WithExpire(time.Millisecond*time.Duration(info.Frequency)*3))
	return ""
}

// Unlock Lock 解锁
func (app *App) Unlock(info Task, id any) {
	cacheName := fmt.Sprintf("cron:%v:%v", info.Type, id)
	app.Redis.NewStringOperation().Del(cacheName)
}

// LockForever 永远上锁
func (app *App) LockForever(info Task, id any) string {
	cacheName := fmt.Sprintf("cron:%v:%v", info.Type, id)
	judgeCache := app.Redis.NewStringOperation().Get(cacheName).UnwrapOr("")
	if judgeCache != "" {
		return judgeCache
	}
	app.Redis.NewStringOperation().Set(cacheName, fmt.Sprintf("已在%v机器永远上锁成功", app.OutsideIp))
	return ""
}
