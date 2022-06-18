package jobs_xorm_model

// Task 任务
type Task struct {
	Id             uint   `xorm:"pk autoincr" json:"id"`     // 记录编号
	Status         string `json:"status"`                    // 状态码
	Params         string `json:"params"`                    // 参数
	ParamsType     string `json:"params_type"`               // 参数类型
	StatusDesc     string `json:"status_desc"`               // 状态描述
	Frequency      int64  `json:"frequency"`                 // 频率（秒单位）
	Number         int64  `json:"number"`                    // 当前次数
	MaxNumber      int64  `json:"max_number"`                // 最大次数
	RunId          string `json:"run_id"`                    // 执行编号
	CustomId       string `json:"custom_id"`                 // 自定义编号
	CustomSequence int64  `json:"custom_sequence"`           // 自定义顺序
	Type           string `json:"type"`                      // 类型
	CreatedIp      string `json:"created_ip"`                // 创建外网IP
	SpecifyIp      string `json:"specify_ip"`                // 指定外网IP
	UpdatedIp      string `json:"updated_ip"`                // 更新外网IP
	Result         string `json:"result"`                    // 结果
	CreatedAt      string `xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt      string `xorm:"created" json:"updated_at"` // 更新时间
	DeletedAt      string `xorm:"deleted" json:"deleted_at"` // 删除时间
}

func (m *Task) TableName() string {
	return "task"
}
