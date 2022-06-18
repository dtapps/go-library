package gojobs

import (
	"fmt"
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/gojobs/jobs_gorm_model"
	"time"
)

// Lock 上锁
func (j *JobsGorm) Lock(info jobs_gorm_model.Task, id any) string {
	cacheName := fmt.Sprintf("cron:%v:%v", info.Type, id)
	judgeCache := j.redis.NewStringOperation().Get(cacheName).UnwrapOr("")
	if judgeCache != "" {
		return judgeCache
	}
	j.redis.NewStringOperation().Set(cacheName, fmt.Sprintf("已在%v机器上锁成功", j.outsideIp), dorm.WithExpire(time.Millisecond*time.Duration(info.Frequency)*3))
	return ""
}

// Unlock Lock 解锁
func (j *JobsGorm) Unlock(info jobs_gorm_model.Task, id any) {
	cacheName := fmt.Sprintf("cron:%v:%v", info.Type, id)
	j.redis.NewStringOperation().Del(cacheName)
}

// LockForever 永远上锁
func (j *JobsGorm) LockForever(info jobs_gorm_model.Task, id any) string {
	cacheName := fmt.Sprintf("cron:%v:%v", info.Type, id)
	judgeCache := j.redis.NewStringOperation().Get(cacheName).UnwrapOr("")
	if judgeCache != "" {
		return judgeCache
	}
	j.redis.NewStringOperation().Set(cacheName, fmt.Sprintf("已在%v机器永远上锁成功", j.outsideIp))
	return ""
}
