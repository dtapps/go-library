package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goarray"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/golock"
	"log"
	"runtime"
)

type JobsGormConfig struct {
	GormClient       *dorm.GormClient  // 数据库驱动
	RedisClient      *dorm.RedisClient // 缓存数据库驱动
	CurrentIp        string            // 当前ip
	LockKeyPrefix    string            // 锁Key前缀 xxx_lock
	LockKeySeparator string            // 锁Key分隔符 :
	CornKeyPrefix    string            // 任务Key前缀 xxx_cron
	CornKeyCustom    string            // 任务Key自定义 xxx_cron_自定义  xxx_cron_自定义_*
	Debug            bool              // 调试
}

// JobsGorm Gorm数据库驱动
type JobsGorm struct {
	gormClient  *dorm.GormClient  // 数据库驱动
	redisClient *dorm.RedisClient // 缓存驱动
	lockClient  *golock.LockRedis // 锁驱动
	config      struct {
		debug            bool   // 调试
		runVersion       string // 运行版本
		os               string // 系统类型
		arch             string // 系统架构
		maxProCs         int    // CPU核数
		version          string // GO版本
		macAddrS         string // Mac地址
		insideIp         string // 内网ip
		outsideIp        string // 外网ip
		lockKeyPrefix    string // 锁Key前缀 xxx_lock
		lockKeySeparator string // 锁Key分隔符 :
		cornKeyPrefix    string // 任务Key前缀 xxx_cron
		cornKeyCustom    string // 任务Key自定义
	}
}

// NewJobsGorm 初始化
func NewJobsGorm(config *JobsGormConfig) (*JobsGorm, error) {

	// 判断
	if config.LockKeyPrefix == "" {
		return nil, errors.New("需要配置锁Key前缀")
	}
	if config.LockKeySeparator == "" {
		return nil, errors.New("需要配置锁Key分隔符")
	}
	if config.CornKeyPrefix == "" {
		return nil, errors.New("需要配置任务Key前缀")
	}
	if config.CornKeyCustom == "" {
		return nil, errors.New("需要配置任务Key自定义")
	}
	if config.CurrentIp == "" {
		return nil, errors.New("需要配置当前的IP")
	}
	if config.GormClient == nil {
		return nil, errors.New("需要配置数据库驱动")
	}
	if config.RedisClient == nil {
		return nil, errors.New("需要配置缓存数据库驱动")
	}

	c := &JobsGorm{}
	c.gormClient = config.GormClient
	c.redisClient = config.RedisClient
	c.config.outsideIp = config.CurrentIp
	c.config.lockKeyPrefix = config.LockKeyPrefix
	c.config.lockKeySeparator = config.LockKeySeparator
	c.config.cornKeyPrefix = config.CornKeyPrefix
	c.config.cornKeyCustom = config.CornKeyCustom
	c.config.debug = config.Debug

	// 锁
	c.lockClient = golock.NewLockRedis(c.redisClient)

	// 配置信息
	c.config.runVersion = go_library.Version()
	c.config.os = runtime.GOOS
	c.config.arch = runtime.GOARCH
	c.config.maxProCs = runtime.GOMAXPROCS(0)
	c.config.version = runtime.Version()
	c.config.macAddrS = goarray.TurnString(goip.GetMacAddr(context.Background()))
	c.config.insideIp = goip.GetInsideIp(context.Background())

	// 创建模型
	err := c.gormClient.Db.AutoMigrate(
		&jobs_gorm_model.Task{},
		&jobs_gorm_model.TaskLog{},
		&jobs_gorm_model.TaskLogRun{},
		&jobs_gorm_model.TaskIp{},
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("创建任务模型失败：%v\n", err))
	}

	if c.config.debug == true {
		log.Printf("JOBS配置：%+v\n", c.config)
	}

	return c, nil
}
