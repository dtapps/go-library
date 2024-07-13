package gojobs

import (
	"context"
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gotime"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
	"time"
)

// GormModelTask 任务
type GormModelTask struct {
	ID             uint           `gorm:"primaryKey;comment:记录编号" json:"id,omitempty"`                      // 记录编号
	Status         string         `gorm:"index;comment:状态码" json:"status,omitempty"`                        // 状态码
	StatusDesc     string         `gorm:"comment:状态描述" json:"status_desc,omitempty"`                        // 状态描述
	Params         string         `gorm:"comment:参数" json:"params,omitempty"`                               // 参数
	Frequency      int64          `gorm:"index;comment:频率(秒单位)" json:"frequency,omitempty"`                 // 频率(秒单位)
	Spec           string         `gorm:"index;comment:cron表达式" json:"spec,omitempty"`                      // cron表达式
	Number         int64          `gorm:"comment:当前次数" json:"number,omitempty"`                             // 当前次数
	MaxNumber      int64          `gorm:"comment:最大次数" json:"max_number,omitempty"`                         // 最大次数
	RunID          string         `gorm:"comment:执行编号" json:"run_id,omitempty"`                             // 执行编号
	CustomID       string         `gorm:"index;comment:自定义编号" json:"custom_id,omitempty"`                   // 自定义编号
	CustomSequence int64          `gorm:"comment:自定义顺序" json:"custom_sequence,omitempty"`                   // 自定义顺序
	Type           string         `gorm:"index;comment:类型" json:"type,omitempty"`                           // 类型
	TypeName       string         `gorm:"comment:类型名称" json:"type_name,omitempty"`                          // 类型名称
	CreatedIP      string         `gorm:"default:0.0.0.0;comment:创建外网IP" json:"created_ip,omitempty"`       // 创建外网IP
	SpecifyIP      string         `gorm:"default:0.0.0.0;index;comment:指定外网IP" json:"specify_ip,omitempty"` // 指定外网IP
	UpdatedIP      string         `gorm:"default:0.0.0.0;comment:更新外网IP" json:"updated_ip,omitempty"`       // 更新外网IP
	Result         string         `gorm:"comment:结果" json:"result,omitempty"`                               // 结果
	NextRunTime    time.Time      `gorm:"comment:下次运行时间" json:"next_run_time,omitempty"`                    // 下次运行时间
	CreatedAt      time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at,omitempty"`          // 创建时间
	UpdatedAt      time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at,omitempty"`          // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at,omitempty"`                   // 删除时间
}

// TaskTakeId 编号查询任务
func (c *Client) TaskTakeId(ctx context.Context, tx *gorm.DB, id uint) (result GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("id = ?", id).
		Take(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result
	}
	if err != nil {
		err = fmt.Errorf("编号查询任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return result
}

// TaskTake 自定义编号查询任务
func (c *Client) TaskTake(ctx context.Context, tx *gorm.DB, customId string) (result GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("custom_id = ?", customId).
		Take(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result
	}
	if err != nil {
		err = fmt.Errorf("自定义编号查询任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return result
}

// 自定义编号加状态查询任务
func (c *Client) taskTake(ctx context.Context, tx *gorm.DB, customId, status string) (result GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("custom_id = ?", customId).
		Where("status = ?", status).
		Take(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result
	}
	if err != nil {
		err = fmt.Errorf("自定义编号加状态查询任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return result
}

// TaskTakeIn 查询单任务 - 任务运行
func (c *Client) TaskTakeIn(ctx context.Context, tx *gorm.DB, customId string) GormModelTask {
	return c.taskTake(ctx, tx, customId, TASK_IN)
}

// TaskTakeSuccess 查询单任务 - 任务完成
func (c *Client) TaskTakeSuccess(ctx context.Context, tx *gorm.DB, customId string) GormModelTask {
	return c.taskTake(ctx, tx, customId, TASK_SUCCESS)
}

// TaskTakeError 查询单任务 - 任务异常
func (c *Client) TaskTakeError(ctx context.Context, tx *gorm.DB, customId string) GormModelTask {
	return c.taskTake(ctx, tx, customId, TASK_ERROR)
}

// TaskTakeTimeout 查询单任务 - 任务超时
func (c *Client) TaskTakeTimeout(ctx context.Context, tx *gorm.DB, customId string) GormModelTask {
	return c.taskTake(ctx, tx, customId, TASK_TIMEOUT)
}

// TaskTakeWait 查询单任务 - 任务等待
func (c *Client) TaskTakeWait(ctx context.Context, tx *gorm.DB, customId string) GormModelTask {
	return c.taskTake(ctx, tx, customId, TASK_WAIT)
}

// TaskTypeTake 查询单任务
func (c *Client) TaskTypeTake(ctx context.Context, tx *gorm.DB, customId, Type string) (result GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("custom_id = ?", customId).
		Where("type = ?", Type).
		Take(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result
	}
	if err != nil {
		err = fmt.Errorf("查询单任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return result
}

// 查询单任务
func (c *Client) taskTypeTake(ctx context.Context, tx *gorm.DB, customId, Type, status string) (result GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("custom_id = ?", customId).Where("type = ?", Type).
		Where("status = ?", status).
		Take(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result
	}
	if err != nil {
		err = fmt.Errorf("查询单任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return result
}

// TaskTypeTakeIn 查询单任务 - 任务运行
func (c *Client) TaskTypeTakeIn(ctx context.Context, tx *gorm.DB, customId, Type string) GormModelTask {
	return c.taskTypeTake(ctx, tx, customId, Type, TASK_IN)
}

// TaskTypeTakeSuccess 查询单任务 - 任务完成
func (c *Client) TaskTypeTakeSuccess(ctx context.Context, tx *gorm.DB, customId, Type string) GormModelTask {
	return c.taskTypeTake(ctx, tx, customId, Type, TASK_SUCCESS)
}

// TaskTypeTakeError 查询单任务 - 任务异常
func (c *Client) TaskTypeTakeError(ctx context.Context, tx *gorm.DB, customId, Type string) GormModelTask {
	return c.taskTypeTake(ctx, tx, customId, Type, TASK_ERROR)
}

// TaskTypeTakeTimeout 查询单任务 - 任务超时
func (c *Client) TaskTypeTakeTimeout(ctx context.Context, tx *gorm.DB, customId, Type string) GormModelTask {
	return c.taskTypeTake(ctx, tx, customId, Type, TASK_TIMEOUT)
}

// TaskTypeTakeWait 查询单任务 - 任务等待
func (c *Client) TaskTypeTakeWait(ctx context.Context, tx *gorm.DB, customId, Type string) GormModelTask {
	return c.taskTypeTake(ctx, tx, customId, Type, TASK_WAIT)
}

// TaskFindAll 查询多任务
func (c *Client) TaskFindAll(ctx context.Context, tx *gorm.DB, frequency int64) (results []GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("frequency = ?", frequency).
		Order("id asc").
		Find(&results).Error
	if err != nil {
		err = fmt.Errorf("查询多任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return results
}

// TaskFindAllType 查询多任务
func (c *Client) TaskFindAllType(ctx context.Context, tx *gorm.DB, Type string, frequency int64) (results []GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("type = ?", Type).
		Where("frequency = ?", frequency).
		Order("id asc").
		Find(&results).Error
	if err != nil {
		err = fmt.Errorf("查询多任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return results
}

// 查询多任务
func (c *Client) taskFindAll(ctx context.Context, tx *gorm.DB, frequency int64, status string) (results []GormModelTask) {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("frequency = ?", frequency).
		Where("status = ?", status).
		Order("id asc").
		Find(&results).Error
	if err != nil {
		err = fmt.Errorf("查询多任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return results
}

// 查询多任务
func (c *Client) taskFindAllType(ctx context.Context, tx *gorm.DB, Type string, frequency int64, status string) (results []GormModelTask) {
	if frequency == 0 {
		err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
			Where("type = ?", Type).
			Where("status = ?", status).
			Order("id asc").
			Find(&results).Error
		if err != nil {
			err = fmt.Errorf("查询多任务：%v", err)
			TraceRecordError(ctx, err)
			TraceSetStatus(ctx, codes.Error, err.Error())
		}
		return results
	}
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("type = ?", Type).
		Where("frequency = ?", frequency).
		Where("status = ?", status).
		Order("id asc").
		Find(&results).Error
	if err != nil {
		err = fmt.Errorf("查询多任务：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return results
}

// TaskFindAllIn 查询多任务 - 任务运行
func (c *Client) TaskFindAllIn(ctx context.Context, tx *gorm.DB, frequency int64) []GormModelTask {
	return c.taskFindAll(ctx, tx, frequency, TASK_IN)
}

// TaskFindAllInType 查询多任务 - 任务运行
func (c *Client) TaskFindAllInType(ctx context.Context, tx *gorm.DB, Type string) []GormModelTask {
	return c.taskFindAllType(ctx, tx, Type, 0, TASK_IN)
}

// TaskFindAllSuccess 查询多任务 - 任务完成
func (c *Client) TaskFindAllSuccess(ctx context.Context, tx *gorm.DB, frequency int64) []GormModelTask {
	return c.taskFindAll(ctx, tx, frequency, TASK_SUCCESS)
}

// TaskFindAllSuccessType 查询多任务 - 任务完成
func (c *Client) TaskFindAllSuccessType(ctx context.Context, tx *gorm.DB, Type string) []GormModelTask {
	return c.taskFindAllType(ctx, tx, Type, 0, TASK_SUCCESS)
}

// TaskFindAllError 查询多任务 - 任务异常
func (c *Client) TaskFindAllError(ctx context.Context, tx *gorm.DB, frequency int64) []GormModelTask {
	return c.taskFindAll(ctx, tx, frequency, TASK_ERROR)
}

// TaskFindAllErrorType 查询多任务 - 任务异常
func (c *Client) TaskFindAllErrorType(ctx context.Context, tx *gorm.DB, Type string) []GormModelTask {
	return c.taskFindAllType(ctx, tx, Type, 0, TASK_ERROR)
}

// TaskFindAllTimeout 查询多任务 - 任务超时
func (c *Client) TaskFindAllTimeout(ctx context.Context, tx *gorm.DB, frequency int64) []GormModelTask {
	return c.taskFindAll(ctx, tx, frequency, TASK_TIMEOUT)
}

// TaskFindAllTimeoutType 查询多任务 - 任务超时
func (c *Client) TaskFindAllTimeoutType(ctx context.Context, tx *gorm.DB, Type string) []GormModelTask {
	return c.taskFindAllType(ctx, tx, Type, 0, TASK_TIMEOUT)
}

// TaskFindAllWait 查询多任务 - 任务等待
func (c *Client) TaskFindAllWait(ctx context.Context, tx *gorm.DB, frequency int64) []GormModelTask {
	return c.taskFindAll(ctx, tx, frequency, TASK_WAIT)
}

// TaskFindAllWaitType 查询多任务 - 任务等待
func (c *Client) TaskFindAllWaitType(ctx context.Context, tx *gorm.DB, Type string) []GormModelTask {
	return c.taskFindAllType(ctx, tx, Type, 0, TASK_WAIT)
}

// StartTask 任务启动
func (c *Client) StartTask(ctx context.Context, tx *gorm.DB, id uint) error {
	err := c.EditTask(ctx, tx, id).
		Select("status", "status_desc").
		Updates(GormModelTask{
			Status:     TASK_IN,
			StatusDesc: "启动任务",
		}).Error
	if err != nil {
		err = fmt.Errorf("任务启动失败：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// StartTaskCustom 任务启动自定义
func (c *Client) StartTaskCustom(ctx context.Context, tx *gorm.DB, customId string, customSequence int64) error {
	err := tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("custom_id = ?", customId).
		Where("custom_sequence = ?", customSequence).
		Where("status = ?", TASK_WAIT).
		Select("status", "status_desc").
		Updates(GormModelTask{
			Status:     TASK_IN,
			StatusDesc: "启动任务",
		}).Error
	if err != nil {
		err = fmt.Errorf("任务启动自定义失败：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// EditTask 任务修改
func (c *Client) EditTask(ctx context.Context, tx *gorm.DB, id uint) *gorm.DB {
	return tx.WithContext(ctx).Table(c.gormConfig.taskTableName).
		Where("id = ?", id)
}

// UpdateFrequency 更新任务频率
func (c *Client) UpdateFrequency(ctx context.Context, tx *gorm.DB, id uint, frequency int64) error {
	err := c.EditTask(ctx, tx, id).
		Select("frequency", "next_run_time").
		Updates(GormModelTask{
			Frequency:   frequency,
			NextRunTime: gotime.Current().AfterSeconds(frequency).Time,
		}).Error
	if err != nil {
		err = fmt.Errorf("更新任务频率失败：%v", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}
