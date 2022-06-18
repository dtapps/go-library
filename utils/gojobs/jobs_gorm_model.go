package gojobs

import (
	"go.dtapp.net/library/utils/gojobs/jobs_gorm_model"
	"go.dtapp.net/library/utils/gotime"
	"gorm.io/gorm"
	"log"
	"strings"
)

// TaskTakeId 查询单任务
func (j *JobsGorm) TaskTakeId(tx *gorm.DB, id uint) (result jobs_gorm_model.Task) {
	tx.Where("id = ?", id).Take(&result)
	return result
}

// TaskTake 查询单任务
func (j *JobsGorm) TaskTake(tx *gorm.DB, customId string) (result jobs_gorm_model.Task) {
	tx.Where("custom_id = ?", customId).Take(&result)
	return result
}

// 查询单任务
func (j *JobsGorm) taskTake(tx *gorm.DB, customId, status string) (result jobs_gorm_model.Task) {
	tx.Where("custom_id = ?", customId).Where("status = ?", status).Take(&result)
	return result
}

// TaskTakeIn 查询单任务 - 任务运行
func (j *JobsGorm) TaskTakeIn(tx *gorm.DB, customId string) jobs_gorm_model.Task {
	return j.taskTake(tx, customId, TASK_IN)
}

// TaskTakeSuccess 查询单任务 - 任务完成
func (j *JobsGorm) TaskTakeSuccess(tx *gorm.DB, customId string) jobs_gorm_model.Task {
	return j.taskTake(tx, customId, TASK_SUCCESS)
}

// TaskTakeError 查询单任务 - 任务异常
func (j *JobsGorm) TaskTakeError(tx *gorm.DB, customId string) jobs_gorm_model.Task {
	return j.taskTake(tx, customId, TASK_ERROR)
}

// TaskTakeTimeout 查询单任务 - 任务超时
func (j *JobsGorm) TaskTakeTimeout(tx *gorm.DB, customId string) jobs_gorm_model.Task {
	return j.taskTake(tx, customId, TASK_TIMEOUT)
}

// TaskTakeWait 查询单任务 - 任务等待
func (j *JobsGorm) TaskTakeWait(tx *gorm.DB, customId string) jobs_gorm_model.Task {
	return j.taskTake(tx, customId, TASK_WAIT)
}

// TaskTypeTake 查询单任务
func (j *JobsGorm) TaskTypeTake(tx *gorm.DB, customId, Type string) (result jobs_gorm_model.Task) {
	tx.Where("custom_id = ?", customId).Where("type = ?", Type).Take(&result)
	return result
}

// 查询单任务
func (j *JobsGorm) taskTypeTake(tx *gorm.DB, customId, Type, status string) (result jobs_gorm_model.Task) {
	tx.Where("custom_id = ?", customId).Where("type = ?", Type).Where("status = ?", status).Take(&result)
	return result
}

// TaskTypeTakeIn 查询单任务 - 任务运行
func (j *JobsGorm) TaskTypeTakeIn(tx *gorm.DB, customId, Type string) jobs_gorm_model.Task {
	return j.taskTypeTake(tx, customId, Type, TASK_IN)
}

// TaskTypeTakeSuccess 查询单任务 - 任务完成
func (j *JobsGorm) TaskTypeTakeSuccess(tx *gorm.DB, customId, Type string) jobs_gorm_model.Task {
	return j.taskTypeTake(tx, customId, Type, TASK_SUCCESS)
}

// TaskTypeTakeError 查询单任务 - 任务异常
func (j *JobsGorm) TaskTypeTakeError(tx *gorm.DB, customId, Type string) jobs_gorm_model.Task {
	return j.taskTypeTake(tx, customId, Type, TASK_ERROR)
}

// TaskTypeTakeTimeout 查询单任务 - 任务超时
func (j *JobsGorm) TaskTypeTakeTimeout(tx *gorm.DB, customId, Type string) jobs_gorm_model.Task {
	return j.taskTypeTake(tx, customId, Type, TASK_TIMEOUT)
}

// TaskTypeTakeWait 查询单任务 - 任务等待
func (j *JobsGorm) TaskTypeTakeWait(tx *gorm.DB, customId, Type string) jobs_gorm_model.Task {
	return j.taskTypeTake(tx, customId, Type, TASK_WAIT)
}

// TaskFindAll 查询多任务
func (j *JobsGorm) TaskFindAll(tx *gorm.DB, frequency int64) (results []jobs_gorm_model.Task) {
	tx.Where("frequency = ?", frequency).Order("id asc").Find(&results)
	return results
}

// 查询多任务
func (j *JobsGorm) taskFindAll(tx *gorm.DB, frequency int64, status string) (results []jobs_gorm_model.Task) {
	tx.Where("frequency = ?", frequency).Where("status = ?", status).Order("id asc").Find(&results)
	return results
}

// TaskFindAllIn 查询多任务 - 任务运行
func (j *JobsGorm) TaskFindAllIn(tx *gorm.DB, frequency int64) []jobs_gorm_model.Task {
	return j.taskFindAll(tx, frequency, TASK_IN)
}

// TaskFindAllSuccess 查询多任务 - 任务完成
func (j *JobsGorm) TaskFindAllSuccess(tx *gorm.DB, frequency int64) []jobs_gorm_model.Task {
	return j.taskFindAll(tx, frequency, TASK_SUCCESS)
}

// TaskFindAllError 查询多任务 - 任务异常
func (j *JobsGorm) TaskFindAllError(tx *gorm.DB, frequency int64) []jobs_gorm_model.Task {
	return j.taskFindAll(tx, frequency, TASK_ERROR)
}

// TaskFindAllTimeout 查询多任务 - 任务超时
func (j *JobsGorm) TaskFindAllTimeout(tx *gorm.DB, frequency int64) []jobs_gorm_model.Task {
	return j.taskFindAll(tx, frequency, TASK_TIMEOUT)
}

// TaskFindAllWait 查询多任务 - 任务等待
func (j *JobsGorm) TaskFindAllWait(tx *gorm.DB, frequency int64) []jobs_gorm_model.Task {
	return j.taskFindAll(tx, frequency, TASK_WAIT)
}

// EditTask 任务修改
func (j *JobsGorm) EditTask(tx *gorm.DB, id uint) *gorm.DB {
	return tx.Model(&jobs_gorm_model.Task{}).Where("id = ?", id)
}

// UpdateFrequency 更新任务频率
func (j *JobsGorm) UpdateFrequency(tx *gorm.DB, id uint, frequency int64) *gorm.DB {
	return j.EditTask(tx, id).
		Select("frequency", "updated_at").
		Updates(jobs_gorm_model.Task{
			Frequency: frequency,
			UpdatedAt: gotime.Current().Format(),
		})
}

func (j *JobsGorm) taskIpTake(tx *gorm.DB, taskType, ips string) (result jobs_gorm_model.TaskIp) {
	tx.Where("task_type = ?", taskType).Where("ips = ?", ips).Take(&result)
	return result
}

// TaskIpUpdate 更新ip
func (j *JobsGorm) TaskIpUpdate(tx *gorm.DB, taskType, ips string) *gorm.DB {
	query := j.taskIpTake(tx, taskType, ips)
	if query.Id != 0 {
		return tx
	}
	updateStatus := tx.Create(&jobs_gorm_model.TaskIp{
		TaskType: taskType,
		Ips:      ips,
	})
	if updateStatus.RowsAffected == 0 {
		log.Println("任务更新失败：", updateStatus.Error)
	}
	return updateStatus
}

// TaskIpInit 实例任务ip
func (j *JobsGorm) TaskIpInit(tx *gorm.DB, ips map[string]string) bool {
	if j.outsideIp == "" || j.outsideIp == "0.0.0.0" {
		return false
	}
	tx.Where("ips = ?", j.outsideIp).Delete(&jobs_gorm_model.TaskIp{}) // 删除
	for k, v := range ips {
		if v == "" {
			j.TaskIpUpdate(tx, k, j.outsideIp)
		} else {
			find := strings.Contains(v, ",")
			if find == true {
				// 包含
				parts := strings.Split(v, ",")
				for _, vv := range parts {
					if vv == j.outsideIp {
						j.TaskIpUpdate(tx, k, j.outsideIp)
					}
				}
			} else {
				// 不包含
				if v == j.outsideIp {
					j.TaskIpUpdate(tx, k, j.outsideIp)
				}
			}
		}
	}
	return true
}

// TaskLogRunTake 查询任务执行日志
func (j *JobsGorm) TaskLogRunTake(tx *gorm.DB, taskId uint, runId string) (result jobs_gorm_model.TaskLogRun) {
	tx.Select("id", "os", "arch", "outside_ip", "created_at").Where("task_id = ?", taskId).Where("run_id = ?", runId).Take(&result)
	return result
}
