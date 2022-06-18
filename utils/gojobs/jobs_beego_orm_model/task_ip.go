package jobs_beego_orm_model

// TaskIp 任务Ip
type TaskIp struct {
	Id       int64  `orm:"auto" json:"id"`
	TaskType string `json:"task_type"` // 任务编号
	Ips      string `json:"ips"`       // 任务IP
}

func (m *TaskIp) TableName() string {
	return "task_ip"
}
