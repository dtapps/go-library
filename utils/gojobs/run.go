package gojobs

import (
	"context"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
)

// Run 运行
func (c *Client) Run(ctx context.Context, task jobs_gorm_model.Task, taskResultCode int, taskResultDesc string) {

	runId := gotrace_id.GetTraceIdContext(ctx)
	if runId == "" {
		c.zapLog.WithTraceId(ctx).Sugar().Error("上下文没有跟踪编号")
		return
	}

	c.GormTaskLogRecord(ctx, task, runId, taskResultCode, taskResultDesc)
	if c.mongoConfig.stats {
		c.MongoTaskLogRecord(ctx, task, runId, taskResultCode, taskResultDesc)
	}

	switch taskResultCode {
	case 0:
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("run_id", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				RunId:       runId,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存失败：%s", err.Error())
		}
		return
	case CodeSuccess:
		// 执行成功
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("status_desc", "number", "run_id", "updated_ip", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				StatusDesc:  "执行成功",
				Number:      task.Number + 1,
				RunId:       runId,
				UpdatedIp:   c.config.systemOutsideIp,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存失败：%s", err.Error())
		}
	case CodeEnd:
		// 执行成功、提前结束
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("status", "status_desc", "number", "updated_ip", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				Status:      TASK_SUCCESS,
				StatusDesc:  "结束执行",
				Number:      task.Number + 1,
				UpdatedIp:   c.config.systemOutsideIp,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().Time,
			}).Error
		if err != nil {
			c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存失败：%s", err.Error())
		}
	case CodeError:
		// 执行失败
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("status_desc", "number", "run_id", "updated_ip", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				StatusDesc:  "执行失败",
				Number:      task.Number + 1,
				RunId:       runId,
				UpdatedIp:   c.config.systemOutsideIp,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存失败：%s", err.Error())
		}
	}

	if task.MaxNumber != 0 {
		if task.Number+1 >= task.MaxNumber {
			// 关闭执行
			err := c.EditTask(c.gormClient.GetDb(), task.Id).
				Select("status").
				Updates(jobs_gorm_model.Task{
					Status: TASK_TIMEOUT,
				}).Error
			if err != nil {
				c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存失败：%s", err.Error())
			}
		}
	}
	return
}
