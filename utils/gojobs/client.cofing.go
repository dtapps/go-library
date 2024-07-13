package gojobs

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gorequest"
	"gorm.io/gorm"
)

// 设置配置信息
func (c *Client) setConfig(ctx context.Context, systemOutsideIP string) {
	c.config.systemInsideIP = gorequest.GetInsideIp(ctx)
	c.config.systemOutsideIP = systemOutsideIP
}

// ConfigGormClientFun GORM配置
func (c *Client) ConfigGormClientFun(ctx context.Context, client *gorm.DB, taskTableName string, taskLogStatus bool, taskLogTableName string) error {
	if client == nil {
		return errors.New("请配置 Gorm")
	}

	// 配置数据库
	c.gormConfig.client = client
	c.gormConfig.taskTableName = taskTableName
	if c.gormConfig.taskTableName == "" {
		return errors.New("请配置 Gorm 库名")
	}

	c.gormConfig.taskLogStatus = taskLogStatus
	if c.gormConfig.taskLogStatus {
		c.gormConfig.taskLogTableName = taskLogTableName
		if c.gormConfig.taskLogTableName == "" {
			return errors.New("请配置 Gorm 任务日志表名")
		}
	}

	return nil

	//err := c.gormAutoMigrateTask(ctx)
	//if err != nil {
	//	return err
	//}
	//err = c.gormAutoMigrateTaskLog(ctx)
	//
	//return err
}

// ConfigRedisClientFun REDIS配置
// lockKeyPrefix 锁Key前缀 xxx_lock
// lockKeySeparator 锁Key分隔符 :
// cornKeyPrefix 任务Key前缀 xxx_cron
// cornKeyCustom 任务Key自定义 xxx_cron_自定义  xxx_cron_自定义_*
func (c *Client) ConfigRedisClientFun(ctx context.Context, client *redis.Client, lockKeyPrefix string, lockKeySeparator string, cornKeyPrefix string, cornKeyCustom string) error {
	if client == nil {
		return errors.New("请配置 Redis")
	}

	// 配置缓存
	c.redisConfig.client = client

	// 配置缓存前缀
	c.redisConfig.lockKeyPrefix, c.redisConfig.lockKeySeparator, c.redisConfig.cornKeyPrefix, c.redisConfig.cornKeyCustom = lockKeyPrefix, lockKeySeparator, cornKeyPrefix, cornKeyCustom
	if c.redisConfig.lockKeyPrefix == "" || c.redisConfig.lockKeySeparator == "" || c.redisConfig.cornKeyPrefix == "" || c.redisConfig.cornKeyCustom == "" {
		return errors.New("请配置 Redis 前缀")
	}

	return nil
}
