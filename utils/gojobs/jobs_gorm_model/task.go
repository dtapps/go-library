package jobs_gorm_model

import (
	"gorm.io/gorm"
	"time"
)

// Task 任务
type Task struct {
	Id             uint           `gorm:"primaryKey;comment:记录编号" json:"id"`                      // 记录编号
	Status         string         `gorm:"index;comment:状态码" json:"status"`                        // 状态码
	Params         string         `gorm:"comment:参数" json:"params"`                               // 参数
	ParamsType     string         `gorm:"comment:参数类型" json:"params_type"`                        // 参数类型
	StatusDesc     string         `gorm:"comment:状态描述" json:"status_desc"`                        // 状态描述
	Frequency      int64          `gorm:"index;comment:频率(秒单位)" json:"frequency"`                 // 频率(秒单位)
	Number         int64          `gorm:"comment:当前次数" json:"number"`                             // 当前次数
	MaxNumber      int64          `gorm:"comment:最大次数" json:"max_number"`                         // 最大次数
	RunId          string         `gorm:"index;comment:执行编号" json:"run_id"`                       // 执行编号
	CustomId       string         `gorm:"index;comment:自定义编号" json:"custom_id"`                   // 自定义编号
	CustomSequence int64          `gorm:"index;comment:自定义顺序" json:"custom_sequence"`             // 自定义顺序
	Type           string         `gorm:"index;comment:类型" json:"type"`                           // 类型
	TypeName       string         `gorm:"comment:类型名称" json:"type_name"`                          // 类型名称
	CreatedIp      string         `gorm:"default:0.0.0.0;comment:创建外网IP" json:"created_ip"`       // 创建外网IP
	SpecifyIp      string         `gorm:"default:0.0.0.0;index;comment:指定外网IP" json:"specify_ip"` // 指定外网IP
	UpdatedIp      string         `gorm:"default:0.0.0.0;comment:更新外网IP" json:"updated_ip"`       // 更新外网IP
	Result         string         `gorm:"comment:结果" json:"result"`                               // 结果
	NextRunTime    time.Time      `gorm:"comment:下次运行时间" json:"next_run_time"`                    // 下次运行时间
	CreatedAt      time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`          // 创建时间
	UpdatedAt      time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`          // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at"`                   // 删除时间
}

func (Task) TableName() string {
	return "task"
}
