package jobs

import (
	"go.dtapp.net/library/utils/gotime"
	"log"
)

// Check 任务检查
func (app *App) Check(vs []Task) {
	if app.MainService > 0 && len(vs) > 0 {
		for _, v := range vs {
			diffInSecondWithAbs := gotime.Current().DiffInSecondWithAbs(gotime.SetCurrentParse(v.UpdatedAt).Time)
			if diffInSecondWithAbs >= v.Frequency*3 {
				log.Printf("每隔%v秒任务：%v相差%v秒\n", v.Frequency, v.Id, diffInSecondWithAbs)
				app.Db.Where("task_id = ?", v.Id).Where("run_id = ?", v.RunId).Delete(&TaskLogRun{}) // 删除
			}
		}
	}
}
