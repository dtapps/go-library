package jobs_gorm_model

import "time"

// TaskLog 任务日志模型
type TaskLog struct {
	LogId           uint      `gorm:"primaryKey;comment:日志编号" json:"log_id"`                 // 日志编号
	TaskId          uint      `gorm:"index;comment:任务编号" json:"task_id"`                     // 任务编号
	TaskRunId       string    `gorm:"comment:执行编号" json:"task_run_id"`                       // 执行编号
	TaskResultCode  int       `gorm:"index;comment:执行状态码" json:"task_result_code"`           // 执行状态码
	TaskResultDesc  string    `gorm:"comment:执行结果" json:"task_result_desc"`                  // 执行结果
	SystemHostName  string    `gorm:"comment:主机名" json:"system_host_name"`                   // 主机名
	SystemInsideIp  string    `gorm:"default:0.0.0.0;comment:内网ip" json:"system_inside_ip"`  // 内网ip
	SystemOs        string    `gorm:"comment:系统类型" json:"system_os"`                         // 系统类型
	SystemArch      string    `gorm:"comment:系统架构" json:"system_arch"`                       // 系统架构
	GoVersion       string    `gorm:"comment:go版本" json:"go_version"`                        // go版本
	SdkVersion      string    `gorm:"comment:sdk版本" json:"sdk_version"`                      // sdk版本
	SystemOutsideIp string    `gorm:"default:0.0.0.0;comment:外网ip" json:"system_outside_ip"` // 外网ip
	LogTime         time.Time `gorm:"autoCreateTime;comment:日志时间" json:"log_time"`           // 日志时间
}

func (TaskLog) TableName() string {
	return "task_log"
}
