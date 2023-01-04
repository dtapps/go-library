package gojobs

import (
	"context"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
)

// 创建模型
func (c *Client) autoMigrateTask(ctx context.Context) {
	err := c.gormClient.GetDb().AutoMigrate(&jobs_gorm_model.Task{})
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("创建模型：%s", err)
	}
}

// 创建模型
func (c *Client) autoMigrateTaskLog(ctx context.Context) {
	err := c.gormClient.GetDb().AutoMigrate(&jobs_gorm_model.TaskLog{})
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("创建模型：%s", err)
	}
}

// GormTaskLogDelete 删除
func (c *Client) GormTaskLogDelete(ctx context.Context, hour int64) error {
	err := c.gormClient.GetDb().Where("log_time < ?", gotime.Current().BeforeHour(hour).Format()).Delete(&jobs_gorm_model.TaskLog{}).Error
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("删除失败：%s", err)
	}
	return err
}

// TaskLogRecord 记录
func (c *Client) TaskLogRecord(ctx context.Context, task jobs_gorm_model.Task, taskResultCode int, taskResultDesc string) {
	runId := gotrace_id.GetTraceIdContext(ctx)
	c.GormTaskLogRecord(ctx, task, runId, taskResultCode, taskResultDesc)
}

// GormTaskLogRecord 记录
func (c *Client) GormTaskLogRecord(ctx context.Context, task jobs_gorm_model.Task, runId string, taskResultCode int, taskResultDesc string) {

	taskLog := jobs_gorm_model.TaskLog{
		TaskId:          task.Id,
		TaskRunId:       runId,
		TaskResultCode:  taskResultCode,
		TaskResultDesc:  taskResultDesc,
		SystemHostName:  c.config.systemHostname,
		SystemInsideIp:  c.config.systemInsideIp,
		SystemOs:        c.config.systemOs,
		SystemArch:      c.config.systemKernel,
		GoVersion:       c.config.goVersion,
		SdkVersion:      c.config.sdkVersion,
		SystemOutsideIp: c.config.systemOutsideIp,
	}
	err := c.gormClient.GetDb().Create(&taskLog).Error
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("记录失败：%s", err)
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("记录数据：%+v", taskLog)
	}

}
