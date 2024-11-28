package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gostring"
	"go.dtapp.net/library/utils/gotime"
	"gorm.io/gorm"
)

// ConfigCreateWaitCustomId 创建正在运行任务
type ConfigCreateWaitCustomId struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	Spec           string   // cron表达式
	CustomID       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIP      string   // 指定IP
	CurrentIP      string   // 当前IP
}

// CreateWaitCustomId 创建正在运行任务
func (c *Client) CreateWaitCustomId(ctx context.Context, config *ConfigCreateWaitCustomId) error {
	if config.CurrentIP == "" {
		config.CurrentIP = c.config.systemOutsideIP
	}
	err := config.Tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Create(&GormModelTask{
			Status:         TASK_WAIT,
			Params:         config.Params,
			StatusDesc:     "首次添加等待任务",
			Frequency:      config.Frequency,
			Spec:           config.Spec,
			RunID:          gostring.GetUuId(),
			CustomID:       config.CustomID,
			CustomSequence: config.CustomSequence,
			Type:           config.Type,
			TypeName:       config.TypeName,
			CreatedIP:      config.CurrentIP,
			SpecifyIP:      config.SpecifyIP,
			UpdatedIP:      config.CurrentIP,
			NextRunTime:    gotime.Current().AfterSeconds(config.Frequency).Time,
		}).Error
	if err != nil {
		return fmt.Errorf("创建[%s@%s]任务失败：%s", config.CustomID, config.Type, err)
	}
	return nil
}
