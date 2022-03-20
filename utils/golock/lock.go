package golock

import (
	"context"
	"dtapps/dta/global"
	"dtapps/dta/library/utils/gouuid"
	"github.com/go-redis/redis/v8"
	"time"
)

type lock struct {
	key        string
	expiration time.Duration
	requestId  string
}

func NewLock(key string, expiration time.Duration) *lock {
	requestId := gouuid.GetUuId()
	return &lock{key: key, expiration: expiration, requestId: requestId}
}

// Get 获取锁
func (lk *lock) Get() bool {

	cxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	ok, err := global.GvaRedis.Db.SetNX(cxt, lk.key, lk.requestId, lk.expiration).Result()

	if err != nil {

		return false
	}

	return ok
}

// Release 释放锁
func (lk *lock) Release() error {

	cxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	const luaScript = `
	if redis.call('get', KEYS[1])==ARGV[1] then
		return redis.call('del', KEYS[1])
	else
		return 0
	end
	`

	script := redis.NewScript(luaScript)

	_, err := script.Run(cxt, global.GvaRedis.Db, []string{lk.key}, lk.requestId).Result()
	return err
}
