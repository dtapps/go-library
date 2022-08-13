package gojobs

import (
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
	"log"
)

// CheckManyTask 多任务检查
func (j *JobsGorm) CheckManyTask(tx *gorm.DB, vs []jobs_gorm_model.Task) {
	if len(vs) > 0 {
		for _, v := range vs {
			diffInSecondWithAbs := gotime.Current().DiffInSecondWithAbs(gotime.SetCurrent(v.UpdatedAt).Time)
			if diffInSecondWithAbs >= v.Frequency*3 {
				if j.config.debug == true {
					log.Printf("每隔%v秒任务：%v相差%v秒\n", v.Frequency, v.Id, diffInSecondWithAbs)
				}
				err := tx.Where("task_id = ?", v.Id).Where("run_id = ?", v.RunId).Delete(&jobs_gorm_model.TaskLogRun{}).Error
				if err != nil {
					log.Println("删除失败", err.Error())
				}
			}
		}
	}
}

// CheckSingleTask 单任务检查
func (j *JobsGorm) CheckSingleTask(tx *gorm.DB, v jobs_gorm_model.Task) {
	diffInSecondWithAbs := gotime.Current().DiffInSecondWithAbs(gotime.SetCurrent(v.UpdatedAt).Time)
	if diffInSecondWithAbs >= v.Frequency*3 {
		if j.config.debug == true {
			log.Printf("每隔%v秒任务：%v相差%v秒\n", v.Frequency, v.Id, diffInSecondWithAbs)
		}
		err := tx.Where("task_id = ?", v.Id).Where("run_id = ?", v.RunId).Delete(&jobs_gorm_model.TaskLogRun{}).Error
		if err != nil {
			log.Println("删除失败", err.Error())
		}
	}
}
