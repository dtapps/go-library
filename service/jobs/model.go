package jobs

import (
	"strings"
)

const (
	TASK_IN      = "IN"      // 任务运行
	TASK_SUCCESS = "SUCCESS" // 任务完成
	TASK_ERROR   = "ERROR"   // 任务异常
	TASK_TIMEOUT = "TIMEOUT" // 任务超时
	TASK_WAIT    = "WAIT"    // 任务等待
)

// Task 任务
type Task struct {
	Id             int64
	Status         string `gorm:"type:text" json:"status"`            // 状态码
	Params         string `gorm:"type:text" json:"params"`            // 参数
	ParamsType     string `gorm:"type:text" json:"params_type"`       // 参数类型
	StatusDesc     string `gorm:"type:text" json:"status_desc"`       // 状态描述
	Frequency      int64  `gorm:"type:bigint" json:"frequency"`       // 频率（秒单位）
	Number         int64  `gorm:"type:bigint" json:"number"`          // 当前次数
	MaxNumber      int64  `gorm:"type:bigint" json:"max_number"`      // 最大次数
	RunId          string `gorm:"type:text" json:"run_id"`            // 执行编号
	CustomId       string `gorm:"type:text" json:"custom_id"`         // 自定义编号
	CustomSequence int64  `gorm:"type:bigint" json:"custom_sequence"` // 自定义顺序
	Type           string `gorm:"type:text" json:"type"`              // 类型
	CreatedIp      string `gorm:"type:text" json:"created_ip"`        // 创建外网IP
	SpecifyIp      string `gorm:"type:text" json:"specify_ip"`        // 指定外网IP
	UpdatedIp      string `gorm:"type:text" json:"updated_ip"`        // 更新外网IP
	Result         string `gorm:"type:text" json:"result"`            // 结果
	CreatedAt      string `gorm:"type:text" json:"created_at"`        // 创建时间
	UpdatedAt      string `gorm:"type:text" json:"updated_at"`        // 更新时间
}

func (m *Task) TableName() string {
	return "task"
}

// TaskTake 查询任务
func (app *App) TaskTake(customId string) (result Task) {
	app.Db.Where("custom_id = ?", customId).Where("status = ?", TASK_IN).Take(&result)
	return result
}

// TaskCustomIdTake 查询任务
func (app *App) TaskCustomIdTake(Type, customId string) (result Task) {
	app.Db.Where("type = ?", Type).Where("custom_id = ?", customId).Take(&result)
	return result
}

// TaskCustomIdTakeStatus 查询任务
func (app *App) TaskCustomIdTakeStatus(Type, customId, status string) (result Task) {
	app.Db.Where("type = ?", Type).Where("custom_id = ?", customId).Where("status = ?", status).Take(&result)
	return result
}

// TaskFind 查询任务
func (app *App) TaskFind(frequency int) (results []Task) {
	app.Db.Table("task").Select("task.*").Where("task.frequency = ?", frequency).Where("task.status = ?", TASK_IN).Where("task_ip.ips = ?", app.OutsideIp).Order("task.id asc").Joins("left join task_ip on task_ip.task_type = task.type").Find(&results)
	return app.taskFindCheck(results)
}

// 检查任务
func (app *App) taskFindCheck(lists []Task) (results []Task) {
	for _, v := range lists {
		if v.SpecifyIp == "" {
			results = append(results, v)
		} else {
			if app.OutsideIp == v.SpecifyIp {
				results = append(results, v)
			}
		}
	}
	return results
}

// TaskLog 任务日志
type TaskLog struct {
	Id         int64
	TaskId     int64  `gorm:"type:bigint" json:"task_id"`     // 任务编号
	StatusCode int    `gorm:"type:bigint" json:"status_code"` // 状态码
	Desc       string `gorm:"type:text" json:"desc"`          // 结果
	Version    int    `gorm:"type:bigint" json:"version"`     // 版本
	CreatedAt  string `gorm:"type:text" json:"created_at"`    // 创建时间
}

func (m *TaskLog) TableName() string {
	return "task_log"
}

// TaskLogRun 任务执行日志
type TaskLogRun struct {
	Id         int64
	TaskId     int64  `gorm:"type:bigint" json:"task_id"`    // 任务编号
	RunId      string `gorm:"type:text" json:"run_id"`       // 执行编号
	OutsideIp  string `gorm:"type:text" json:"outside_ip"`   // 外网ip
	InsideIp   string `gorm:"type:text" json:"inside_ip"`    // 内网ip
	Os         string `gorm:"type:text" json:"os"`           // 系统类型
	Arch       string `gorm:"type:text" json:"arch"`         // 系统架构
	Gomaxprocs int    `gorm:"type:bigint" json:"gomaxprocs"` // CPU核数
	GoVersion  string `gorm:"type:text" json:"go_version"`   // GO版本
	MacAddrs   string `gorm:"type:text" json:"mac_addrs"`    // Mac地址
	CreatedAt  string `gorm:"type:text" json:"created_at"`   // 创建时间
}

func (m *TaskLogRun) TableName() string {
	return "task_log_run"
}

// TaskLogRunTake 查询任务执行日志
func (app *App) TaskLogRunTake(taskId int64, runId string) (result TaskLogRun) {
	app.Db.Select("id", "os", "arch", "outside_ip", "created_at").Where("task_id = ?", taskId).Where("run_id = ?", runId).Take(&result)
	return result
}

// TaskIp 任务Ip
type TaskIp struct {
	Id       int64
	TaskType string `gorm:"type:text" json:"task_type"` // 任务编号
	Ips      string `gorm:"type:text" json:"ips"`       // 任务IP
}

func (m *TaskIp) TableName() string {
	return "task_ip"
}

func (app *App) TaskIpUpdate(taskType, ips string) int64 {
	var query TaskIp
	app.Db.Where("task_type = ?", taskType).Where("ips = ?", ips).Take(&query)
	if query.Id != 0 {
		return query.Id
	}
	return app.Db.Create(&TaskIp{
		TaskType: taskType,
		Ips:      ips,
	}).RowsAffected
}

// TaskIpInit 实例任务ip
func (app *App) TaskIpInit(ips map[string]string) bool {
	if app.OutsideIp == "" || app.OutsideIp == "0.0.0.0" {
		return false
	}
	app.Db.Where("ips = ?", app.OutsideIp).Delete(&TaskIp{}) // 删除
	for k, v := range ips {
		if v == "" {
			app.TaskIpUpdate(k, app.OutsideIp)
		} else {
			find := strings.Contains(v, ",")
			if find == true {
				// 包含
				parts := strings.Split(v, ",")
				for _, vv := range parts {
					if vv == app.OutsideIp {
						app.TaskIpUpdate(k, app.OutsideIp)
					}
				}
			} else {
				// 不包含
				if v == app.OutsideIp {
					app.TaskIpUpdate(k, app.OutsideIp)
				}
			}
		}
	}
	return true
}