package gojobs

import (
	"context"
)

const (
	// TASK_IN 任务运行
	TASK_IN = "IN"
	// TASK_CONFLICT 冲突
	TASK_CONFLICT = "CONFLICT"
	// TASK_OBSERVATION 观查
	TASK_OBSERVATION = "OBSERVATION"
	// TASK_SUCCESS 任务完成
	TASK_SUCCESS = "SUCCESS"
	// TASK_ERROR 任务异常
	TASK_ERROR = "ERROR"
	// TASK_TIMEOUT 任务超时
	TASK_TIMEOUT = "TIMEOUT"
	// TASK_WAIT 任务等待
	TASK_WAIT = "WAIT"
)

// Cron
type jobs interface {
	// Run 运行
	Run(ctx context.Context, info GormModelTask, status int, result string)
	// CreateInCustomId 创建正在运行任务
	CreateInCustomId(ctx context.Context, config *ConfigCreateInCustomId) error
	// CreateInCustomIdOnly 创建正在运行唯一任务
	CreateInCustomIdOnly(ctx context.Context, config *ConfigCreateInCustomIdOnly) error
	// CreateInCustomIdMaxNumber 创建正在运行任务并限制数量
	CreateInCustomIdMaxNumber(ctx context.Context, config *ConfigCreateInCustomIdMaxNumber) error
	// CreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
	CreateInCustomIdMaxNumberOnly(ctx context.Context, config *ConfigCreateInCustomIdMaxNumberOnly) error
	// CreateWaitCustomId 创建正在运行任务
	CreateWaitCustomId(ctx context.Context, config *ConfigCreateWaitCustomId) error
}
