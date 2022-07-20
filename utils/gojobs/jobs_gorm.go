package gojobs

import (
	"errors"
	"fmt"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/goarray"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/golock"
	"github.com/go-redis/redis/v8"
	"go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"runtime"
)

// JobsGorm Gorm数据库驱动
type JobsGorm struct {
	db struct {
		gormClient  *gorm.DB         // 数据库驱动
		redisClient *redis.Client    // 缓存数据库驱动
		etcdClient  *clientv3.Client // 分布式缓存驱动
	}
	service struct {
		gormClient      *gorm.DB          // 数据库驱动
		lockRedisClient *golock.LockRedis // 缓存数据库驱动
		lockEtcdClient  *golock.LockEtcd  // 分布式缓存驱动
	} // 服务
	config struct {
		lockPrefix string // 锁Key前缀
		lockType   string // 锁驱动类型
		runVersion string // 运行版本
		os         string // 系统类型
		arch       string // 系统架构
		maxProCs   int    // CPU核数
		version    string // GO版本
		macAddrS   string // Mac地址
		insideIp   string // 内网ip
		outsideIp  string // 外网ip
	} // 配置
}

// NewJobsGorm 初始化
// WithGormClient && WithRedisClient && WithLockPrefix && WithOutsideIp
// WithGormClient && WithEtcdClient && WithLockPrefix && WithOutsideIp
func NewJobsGorm(attrs ...*OperationAttr) (*JobsGorm, error) {

	c := &JobsGorm{}
	for _, attr := range attrs {
		if attr.gormClient != nil {
			c.db.gormClient = attr.gormClient
			c.service.gormClient = attr.gormClient
		}
		if attr.redisClient != nil {
			c.db.redisClient = attr.redisClient
			c.config.lockType = attr.lockType
		}
		if attr.etcdClient != nil {
			c.db.etcdClient = attr.etcdClient
			c.config.lockType = attr.lockType
		}
		if attr.lockPrefix != "" {
			c.config.lockPrefix = attr.lockPrefix
		}
		if attr.outsideIp != "" {
			c.config.outsideIp = attr.outsideIp
		}
	}

	if c.config.outsideIp == "" {
		return nil, errors.New("需要配置当前的IP")
	}

	if c.db.gormClient == nil {
		return nil, errors.New("需要配置数据库驱动")
	}

	c.config.runVersion = go_library.Version()
	c.config.os = runtime.GOOS
	c.config.arch = runtime.GOARCH
	c.config.maxProCs = runtime.GOMAXPROCS(0)
	c.config.version = runtime.Version()
	c.config.macAddrS = goarray.TurnString(goip.GetMacAddr())
	c.config.insideIp = goip.GetInsideIp()

	switch c.config.lockType {
	case lockTypeRedis:

		if c.db.redisClient == nil {
			return nil, errors.New("没有设置REDIS驱动")
		}
		c.service.lockRedisClient = golock.NewLockRedis(c.db.redisClient)

	case lockTypeEtcd:

		if c.db.etcdClient == nil {
			return nil, errors.New("没有设置ETCD驱动")
		}
		c.service.lockEtcdClient = golock.NewLockEtcd(c.db.etcdClient)

	default:
		// 添加任务端不需要
		// return nil, errors.New("驱动为空")
	}

	err := c.service.gormClient.AutoMigrate(
		&jobs_gorm_model.Task{},
		&jobs_gorm_model.TaskLog{},
		&jobs_gorm_model.TaskLogRun{},
		&jobs_gorm_model.TaskIp{},
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("创建任务模型失败：%v\n", err))
	}

	return c, nil
}
