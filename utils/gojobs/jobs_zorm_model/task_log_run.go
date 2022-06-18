package jobs_zorm_model

// TaskLogRun 任务执行日志模型
type TaskLogRun struct {
	Id         uint   `zorm:"primaryKey" json:"id"`        // 记录编号
	TaskId     uint   `json:"task_id"`                     // 任务编号
	RunId      string `json:"run_id"`                      // 执行编号
	OutsideIp  string `json:"outside_ip"`                  // 外网ip
	InsideIp   string `json:"inside_ip"`                   // 内网ip
	Os         string `json:"os"`                          // 系统类型
	Arch       string `json:"arch"`                        // 系统架构
	Gomaxprocs int    `json:"gomaxprocs"`                  // CPU核数
	GoVersion  string `json:"go_version"`                  // GO版本
	MacAddrs   string `json:"mac_addrs"`                   // Mac地址
	CreatedAt  string `zorm:"type:text" json:"created_at"` // 创建时间
}

func (m *TaskLogRun) TableName() string {
	return "task_log_run"
}
