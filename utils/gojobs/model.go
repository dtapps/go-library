package gojobs

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gojson"
	"time"
)

// 创建模型
func (c *Client) autoMigrateTask(ctx context.Context) {
	err := c.gormClient.GetDb().AutoMigrate(&jobs_gorm_model.Task{})
	if err != nil {
		if c.slog.status {
			c.slog.client.WithTraceId(ctx).Error(fmt.Sprintf("创建模型：%s", err))
		}
	}
}

// TaskLog 任务日志模型
type TaskLog struct {
	LogId            uint      `json:"log_id"`             // 日志编号
	TaskId           uint      `json:"task_id"`            // 任务编号
	TaskRunId        string    `json:"task_run_id"`        // 执行编号
	TaskResultStatus string    `json:"task_result_status"` // 执行状态
	TaskResultCode   int       `json:"task_result_code"`   // 执行状态码
	TaskResultDesc   string    `json:"task_result_desc"`   // 执行结果
	SystemHostName   string    `json:"system_host_name"`   // 主机名
	SystemInsideIp   string    `json:"system_inside_ip"`   // 内网ip
	SystemOs         string    `json:"system_os"`          // 系统类型
	SystemArch       string    `json:"system_arch"`        // 系统架构
	GoVersion        string    `json:"go_version"`         // go版本
	SdkVersion       string    `json:"sdk_version"`        // sdk版本
	SystemOutsideIp  string    `json:"system_outside_ip"`  // 外网ip
	LogTime          time.Time `json:"log_time"`           // 日志时间
}

// TaskLogRecord 记录
func (c *Client) TaskLogRecord(ctx context.Context, task jobs_gorm_model.Task, runId string, taskResultCode int, taskResultDesc string) {

	if c.runSlog.status {
		taskLog := TaskLog{
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
		c.runSlog.client.WithTraceId(ctx).Info(gojson.JsonEncodeNoError(taskLog))
	}
}
