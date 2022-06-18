package gojobs

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
	Run(info interface{}, status int, desc string)
	// RunAddLog 任务执行日志
	RunAddLog(id uint, runId string)
	// CreateInCustomId 创建正在运行任务
	CreateInCustomId()
	// CreateInCustomIdOnly 创建正在运行唯一任务
	CreateInCustomIdOnly()
	// CreateInCustomIdMaxNumber 创建正在运行任务并限制数量
	CreateInCustomIdMaxNumber()
	// CreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
	CreateInCustomIdMaxNumberOnly()
}
