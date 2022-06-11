package jobs

import (
	"gorm.io/gorm"
	"log"
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
	Id             uint   `gorm:"primaryKey" json:"id"`        // 记录编号
	Status         string `json:"status"`                      // 状态码
	Params         string `json:"params"`                      // 参数
	ParamsType     string `json:"params_type"`                 // 参数类型
	StatusDesc     string `json:"status_desc"`                 // 状态描述
	Frequency      int64  `json:"frequency"`                   // 频率（秒单位）
	Number         int64  `json:"number"`                      // 当前次数
	MaxNumber      int64  `json:"max_number"`                  // 最大次数
	RunId          string `json:"run_id"`                      // 执行编号
	CustomId       string `json:"custom_id"`                   // 自定义编号
	CustomSequence int64  `json:"custom_sequence"`             // 自定义顺序
	Type           string `json:"type"`                        // 类型
	CreatedIp      string `json:"created_ip"`                  // 创建外网IP
	SpecifyIp      string `json:"specify_ip"`                  // 指定外网IP
	UpdatedIp      string `json:"updated_ip"`                  // 更新外网IP
	Result         string `json:"result"`                      // 结果
	CreatedAt      string `gorm:"type:text" json:"created_at"` // 创建时间
	UpdatedAt      string `gorm:"type:text" json:"updated_at"` // 更新时间
}

func (m *Task) TableName() string {
	return "task"
}

// TaskTake 查询任务
func (app *App) TaskTake(tx *gorm.DB, customId string) (result Task) {
	tx.Where("custom_id = ?", customId).Where("status = ?", TASK_IN).Take(&result)
	return result
}

// TaskCustomIdTake 查询任务
func (app *App) TaskCustomIdTake(tx *gorm.DB, Type, customId string) (result Task) {
	tx.Where("type = ?", Type).Where("custom_id = ?", customId).Take(&result)
	return result
}

// TaskCustomIdTakeStatus 查询任务
func (app *App) TaskCustomIdTakeStatus(tx *gorm.DB, Type, customId, status string) (result Task) {
	tx.Where("type = ?", Type).Where("custom_id = ?", customId).Where("status = ?", status).Take(&result)
	return result
}

// TaskFind 查询任务
func (app *App) TaskFind(tx *gorm.DB, frequency int64) (results []Task) {
	tx.Table("task").Select("task.*").Where("task.frequency = ?", frequency).Where("task.status = ?", TASK_IN).Where("task_ip.ips = ?", app.OutsideIp).Order("task.id asc").Joins("left join task_ip on task_ip.task_type = task.type").Find(&results)
	return app.taskFindCheck(results)
}

// TaskFindAll 查询任务
func (app *App) TaskFindAll(tx *gorm.DB, frequency int64) (results []Task) {
	tx.Where("frequency = ?", frequency).Where("status = ?", TASK_IN).Order("id asc").Find(&results)
	return results
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
	Id         uint   `gorm:"primaryKey" json:"id"`        // 记录编号
	TaskId     uint   `json:"task_id"`                     // 任务编号
	StatusCode int    `json:"status_code"`                 // 状态码
	Desc       string `json:"desc"`                        // 结果
	Version    int    `json:"version"`                     // 版本
	CreatedAt  string `gorm:"type:text" json:"created_at"` // 创建时间
}

func (m *TaskLog) TableName() string {
	return "task_log"
}

// TaskLogRun 任务执行日志
type TaskLogRun struct {
	Id         uint   `gorm:"primaryKey" json:"id"`        // 记录编号
	TaskId     uint   `json:"task_id"`                     // 任务编号
	RunId      string `json:"run_id"`                      // 执行编号
	OutsideIp  string `json:"outside_ip"`                  // 外网ip
	InsideIp   string `json:"inside_ip"`                   // 内网ip
	Os         string `json:"os"`                          // 系统类型
	Arch       string `json:"arch"`                        // 系统架构
	Gomaxprocs int    `json:"gomaxprocs"`                  // CPU核数
	GoVersion  string `json:"go_version"`                  // GO版本
	MacAddrs   string `json:"mac_addrs"`                   // Mac地址
	CreatedAt  string `gorm:"type:text" json:"created_at"` // 创建时间
}

func (m *TaskLogRun) TableName() string {
	return "task_log_run"
}

// TaskLogRunTake 查询任务执行日志
func (app *App) TaskLogRunTake(tx *gorm.DB, taskId uint, runId string) (result TaskLogRun) {
	tx.Select("id", "os", "arch", "outside_ip", "created_at").Where("task_id = ?", taskId).Where("run_id = ?", runId).Take(&result)
	return result
}

// TaskIp 任务Ip
type TaskIp struct {
	Id       int64
	TaskType string `json:"task_type"` // 任务编号
	Ips      string `json:"ips"`       // 任务IP
}

func (m *TaskIp) TableName() string {
	return "task_ip"
}

func (app *App) TaskIpUpdate(tx *gorm.DB, taskType, ips string) *gorm.DB {
	var query TaskIp
	tx.Where("task_type = ?", taskType).Where("ips = ?", ips).Take(&query)
	if query.Id != 0 {
		return tx
	}
	updateStatus := tx.Create(&TaskIp{
		TaskType: taskType,
		Ips:      ips,
	})
	if updateStatus.RowsAffected == 0 {
		log.Println("任务更新失败：", updateStatus.Error)
	}
	return updateStatus
}

// TaskIpInit 实例任务ip
func (app *App) TaskIpInit(tx *gorm.DB, ips map[string]string) bool {
	if app.OutsideIp == "" || app.OutsideIp == "0.0.0.0" {
		return false
	}
	tx.Where("ips = ?", app.OutsideIp).Delete(&TaskIp{}) // 删除
	for k, v := range ips {
		if v == "" {
			app.TaskIpUpdate(tx, k, app.OutsideIp)
		} else {
			find := strings.Contains(v, ",")
			if find == true {
				// 包含
				parts := strings.Split(v, ",")
				for _, vv := range parts {
					if vv == app.OutsideIp {
						app.TaskIpUpdate(tx, k, app.OutsideIp)
					}
				}
			} else {
				// 不包含
				if v == app.OutsideIp {
					app.TaskIpUpdate(tx, k, app.OutsideIp)
				}
			}
		}
	}
	return true
}
