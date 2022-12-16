package gojobs

import (
	"context"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
)

const (
	TASK_IN      = "IN"      // 任务运行
	TASK_SUCCESS = "SUCCESS" // 任务完成
	TASK_ERROR   = "ERROR"   // 任务异常
	TASK_TIMEOUT = "TIMEOUT" // 任务超时
	TASK_WAIT    = "WAIT"    // 任务等待
)

// Cron
type jobs interface {
	// Run 运行
	Run(ctx context.Context, info jobs_gorm_model.Task, status int, result string)
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
