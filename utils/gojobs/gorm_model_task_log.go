package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gotime"
	"go.opentelemetry.io/otel/codes"
	"time"
)

// GormModelTaskLog 任务日志
type GormModelTaskLog struct {
	LogID           uint      `gorm:"primaryKey;comment:【日志】编号" json:"log_id"`                            // 【日志】编号
	LogTime         time.Time `gorm:"autoCreateTime;index;comment:【日志】时间" json:"log_time"`                // 【日志】时间
	TaskID          uint      `gorm:"index;comment:【任务】编号" json:"task_id"`                                // 【任务】编号
	TaskRunID       string    `gorm:"comment:【任务】执行编号" json:"task_run_id"`                                //【任务】执行编号
	TaskResultCode  int       `gorm:"index;comment:【任务】执行状态码" json:"task_result_code"`                    //【任务】执行状态码
	TaskResultDesc  string    `gorm:"comment:【任务】执行结果" json:"task_result_desc"`                           //【任务】执行结果
	SystemInsideIP  string    `gorm:"default:0.0.0.0;comment:【系统】内网IP" json:"system_inside_ip,omitempty"` //【系统】内网IP
	SystemOutsideIP string    `gorm:"default:0.0.0.0;comment:【系统】外网IP" json:"system_outside_ip"`          //【系统】外网IP
}

// 创建模型
func (c *Client) gormAutoMigrateTaskLog(ctx context.Context) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		AutoMigrate(&GormModelTaskLog{})
	if err != nil {
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogDelete 删除
func (c *Client) GormTaskLogDelete(ctx context.Context, hour int64) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).
		Delete(&GormModelTaskLog{}).Error
	if err != nil {
		err = fmt.Errorf("删除失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogInDelete 删除任务运行
func (c *Client) GormTaskLogInDelete(ctx context.Context, hour int64) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Where("task_result_status = ?", TASK_IN).Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).
		Delete(&GormModelTaskLog{}).Error
	if err != nil {
		err = fmt.Errorf("删除失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogSuccessDelete 删除任务完成
func (c *Client) GormTaskLogSuccessDelete(ctx context.Context, hour int64) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Where("task_result_status = ?", TASK_SUCCESS).Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).
		Delete(&GormModelTaskLog{}).Error
	if err != nil {
		err = fmt.Errorf("删除失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogErrorDelete 删除任务异常
func (c *Client) GormTaskLogErrorDelete(ctx context.Context, hour int64) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Where("task_result_status = ?", TASK_ERROR).Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).
		Delete(&GormModelTaskLog{}).Error
	if err != nil {
		err = fmt.Errorf("删除失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogTimeoutDelete 删除任务超时
func (c *Client) GormTaskLogTimeoutDelete(ctx context.Context, hour int64) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Where("task_result_status = ?", TASK_TIMEOUT).Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).
		Delete(&GormModelTaskLog{}).Error
	if err != nil {
		err = fmt.Errorf("删除失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogWaitDelete 删除任务等待
func (c *Client) GormTaskLogWaitDelete(ctx context.Context, hour int64) error {
	if c.gormConfig.taskLogStatus == false {
		return nil
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Where("task_result_status = ?", TASK_WAIT).
		Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).
		Delete(&GormModelTaskLog{}).Error
	if err != nil {
		err = fmt.Errorf("删除失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
	return err
}

// GormTaskLogRecord 记录
func (c *Client) GormTaskLogRecord(ctx context.Context, task GormModelTask, runId string, taskResultCode int, taskResultDesc string) {
	taskLog := GormModelTaskLog{
		TaskID:          task.ID,                  //【任务】编号
		TaskRunID:       runId,                    //【任务】执行编号
		TaskResultCode:  taskResultCode,           //【任务】执行状态码
		TaskResultDesc:  taskResultDesc,           //【任务】执行结果
		SystemInsideIP:  c.config.systemInsideIP,  //【系统】内网IP
		SystemOutsideIP: c.config.systemOutsideIP, //【系统】外网IP
	}
	err := c.gormConfig.client.WithContext(ctx).Table(c.gormConfig.taskLogTableName).
		Create(&taskLog).Error
	if err != nil {
		err = fmt.Errorf("记录失败：%s", err)
		TraceRecordError(ctx, err)
		TraceSetStatus(ctx, codes.Error, err.Error())
	}
}
