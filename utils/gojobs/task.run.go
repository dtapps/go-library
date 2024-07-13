package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"log/slog"
)

// Run 运行
func (c *Client) Run(ctx context.Context, task GormModelTask, taskResultCode int, taskResultDesc string) {

	runId := TraceGetTraceID(ctx)
	if runId == "" {
		runId = gorequest.GetRequestIDContext(ctx)
		if runId == "" {
			slog.InfoContext(ctx, "上下文没有跟踪编号")
			return
		}
	}

	// OpenTelemetry链路追踪
	TraceSetAttributes(ctx, attribute.Int64("task.info.id", int64(task.ID)))
	TraceSetAttributes(ctx, attribute.String("task.info.status", task.Status))
	TraceSetAttributes(ctx, attribute.String("task.info.params", task.Params))
	TraceSetAttributes(ctx, attribute.Int64("task.info.number", task.Number))
	TraceSetAttributes(ctx, attribute.Int64("task.info.max_number", task.MaxNumber))
	TraceSetAttributes(ctx, attribute.String("task.info.custom_id", task.CustomID))
	TraceSetAttributes(ctx, attribute.Int64("task.info.custom_sequence", task.CustomSequence))
	TraceSetAttributes(ctx, attribute.String("task.info.type", task.Type))
	TraceSetAttributes(ctx, attribute.String("task.info.type_name", task.TypeName))
	TraceSetAttributes(ctx, attribute.String("task.run.id", runId))
	TraceSetAttributes(ctx, attribute.Int("task.run.code", taskResultCode))
	TraceSetAttributes(ctx, attribute.String("task.run.desc", taskResultDesc))

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
			TraceRecordError(ctx, err)
			TraceSetStatus(ctx, codes.Error, err.Error())
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
			TraceRecordError(ctx, err)
			TraceSetStatus(ctx, codes.Error, err.Error())
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
			TraceRecordError(ctx, err)
			TraceSetStatus(ctx, codes.Error, err.Error())
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
			TraceRecordError(ctx, err)
			TraceSetStatus(ctx, codes.Error, err.Error())
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
				TraceRecordError(ctx, err)
				TraceSetStatus(ctx, codes.Error, err.Error())
			}
		}
	}
	return
}
