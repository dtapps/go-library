package jobs

import (
	"go.dtapp.net/library/gojson"
	"go.dtapp.net/library/goredis"
	"go.dtapp.net/library/gotime"
	"go.dtapp.net/library/gouuid"
	"gorm.io/gorm"
	"log"
)

type App struct {
	RunVersion  int         `json:"run_version"`  // 运行版本
	Os          string      `json:"os"`           // 系统类型
	Arch        string      `json:"arch"`         // 系统架构
	MaxProCs    int         `json:"max_pro_cs"`   // CPU核数
	Version     string      `json:"version"`      // GO版本
	MacAddrS    string      `json:"mac_addr_s"`   // Mac地址
	InsideIp    string      `json:"inside_ip"`    // 内网ip
	OutsideIp   string      `json:"outside_ip"`   // 外网ip
	MainService int         `json:"main_service"` // 主要服务
	Db          *gorm.DB    // 数据库
	Redis       goredis.App // 缓存数据库服务
}

// Add 添加任务
func (app *App) Add(tx *gorm.DB, Type string, params interface{}, frequency int64) *gorm.DB {
	return tx.Create(&Task{
		Status:     TASK_IN,
		Params:     gojson.JsonEncodeNoError(params),
		StatusDesc: "首次添加任务",
		Frequency:  frequency,
		RunId:      gouuid.GetUuId(),
		Type:       Type,
		CreatedIp:  app.OutsideIp,
		UpdatedIp:  app.OutsideIp,
		CreatedAt:  gotime.Current().Format(),
		UpdatedAt:  gotime.Current().Format(),
	})
}

// AddCustomId 添加任务
func (app *App) AddCustomId(tx *gorm.DB, Type string, params interface{}, frequency int64, customId string) *gorm.DB {
	query := app.TaskCustomIdTake(tx, Type, customId)
	if query.Id != 0 {
		return tx
	}
	return tx.Create(&Task{
		Status:     TASK_IN,
		Params:     gojson.JsonEncodeNoError(params),
		StatusDesc: "首次添加任务",
		Frequency:  frequency,
		RunId:      gouuid.GetUuId(),
		CustomId:   customId,
		Type:       Type,
		CreatedIp:  app.OutsideIp,
		UpdatedIp:  app.OutsideIp,
		CreatedAt:  gotime.Current().Format(),
		UpdatedAt:  gotime.Current().Format(),
	})
}

// AddCustomIdMaxNumber 添加任务并设置最大数量
func (app *App) AddCustomIdMaxNumber(tx *gorm.DB, Type string, params interface{}, frequency int64, customId string, maxNumber int64) *gorm.DB {
	query := app.TaskCustomIdTakeStatus(tx, Type, customId, TASK_IN)
	if query.Id != 0 {
		return tx
	}
	return tx.Create(&Task{
		Status:     TASK_IN,
		Params:     gojson.JsonEncodeNoError(params),
		StatusDesc: "首次添加任务",
		Frequency:  frequency,
		MaxNumber:  maxNumber,
		RunId:      gouuid.GetUuId(),
		CustomId:   customId,
		Type:       Type,
		CreatedIp:  app.OutsideIp,
		UpdatedIp:  app.OutsideIp,
		CreatedAt:  gotime.Current().Format(),
		UpdatedAt:  gotime.Current().Format(),
	})
}

type TaskParams = Task

// AddInOrder 添加订单可执行任务
func (app *App) AddInOrder(tx *gorm.DB, Type string, params interface{}, frequency int64) *gorm.DB {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsOrderType
	return app.AddIn(tx, param, params)
}

// AddInOrderCustomId 添加订单可执行任务
func (app *App) AddInOrderCustomId(tx *gorm.DB, Type string, params interface{}, frequency int64, customId string) *gorm.DB {
	query := app.TaskCustomIdTakeStatus(tx, Type, customId, TASK_IN)
	if query.Id != 0 {
		return tx
	}
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.CustomId = customId
	param.ParamsType = ParamsOrderType
	return app.AddIn(tx, param, params)
}

// AddInOrderCustomIdSpecifyIp 添加订单可执行任务
func (app *App) AddInOrderCustomIdSpecifyIp(tx *gorm.DB, Type string, params interface{}, frequency int64, customId, specifyIp string) *gorm.DB {
	query := app.TaskCustomIdTakeStatus(tx, Type, customId, TASK_IN)
	if query.Id != 0 {
		return tx
	}
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.CustomId = customId
	param.SpecifyIp = specifyIp
	param.ParamsType = ParamsOrderType
	return app.AddIn(tx, param, params)
}

// AddInMerchantGoldenBean 添加商家金豆可执行任务
func (app *App) AddInMerchantGoldenBean(tx *gorm.DB, Type string, params interface{}, frequency int64) *gorm.DB {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsMerchantGoldenBeanType
	return app.AddIn(tx, param, params)
}

// AddInTeamInv 添加团队邀请可执行任务
func (app *App) AddInTeamInv(tx *gorm.DB, Type string, params interface{}, frequency int64) *gorm.DB {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsTeamInvType
	return app.AddIn(tx, param, params)
}

// AddInUserShareInvitation 邀请好友
func (app *App) AddInUserShareInvitation(tx *gorm.DB, Type string, params interface{}, frequency int64) *gorm.DB {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	return app.AddIn(tx, param, params)
}

// AddInNewService 添加企业自定义可执行任务
func (app *App) AddInNewService(tx *gorm.DB, Type string, params interface{}, frequency int64) *gorm.DB {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsNewServiceType
	return app.AddIn(tx, param, params)
}

// AddInOrderCustomIdObservation 添加观察接口任务
func (app *App) AddInOrderCustomIdObservation(tx *gorm.DB, Type string, customId string) *gorm.DB {
	query := app.TaskCustomIdTakeStatus(tx, Type, customId, TASK_IN)
	if query.Id != 0 {
		return tx
	}
	var param TaskParams
	param.Type = Type
	param.Frequency = 3600
	param.MaxNumber = 24 * 5 // 一个星期
	param.CustomId = customId
	param.ParamsType = ParamsOrderType
	return app.AddIn(tx, param, ParamsOrderId{
		OrderId: customId,
	})
}

// AddInOrderCustomIdObservationClone 观察接口关闭
func (app *App) AddInOrderCustomIdObservationClone(tx *gorm.DB, Type string, customId string) *gorm.DB {
	query := app.TaskCustomIdTakeStatus(tx, Type, customId, TASK_IN)
	if query.Id == 0 {
		return tx
	}
	return app.Edit(tx, query.Id).Select("status", "status_desc", "run_id", "updated_ip", "updated_at").
		Updates(Task{
			Status:     TASK_SUCCESS,
			StatusDesc: "已完成，停止观察",
			RunId:      gouuid.GetUuId(),
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
		})
}

// AddIn 添加可执行任务
// params.Type 任务类型
// params.Frequency 任务频率
// params.CustomId 自定义编号
// params 任务参数
func (app *App) AddIn(tx *gorm.DB, param TaskParams, params interface{}) *gorm.DB {
	param.Status = TASK_IN
	param.StatusDesc = "首次添加任务"
	param.RunId = gouuid.GetUuId()
	param.Params = gojson.JsonEncodeNoError(params)
	param.CreatedIp = app.OutsideIp
	param.UpdatedIp = app.OutsideIp
	param.CreatedAt = gotime.Current().Format()
	param.UpdatedAt = gotime.Current().Format()
	status := tx.Create(&param)
	if status.RowsAffected == 0 {
		log.Println("AddIn：", status.Error)
	}
	return status
}

// AddWaitNewServiceNext 添加企业自定义下一步等待执行任务
func (app *App) AddWaitNewServiceNext(tx *gorm.DB, param TaskParams, params interface{}) *gorm.DB {
	param.ParamsType = ParamsNewServiceNextType
	return app.AddWait(tx, param, params)
}

// AddWait 添加等待执行任务
// params.Type 任务类型
// params.Frequency 任务频率
// params.CustomId 自定义编号
// params.CustomSequence 自定义顺序
// params 任务参数
func (app *App) AddWait(tx *gorm.DB, param TaskParams, params interface{}) *gorm.DB {
	param.Status = TASK_WAIT
	param.StatusDesc = "首次添加任务"
	param.RunId = gouuid.GetUuId()
	param.Params = gojson.JsonEncodeNoError(params)
	param.CreatedIp = app.OutsideIp
	param.UpdatedIp = app.OutsideIp
	param.CreatedAt = gotime.Current().Format()
	param.UpdatedAt = gotime.Current().Format()
	return tx.Create(&param)
}

// Edit 任务修改
func (app *App) Edit(tx *gorm.DB, id uint) *gorm.DB {
	return tx.Model(&Task{}).Where("id = ?", id)
}

// UpdateFrequency 更新任务频率
func (app *App) UpdateFrequency(tx *gorm.DB, id uint, frequency int64) *gorm.DB {
	return app.Edit(tx, id).
		Updates(map[string]interface{}{
			"frequency": frequency,
		})
}

// Start 任务启动
func (app *App) Start(tx *gorm.DB, customId string, customSequence int64) *gorm.DB {
	return tx.Model(&Task{}).
		Where("custom_id = ?", customId).
		Where("custom_sequence = ?", customSequence).
		Where("status = ?", TASK_WAIT).
		Select("status", "status_desc", "updated_ip", "updated_at").
		Updates(Task{
			Status:     TASK_IN,
			StatusDesc: "启动任务",
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
		})
}

// RunAddLog 任务执行日志
func (app *App) RunAddLog(tx *gorm.DB, id uint, runId string) *gorm.DB {
	return tx.Create(&TaskLogRun{
		TaskId:     id,
		RunId:      runId,
		InsideIp:   app.InsideIp,
		OutsideIp:  app.OutsideIp,
		Os:         app.Os,
		Arch:       app.Arch,
		Gomaxprocs: app.MaxProCs,
		GoVersion:  app.Version,
		MacAddrs:   app.MacAddrS,
		CreatedAt:  gotime.Current().Format(),
	})
}

// Run 任务执行
func (app *App) Run(tx *gorm.DB, info Task, status int, desc string) {
	// 请求函数记录
	statusCreate := tx.Create(&TaskLog{
		TaskId:     info.Id,
		StatusCode: status,
		Desc:       desc,
		Version:    app.RunVersion,
		CreatedAt:  gotime.Current().Format(),
	})
	if statusCreate.RowsAffected == 0 {
		log.Println("statusCreate", statusCreate.Error)
	}
	if status == 0 {
		statusEdit := app.Edit(tx, info.Id).Select("run_id").Updates(Task{
			RunId: gouuid.GetUuId(),
		})
		if statusEdit.RowsAffected == 0 {
			log.Println("statusEdit", statusEdit.Error)
		}
		return
	}
	// 任务
	if status == CodeSuccess {
		// 执行成功
		statusEdit := app.Edit(tx, info.Id).Select("status_desc", "number", "run_id", "updated_ip", "updated_at", "result").Updates(Task{
			StatusDesc: "执行成功",
			Number:     info.Number + 1,
			RunId:      gouuid.GetUuId(),
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
			Result:     desc,
		})
		if statusEdit.RowsAffected == 0 {
			log.Println("statusEdit", statusEdit.Error)
		}
	}
	if status == CodeEnd {
		// 执行成功、提前结束
		statusEdit := app.Edit(tx, info.Id).Select("status", "status_desc", "number", "updated_ip", "updated_at", "result").Updates(Task{
			Status:     TASK_SUCCESS,
			StatusDesc: "结束执行",
			Number:     info.Number + 1,
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
			Result:     desc,
		})
		if statusEdit.RowsAffected == 0 {
			log.Println("statusEdit", statusEdit.Error)
		}
	}
	if status == CodeError {
		// 执行失败
		statusEdit := app.Edit(tx, info.Id).Select("status_desc", "number", "run_id", "updated_ip", "updated_at", "result").Updates(Task{
			StatusDesc: "执行失败",
			Number:     info.Number + 1,
			RunId:      gouuid.GetUuId(),
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
			Result:     desc,
		})
		if statusEdit.RowsAffected == 0 {
			log.Println("statusEdit", statusEdit.Error)
		}
	}
	if info.MaxNumber != 0 {
		if info.Number+1 >= info.MaxNumber {
			// 关闭执行
			statusEdit := app.Edit(tx, info.Id).Select("status").Updates(Task{
				Status: TASK_TIMEOUT,
			})
			if statusEdit.RowsAffected == 0 {
				log.Println("statusEdit", statusEdit.Error)
			}
		}
	}
}
