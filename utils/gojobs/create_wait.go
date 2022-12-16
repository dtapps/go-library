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

// ConfigCreateWaitCustomId 创建正在运行任务
type ConfigCreateWaitCustomId struct {
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

// CreateWaitCustomId 创建正在运行任务
func (c *Client) CreateWaitCustomId(ctx context.Context, config *ConfigCreateWaitCustomId) error {
	if config.CurrentIp == "" {
		config.CurrentIp = c.config.systemOutsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_WAIT,
		Params:         config.Params,
		StatusDesc:     "首次添加等待任务",
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
