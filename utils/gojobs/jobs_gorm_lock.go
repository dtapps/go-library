package gojobs

import (
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gojobs/jobs_gorm_model"
)

// Lock 上锁
func (j *JobsGorm) Lock(info jobs_gorm_model.Task, id any) (string, error) {

	if j.config.lockType == "" {
		return "", errors.New("没有配置")
	}

	var (
		redisKey = fmt.Sprintf("%s:%v:%v", j.config.lockPrefix, info.Type, id)
		etcdKey  = fmt.Sprintf("%s/%v/%v", j.config.lockPrefix, info.Type, id)
		val      = fmt.Sprintf("已在%s@%s机器上锁成功", j.config.insideIp, j.config.outsideIp)
		ttl      = info.Frequency * 3
	)

	if j.config.lockType == lockTypeRedis {
		return j.service.lockRedisClient.Lock(redisKey, val, ttl)
	}

	return j.service.lockEtcdClient.Lock(etcdKey, val, ttl)
}

// Unlock Lock 解锁
func (j *JobsGorm) Unlock(info jobs_gorm_model.Task, id any) error {

	if j.config.lockType == "" {
		return errors.New("没有配置")
	}

	var (
		redisKey = fmt.Sprintf("%s:%v:%v", j.config.lockPrefix, info.Type, id)
		etcdKey  = fmt.Sprintf("%s/%v/%v", j.config.lockPrefix, info.Type, id)
	)

	if j.config.lockType == lockTypeRedis {
		return j.service.lockRedisClient.Unlock(redisKey)
	}

	return j.service.lockEtcdClient.Unlock(etcdKey)
}

// LockForever 永远上锁
func (j *JobsGorm) LockForever(info jobs_gorm_model.Task, id any) (string, error) {

	if j.config.lockType == "" {
		return "", errors.New("没有配置")
	}

	var (
		redisKey = fmt.Sprintf("%s:%v:%v", j.config.lockPrefix, info.Type, id)
		etcdKey  = fmt.Sprintf("%s/%v/%v", j.config.lockPrefix, info.Type, id)
		val      = fmt.Sprintf("已在%s@%s机器永远上锁成功", j.config.insideIp, j.config.outsideIp)
	)

	if j.config.lockType == lockTypeRedis {
		return j.service.lockRedisClient.LockForever(redisKey, val)
	}

	return j.service.lockEtcdClient.LockForever(etcdKey, val)
}
