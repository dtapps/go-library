package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gostring"
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
)

// ConfigCreateInCustomId 创建正在运行任务
type ConfigCreateInCustomId struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前外网IP
}

// CreateInCustomId 创建正在运行任务
func (c *Client) CreateInCustomId(ctx context.Context, config *ConfigCreateInCustomId) error {
	if config.CurrentIp == "" {
		config.CurrentIp = c.config.systemOutsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
		NextRunTime:    gotime.Current().AfterSeconds(config.Frequency).Time,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}

// ConfigCreateInCustomIdOnly 创建正在运行唯一任务
type ConfigCreateInCustomIdOnly struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前外网IP
}

// CreateInCustomIdOnly 创建正在运行唯一任务
func (c *Client) CreateInCustomIdOnly(ctx context.Context, config *ConfigCreateInCustomIdOnly) error {
	query := c.TaskTypeTakeIn(ctx, config.Tx, config.CustomId, config.Type)
	if query.Id != 0 {
		return TaskIsExist
	}
	if config.CurrentIp == "" {
		config.CurrentIp = c.config.systemOutsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
		NextRunTime:    gotime.Current().AfterSeconds(config.Frequency).Time,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}

// ConfigCreateInCustomIdMaxNumber 创建正在运行任务并限制数量
type ConfigCreateInCustomIdMaxNumber struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	MaxNumber      int64    // 最大次数
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前外网IP
}

// CreateInCustomIdMaxNumber 创建正在运行任务并限制数量
func (c *Client) CreateInCustomIdMaxNumber(ctx context.Context, config *ConfigCreateInCustomIdMaxNumber) error {
	if config.CurrentIp == "" {
		config.CurrentIp = c.config.systemOutsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		MaxNumber:      config.MaxNumber,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
		NextRunTime:    gotime.Current().AfterSeconds(config.Frequency).Time,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}

// ConfigCreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
type ConfigCreateInCustomIdMaxNumberOnly struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	MaxNumber      int64    // 最大次数
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前外网IP
}

// CreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
func (c *Client) CreateInCustomIdMaxNumberOnly(ctx context.Context, config *ConfigCreateInCustomIdMaxNumberOnly) error {
	query := c.TaskTypeTakeIn(ctx, config.Tx, config.CustomId, config.Type)
	if query.Id != 0 {
		return TaskIsExist
	}
	if config.CurrentIp == "" {
		config.CurrentIp = c.config.systemOutsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		MaxNumber:      config.MaxNumber,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
		NextRunTime:    gotime.Current().AfterSeconds(config.Frequency).Time,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}
