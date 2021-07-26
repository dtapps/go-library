package v20210726

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	Rdb  *redis.Client
	RdbC *redis.ClusterClient
)

// InitRedis 初始化连接 普通连接
func InitRedis(host string, port int, password string, db int) (err error) {
	dsn := fmt.Sprintf("%s:%v", host, port)
	fmt.Printf("【redis.普通】数据库配置 %s \n", dsn)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: password, // no password set
		DB:       db,       // use default DB
		PoolSize: 100,      // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}

// InitSentinelRedis 初始化连接 哨兵模式
func InitSentinelRedis(adds []string, masterName string, password string, db int) (err error) {
	fmt.Printf("【redis.哨兵】数据库配置 %s \n", adds)
	Rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterName,
		SentinelAddrs: adds,
		Password:      password, // no password set
		DB:            db,       // use default DB
		PoolSize:      100,      // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}

// InitClusterRedis 初始化连接 集群
func InitClusterRedis(adds []string, password string) (err error) {
	fmt.Printf("【redis.集群】数据库配置 %v \n", adds)
	RdbC = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    adds,
		Password: password, // no password set
		PoolSize: 100,      // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RdbC.Ping(ctx).Result()
	return err
}
