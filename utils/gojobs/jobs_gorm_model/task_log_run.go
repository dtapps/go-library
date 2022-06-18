package jobs_gorm_model

// TaskLogRun 任务执行日志模型
type TaskLogRun struct {
	Id         uint   `gorm:"primaryKey;comment:记录编号" json:"id"`        // 记录编号
	TaskId     uint   `gorm:"comment:任务编号" json:"task_id"`              // 任务编号
	RunId      string `gorm:"comment:执行编号" json:"run_id"`               // 执行编号
	OutsideIp  string `gorm:"comment:外网ip" json:"outside_ip"`           // 外网ip
	InsideIp   string `gorm:"comment:内网ip" json:"inside_ip"`            // 内网ip
	Os         string `gorm:"comment:系统类型" json:"os"`                   // 系统类型
	Arch       string `gorm:"comment:系统架构" json:"arch"`                 // 系统架构
	Gomaxprocs int    `gorm:"comment:CPU核数" json:"gomaxprocs"`          // CPU核数
	GoVersion  string `gorm:"comment:GO版本" json:"go_version"`           // GO版本
	MacAddrs   string `gorm:"comment:Mac地址" json:"mac_addrs"`           // Mac地址
	CreatedAt  string `gorm:"type:text;comment:创建时间" json:"created_at"` // 创建时间
}

func (m *TaskLogRun) TableName() string {
	return "task_log_run"
}
