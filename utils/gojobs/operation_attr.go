package gojobs

import (
	"github.com/go-redis/redis/v8"
	"go.dtapp.net/library/utils/goip"
	"go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

const (
	lockTypeRedis = "redis"
	lockTypeEtcd  = "etcd"
)

// OperationAttr 操作属性
type OperationAttr struct {
	gormClient  *gorm.DB         // 数据库驱动
	redisClient *redis.Client    // 缓存数据库驱动
	etcdClient  *clientv3.Client // 分布式缓存驱动
	lockPrefix  string           // 锁Key前缀
	ipService   *goip.Client     // ip服务
	lockType    string           // 锁驱动类型
}

// WithGormClient 设置数据库驱动
func WithGormClient(client *gorm.DB) *OperationAttr {
	return &OperationAttr{gormClient: client}
}

// WithRedisClient 设置缓存数据库驱动
func WithRedisClient(redisClient *redis.Client) *OperationAttr {
	return &OperationAttr{redisClient: redisClient, lockType: lockTypeRedis}
}

// WithEtcdClient 设置分布式缓存驱动
func WithEtcdClient(etcdClient *clientv3.Client) *OperationAttr {
	return &OperationAttr{etcdClient: etcdClient, lockType: lockTypeEtcd}
}

// WithLockPrefix 设置锁Key前缀
// redis：fmt.Sprintf("cron:lock:%v:%v", info.Type, id)
// etcd：fmt.Sprintf("cron/lock/%v/%v", info.Type, id)
func WithLockPrefix(lockPrefix string) *OperationAttr {
	return &OperationAttr{lockPrefix: lockPrefix}
}

// WithIpService 设置ip服务
func WithIpService(ipService *goip.Client) *OperationAttr {
	return &OperationAttr{ipService: ipService}
}
