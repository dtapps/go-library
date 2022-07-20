package jobs

import (
	"github.com/dtapps/go-library/utils/gotime"
	"gorm.io/gorm"
	"log"
)

// Check 任务检查
func (app *App) Check(tx *gorm.DB, vs []Task) {
	if app.MainService > 0 && len(vs) > 0 {
		for _, v := range vs {
			diffInSecondWithAbs := gotime.Current().DiffInSecondWithAbs(gotime.SetCurrentParse(v.UpdatedAt).Time)
			if diffInSecondWithAbs >= v.Frequency*3 {
				log.Printf("每隔%v秒任务：%v相差%v秒\n", v.Frequency, v.Id, diffInSecondWithAbs)
				tx.Where("task_id = ?", v.Id).Where("run_id = ?", v.RunId).Delete(&TaskLogRun{}) // 删除
			}
		}
	}
}
