package jobs_xorm_model

// TaskLog 任务日志模型
type TaskLog struct {
	Id         uint   `xorm:"pk autoincr" json:"id"`     // 记录编号
	TaskId     uint   `json:"task_id"`                   // 任务编号
	StatusCode int    `json:"status_code"`               // 状态码
	Desc       string `json:"desc"`                      // 结果
	Version    int    `json:"version"`                   // 版本
	CreatedAt  string `xorm:"created" json:"created_at"` // 创建时间
}

func (m *TaskLog) TableName() string {
	return "task_log"
}
