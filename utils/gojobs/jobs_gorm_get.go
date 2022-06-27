package gojobs

import (
	"github.com/go-redis/redis/v8"
	"go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

// GetDb 数据库驱动
func (j *JobsGorm) GetDb() *gorm.DB {
	return j.service.gormClient
}

// GetRedis 缓存数据库驱动
func (j *JobsGorm) GetRedis() *redis.Client {
	return j.db.redisClient
}

// GetEtcd 分布式缓存驱动
func (j *JobsGorm) GetEtcd() *clientv3.Client {
	return j.db.etcdClient
}
