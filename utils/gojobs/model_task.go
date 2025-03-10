package gojobs

import (
	"time"
)

// ModelTask 任务
type ModelTask struct {
	ID             int64     `json:"id,omitempty"`              // 记录编号
	Status         string    `json:"status,omitempty"`          // 状态码
	StatusDesc     string    `json:"status_desc,omitempty"`     // 状态描述
	Params         []byte    `json:"params,omitempty"`          // 参数
	Frequency      int64     `json:"frequency,omitempty"`       // 频率(秒单位)
	Spec           string    `json:"spec,omitempty"`            // cron表达式
	Number         int64     `json:"number,omitempty"`          // 当前次数
	MaxNumber      int64     `json:"max_number,omitempty"`      // 最大次数
	RunID          string    `json:"run_id,omitempty"`          // 执行编号
	CustomID       string    `json:"custom_id,omitempty"`       // 自定义编号
	CustomSequence int64     `json:"custom_sequence,omitempty"` // 自定义顺序
	Type           string    `json:"type,omitempty"`            // 类型
	TypeName       string    `json:"type_name,omitempty"`       // 类型名称
	CreatedIP      string    `json:"created_ip,omitempty"`      // 创建外网IP
	SpecifyIP      string    `json:"specify_ip,omitempty"`      // 指定外网IP
	UpdatedIP      string    `json:"updated_ip,omitempty"`      // 更新外网IP
	Result         string    `json:"result,omitempty"`          // 结果
	NextRunTime    time.Time `json:"next_run_time,omitempty"`   // 下次运行时间
	CreatedAt      time.Time `json:"created_at,omitempty"`      // 创建时间
	UpdatedAt      time.Time `json:"updated_at,omitempty"`      // 更新时间
}
