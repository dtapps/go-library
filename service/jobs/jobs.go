package jobs

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/goredis"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gouuid"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	RunVersion   int         `json:"run_version"`    // 运行版本
	Os           string      `json:"os"`             // 系统类型
	Arch         string      `json:"arch"`           // 系统架构
	MaxProCs     int         `json:"max_pro_cs"`     // CPU核数
	Version      string      `json:"version"`        // GO版本
	MacAddrS     string      `json:"mac_addr_s"`     // Mac地址
	InsideIp     string      `json:"inside_ip"`      // 内网ip
	OutsideIp    string      `json:"outside_ip"`     // 外网ip
	MainService  int         `json:"main_service"`   // 主要服务
	AddIpService int         `json:"add_ip_service"` // 添加IP
	Db           *gorm.DB    // 数据库
	Redis        goredis.App // 缓存数据库服务
}

// Add 添加任务
func (app *App) Add(Type string, params interface{}, frequency int64) int64 {
	return app.Db.Create(&Task{
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
	}).RowsAffected
}

// AddCustomId 添加任务
func (app *App) AddCustomId(Type string, params interface{}, frequency int64, customId string) int64 {
	query := app.TaskCustomIdTake(Type, customId)
	if query.Id != 0 {
		return 0
	}
	return app.Db.Create(&Task{
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
	}).RowsAffected
}

// AddCustomIdMaxNumber 添加任务并设置最大数量
func (app *App) AddCustomIdMaxNumber(Type string, params interface{}, frequency int64, customId string, maxNumber int64) int64 {
	query := app.TaskCustomIdTakeStatus(Type, customId, TASK_IN)
	if query.Id != 0 {
		return 0
	}
	return app.Db.Create(&Task{
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
	}).RowsAffected
}

type TaskParams = Task

// AddInOrder 添加订单可执行任务
func (app *App) AddInOrder(Type string, params interface{}, frequency int64) int64 {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsOrderType
	return app.AddIn(param, params)
}

// AddInOrderCustomId 添加订单可执行任务
func (app *App) AddInOrderCustomId(Type string, params interface{}, frequency int64, customId string) int64 {
	query := app.TaskCustomIdTakeStatus(Type, customId, TASK_IN)
	if query.Id != 0 {
		return 0
	}
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.CustomId = customId
	param.ParamsType = ParamsOrderType
	return app.AddIn(param, params)
}

// AddInOrderCustomIdSpecifyIp 添加订单可执行任务
func (app *App) AddInOrderCustomIdSpecifyIp(Type string, params interface{}, frequency int64, customId, specifyIp string) int64 {
	query := app.TaskCustomIdTakeStatus(Type, customId, TASK_IN)
	if query.Id != 0 {
		return 0
	}
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.CustomId = customId
	param.SpecifyIp = specifyIp
	param.ParamsType = ParamsOrderType
	return app.AddIn(param, params)
}

// AddInMerchantGoldenBean 添加商家金豆可执行任务
func (app *App) AddInMerchantGoldenBean(Type string, params interface{}, frequency int64) int64 {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsMerchantGoldenBeanType
	return app.AddIn(param, params)
}

// AddInTeamInv 添加团队邀请可执行任务
func (app *App) AddInTeamInv(Type string, params interface{}, frequency int64) int64 {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsTeamInvType
	return app.AddIn(param, params)
}

// AddInNewService 添加企业自定义可执行任务
func (app *App) AddInNewService(Type string, params interface{}, frequency int64) int64 {
	var param TaskParams
	param.Type = Type
	param.Frequency = frequency
	param.ParamsType = ParamsNewServiceType
	return app.AddIn(param, params)
}

// AddInNewServiceNext 添加企业自定义下一步可执行任务
func (app *App) AddInNewServiceNext(param TaskParams, params interface{}) int64 {
	param.ParamsType = ParamsNewServiceNextType
	return app.AddIn(param, params)
}

// AddInRepairMerchantAccountQuantity 添加可执行任务
func (app *App) AddInRepairMerchantAccountQuantity() int64 {
	var param TaskParams
	param.Type = TypeRepairMerchantAccountQuantityTesting
	param.Frequency = 90000
	return app.AddIn(param, map[string]interface{}{})
}

// AddInRepairMerchantAccountQuantityLevel 添加修复商家账号数量下一步可执行任务
func (app *App) AddInRepairMerchantAccountQuantityLevel(param TaskParams, params interface{}) int64 {
	param.ParamsType = ParamsRepairMerchantAccountQuantityLevelType
	return app.AddIn(param, params)
}

// AddInOrderCustomIdObservation 添加观察接口任务
func (app *App) AddInOrderCustomIdObservation(Type string, customId string) int64 {
	query := app.TaskCustomIdTakeStatus(Type, customId, TASK_IN)
	if query.Id != 0 {
		return query.Id
	}
	var param TaskParams
	param.Type = Type
	param.Frequency = 3600
	param.MaxNumber = 24 * 5 // 一个星期
	param.CustomId = customId
	param.ParamsType = ParamsOrderType
	return app.AddIn(param, ParamsOrderId{
		OrderId: customId,
	})
}

// AddInOrderCustomIdObservationClone 观察接口关闭
func (app *App) AddInOrderCustomIdObservationClone(Type string, customId string) int64 {
	query := app.TaskCustomIdTakeStatus(Type, customId, TASK_IN)
	if query.Id == 0 {
		return 0
	}
	return app.Edit(query.Id).Select("status", "status_desc", "run_id", "updated_ip", "updated_at").Updates(Task{
		Status:     TASK_SUCCESS,
		StatusDesc: "已完成，停止观察",
		RunId:      gouuid.GetUuId(),
		UpdatedIp:  app.OutsideIp,
		UpdatedAt:  gotime.Current().Format(),
	}).RowsAffected
}

// AddIn 添加可执行任务
// params.Type 任务类型
// params.Frequency 任务频率
// params.CustomId 自定义编号
// params 任务参数
func (app *App) AddIn(param TaskParams, params interface{}) int64 {
	param.Status = TASK_IN
	param.StatusDesc = "首次添加任务"
	param.RunId = gouuid.GetUuId()
	param.Params = gojson.JsonEncodeNoError(params)
	param.CreatedIp = app.OutsideIp
	param.UpdatedIp = app.OutsideIp
	param.CreatedAt = gotime.Current().Format()
	param.UpdatedAt = gotime.Current().Format()
	return app.Db.Create(&param).RowsAffected
}

// AddWaitNewServiceNext 添加企业自定义下一步等待执行任务
func (app *App) AddWaitNewServiceNext(param TaskParams, params interface{}) int64 {
	param.ParamsType = ParamsNewServiceNextType
	return app.AddWait(param, params)
}

// AddWaitRepairMerchantAccountQuantityLevel 添加修复商家账号数量下一步等待执行任务
func (app *App) AddWaitRepairMerchantAccountQuantityLevel(param TaskParams, params interface{}) int64 {
	param.ParamsType = ParamsRepairMerchantAccountQuantityLevelType
	return app.AddWait(param, params)
}

// AddInKashangwlPriceCustomId 添加刷新卡商网价格任务
func (app *App) AddInKashangwlPriceCustomId(productId int64) int64 {
	query := app.TaskCustomIdTakeStatus(TypeSyncGoodsPriceSingleKashangwl, fmt.Sprintf("%v", productId), TASK_IN)
	if query.Id != 0 {
		return query.Id
	}
	var param TaskParams
	param.Type = TypeSyncGoodsPriceSingleKashangwl
	param.Frequency = 1800
	param.CustomId = fmt.Sprintf("%v", productId)
	param.ParamsType = ParamsKashangwlType
	return app.AddIn(param, ParamsKashangwlId{
		ProductID: productId,
	})
}

// AddWait 添加等待执行任务
// params.Type 任务类型
// params.Frequency 任务频率
// params.CustomId 自定义编号
// params.CustomSequence 自定义顺序
// params 任务参数
func (app *App) AddWait(param TaskParams, params interface{}) int64 {
	param.Status = TASK_WAIT
	param.StatusDesc = "首次添加任务"
	param.RunId = gouuid.GetUuId()
	param.Params = gojson.JsonEncodeNoError(params)
	param.CreatedIp = app.OutsideIp
	param.UpdatedIp = app.OutsideIp
	param.CreatedAt = gotime.Current().Format()
	param.UpdatedAt = gotime.Current().Format()
	return app.Db.Create(&param).RowsAffected
}

// Edit 任务修改
func (app *App) Edit(id int64) *gorm.DB {
	return app.Db.Model(&Task{}).Where("id = ?", id)
}

// UpdateFrequency 更新任务频率
func (app *App) UpdateFrequency(id, frequency int64) *gorm.DB {
	return app.Edit(id).Updates(map[string]interface{}{
		"frequency": frequency,
	})
}

// Start 任务启动
func (app *App) Start(customId string, customSequence int64) int64 {
	return app.Db.Model(&Task{}).
		Where("custom_id = ?", customId).
		Where("custom_sequence = ?", customSequence).
		Where("status = ?", TASK_WAIT).
		Select("status", "status_desc", "updated_ip", "updated_at").Updates(Task{
		Status:     TASK_IN,
		StatusDesc: "启动任务",
		UpdatedIp:  app.OutsideIp,
		UpdatedAt:  gotime.Current().Format(),
	}).RowsAffected
}

// RunAddLog 任务执行日志
func (app *App) RunAddLog(id int64, runId string) int64 {
	return app.Db.Create(&TaskLogRun{
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
	}).RowsAffected
}

type CronParamsResp struct {
	ParamsOrderId
	ParamsMerchantUserIdOpenid
	ParamsTaskId
	ParamsTaskIdNext
	ParamsWechat
	ParamsTeamInv
	ParamsRepairMerchantAccountQuantityLevel
}

// Run 任务执行
func (app *App) Run(info Task, status int, desc string) {
	// 请求函数记录
	app.Db.Create(&TaskLog{
		TaskId:     info.Id,
		StatusCode: status,
		Desc:       desc,
		Version:    app.RunVersion,
		CreatedAt:  gotime.Current().Format(),
	})
	if status == 0 {
		app.Edit(info.Id).Select("run_id").Updates(Task{
			RunId: gouuid.GetUuId(),
		})
		return
	}
	// 任务
	if status == http.StatusOK {
		// 执行成功
		app.Edit(info.Id).Select("status_desc", "number", "run_id", "updated_ip", "updated_at", "result").Updates(Task{
			StatusDesc: "执行成功",
			Number:     info.Number + 1,
			RunId:      gouuid.GetUuId(),
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
			Result:     desc,
		})
	}
	if status == http.StatusCreated {
		// 执行成功、提前结束
		app.Edit(info.Id).Select("status", "status_desc", "number", "updated_ip", "updated_at", "result").Updates(Task{
			Status:     TASK_SUCCESS,
			StatusDesc: "结束执行",
			Number:     info.Number + 1,
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
			Result:     desc,
		})
	}
	if status == http.StatusInternalServerError {
		// 执行失败
		app.Edit(info.Id).Select("status_desc", "number", "run_id", "updated_ip", "updated_at", "result").Updates(Task{
			StatusDesc: "执行失败",
			Number:     info.Number + 1,
			RunId:      gouuid.GetUuId(),
			UpdatedIp:  app.OutsideIp,
			UpdatedAt:  gotime.Current().Format(),
			Result:     desc,
		})
	}
	if info.MaxNumber != 0 {
		if info.Number+1 >= info.MaxNumber {
			// 关闭执行
			app.Edit(info.Id).Select("status").Updates(Task{
				Status: TASK_TIMEOUT,
			})
		}
	}
}
