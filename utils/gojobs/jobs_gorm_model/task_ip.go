package jobs_gorm_model

// TaskIp 任务Ip
type TaskIp struct {
	Id       int64  `gorm:"primaryKey;comment:记录编号" json:"id"` // 记录编号
	TaskType string `gorm:"comment:任务编号" json:"task_type"`     // 任务编号
	Ips      string `gorm:"comment:任务IP" json:"ips"`           // 任务IP
}

func (m *TaskIp) TableName() string {
	return "task_ip"
}
