package gojobs

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"time"
)

// Lock 上锁
func (j *JobsGorm) Lock(info jobs_gorm_model.Task, id any) (string, error) {
	return j.lockClient.Lock(fmt.Sprintf("%s%s%v%s%v", j.config.lockKeyPrefix, j.config.lockKeySeparator, info.Type, j.config.lockKeySeparator, id), fmt.Sprintf("已在%s@%s机器上锁成功", j.config.insideIp, j.config.outsideIp), time.Duration(info.Frequency)*3*time.Second)
}

// Unlock Lock 解锁
func (j *JobsGorm) Unlock(info jobs_gorm_model.Task, id any) error {
	return j.lockClient.Unlock(fmt.Sprintf("%s%s%v%s%v", j.config.lockKeyPrefix, j.config.lockKeySeparator, info.Type, j.config.lockKeySeparator, id))
}

// LockForever 永远上锁
func (j *JobsGorm) LockForever(info jobs_gorm_model.Task, id any) (string, error) {
	return j.lockClient.LockForever(fmt.Sprintf("%s%s%v%s%v", j.config.lockKeyPrefix, j.config.lockKeySeparator, info.Type, j.config.lockKeySeparator, id), fmt.Sprintf("已在%s@%s机器永远上锁成功", j.config.insideIp, j.config.outsideIp))
}
