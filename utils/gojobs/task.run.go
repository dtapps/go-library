package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"log/slog"
)

// Run 运行
func (c *Client) Run(ctx context.Context, task GormModelTask, taskResultCode int, taskResultDesc string) {

	runId := gorequest.GetRequestIDContext(ctx)
	if runId == "" {
		slog.InfoContext(ctx, "上下文没有跟踪编号")
		return
	}

	if c.gormConfig.taskLogStatus {
		go c.GormTaskLogRecord(ctx, task, runId, taskResultCode, taskResultDesc)
	}

	switch taskResultCode {
	case 0:
		err := c.EditTask(ctx, c.gormConfig.client, task.ID).
			Select("run_id", "result", "next_run_time").
			Updates(GormModelTask{
				RunID:       runId,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			err = fmt.Errorf("保存失败：%s", err)
		}
		return
	case CodeSuccess:
		// 执行成功
		err := c.EditTask(ctx, c.gormConfig.client, task.ID).
			Select("status_desc", "number", "run_id", "updated_ip", "result", "next_run_time").
			Updates(GormModelTask{
				StatusDesc:  "执行成功",
				Number:      task.Number + 1,
				RunID:       runId,
				UpdatedIP:   c.config.systemOutsideIP,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			err = fmt.Errorf("保存失败：%s", err)
		}
	case CodeEnd:
		// 执行成功、提前结束
		err := c.EditTask(ctx, c.gormConfig.client, task.ID).
			Select("status", "status_desc", "number", "updated_ip", "result", "next_run_time").
			Updates(GormModelTask{
				Status:      TASK_SUCCESS,
				StatusDesc:  "结束执行",
				Number:      task.Number + 1,
				UpdatedIP:   c.config.systemOutsideIP,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().Time,
			}).Error
		if err != nil {
			err = fmt.Errorf("保存失败：%s", err)
		}
	case CodeError:
		// 执行失败
		err := c.EditTask(ctx, c.gormConfig.client, task.ID).
			Select("status_desc", "number", "run_id", "updated_ip", "result", "next_run_time").
			Updates(GormModelTask{
				StatusDesc:  "执行失败",
				Number:      task.Number + 1,
				RunID:       runId,
				UpdatedIP:   c.config.systemOutsideIP,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			err = fmt.Errorf("保存失败：%s", err)
		}
	}

	if task.MaxNumber != 0 {
		if task.Number+1 >= task.MaxNumber {
			// 关闭执行
			err := c.EditTask(ctx, c.gormConfig.client, task.ID).
				Select("status").
				Updates(GormModelTask{
					Status: TASK_TIMEOUT,
				}).Error
			if err != nil {
				err = fmt.Errorf("保存失败：%s", err)
			}
		}
	}
	return
}
