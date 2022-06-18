package jobs_gorm_model

// TaskLog 任务日志模型
type TaskLog struct {
	Id         uint   `gorm:"primaryKey;comment:记录编号" json:"id"`        // 记录编号
	TaskId     uint   `gorm:"comment:任务编号" json:"task_id"`              // 任务编号
	StatusCode int    `gorm:"comment:状态码" json:"status_code"`           // 状态码
	Desc       string `gorm:"comment:结果" json:"desc"`                   // 结果
	Version    string `gorm:"comment:版本" json:"version"`                // 版本
	CreatedAt  string `gorm:"type:text;comment:创建时间" json:"created_at"` // 创建时间
}

func (m *TaskLog) TableName() string {
	return "task_log"
}
