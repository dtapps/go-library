package golock

import (
	"go.dtapp.net/library/utils/goredis"
	"time"
)

type ConfigLockRedis struct {
	Key            string
	KeyContent     string
	ExpirationTime time.Duration
}

type LockRedis struct {
	config ConfigLockRedis
	db     goredis.App
}

func NewLockRedis(db goredis.App) *LockRedis {
	return &LockRedis{db: db}
}

// Lock 上锁
func (lockRedis *LockRedis) Lock() bool {
	judgeCache := lockRedis.db.NewStringOperation().Get(lockRedis.config.Key).UnwrapOr("")
	if judgeCache != "" {
		return true
	}
	lockRedis.db.NewStringOperation().Set(lockRedis.config.Key, lockRedis.config.KeyContent, goredis.WithExpire(lockRedis.config.ExpirationTime))
	return true
}

// Unlock Lock 解锁
func (lockRedis *LockRedis) Unlock() {
	lockRedis.db.NewStringOperation().Del(lockRedis.config.Key)
}

// LockForever 永远上锁
func (lockRedis *LockRedis) LockForever() bool {
	judgeCache := lockRedis.db.NewStringOperation().Get(lockRedis.config.Key).UnwrapOr("")
	if judgeCache != "" {
		return true
	}
	lockRedis.db.NewStringOperation().Set(lockRedis.config.Key, lockRedis.config.KeyContent)
	return true
}