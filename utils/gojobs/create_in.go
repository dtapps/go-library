package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gostring"
	"go.dtapp.net/library/utils/gotime"
	"gorm.io/gorm"
)

// ConfigCreateInCustomId 创建正在运行任务
type ConfigCreateInCustomId struct {
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

// CreateInCustomId 创建正在运行任务
func (c *Client) CreateInCustomId(ctx context.Context, config *ConfigCreateInCustomId) error {
	if config.CurrentIP == "" {
		config.CurrentIP = c.config.systemOutsideIP
	}
	err := config.Tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Create(&GormModelTask{
			Status:         TASK_IN,
			Params:         config.Params,
			StatusDesc:     "首次添加任务",
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

// ConfigCreateInCustomIdOnly 创建正在运行唯一任务
type ConfigCreateInCustomIdOnly struct {
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

// CreateInCustomIdOnly 创建正在运行唯一任务
func (c *Client) CreateInCustomIdOnly(ctx context.Context, config *ConfigCreateInCustomIdOnly) error {
	query := c.TaskTypeTakeIn(ctx, config.Tx, config.CustomID, config.Type)
	if query.ID != 0 {
		return fmt.Errorf("任务已存在")
	}
	if config.CurrentIP == "" {
		config.CurrentIP = c.config.systemOutsideIP
	}
	err := config.Tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Create(&GormModelTask{
			Status:         TASK_IN,
			Params:         config.Params,
			StatusDesc:     "首次添加任务",
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

// ConfigCreateInCustomIdMaxNumber 创建正在运行任务并限制数量
type ConfigCreateInCustomIdMaxNumber struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	Spec           string   // cron表达式
	MaxNumber      int64    // 最大次数
	CustomID       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIP      string   // 指定IP
	CurrentIP      string   // 当前IP
}

// CreateInCustomIdMaxNumber 创建正在运行任务并限制数量
func (c *Client) CreateInCustomIdMaxNumber(ctx context.Context, config *ConfigCreateInCustomIdMaxNumber) error {
	if config.CurrentIP == "" {
		config.CurrentIP = c.config.systemOutsideIP
	}
	err := config.Tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Create(&GormModelTask{
			Status:         TASK_IN,
			Params:         config.Params,
			StatusDesc:     "首次添加任务",
			Frequency:      config.Frequency,
			Spec:           config.Spec,
			MaxNumber:      config.MaxNumber,
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

// ConfigCreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
type ConfigCreateInCustomIdMaxNumberOnly struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	Spec           string   // cron表达式
	MaxNumber      int64    // 最大次数
	CustomID       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIP      string   // 指定IP
	CurrentIP      string   // 当前IP
}

// CreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
func (c *Client) CreateInCustomIdMaxNumberOnly(ctx context.Context, config *ConfigCreateInCustomIdMaxNumberOnly) error {
	query := c.TaskTypeTakeIn(ctx, config.Tx, config.CustomID, config.Type)
	if query.ID != 0 {
		err := fmt.Errorf("任务[%s@%s]已存在", config.CustomID, config.Type)
		return err
	}
	if config.CurrentIP == "" {
		config.CurrentIP = c.config.systemOutsideIP
	}
	err := config.Tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Create(&GormModelTask{
			Status:         TASK_IN,
			Params:         config.Params,
			StatusDesc:     "首次添加任务",
			Frequency:      config.Frequency,
			Spec:           config.Spec,
			MaxNumber:      config.MaxNumber,
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
