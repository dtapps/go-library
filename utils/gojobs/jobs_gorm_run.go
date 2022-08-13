package gojobs

import (
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gostring"
	"gorm.io/gorm"
	"log"
)

// Run 运行
func (j *JobsGorm) Run(info jobs_gorm_model.Task, status int, desc string) {
	// 请求函数记录
	err := j.gormClient.Db.Create(&jobs_gorm_model.TaskLog{
		TaskId:     info.Id,
		StatusCode: status,
		Desc:       desc,
		Version:    j.config.runVersion,
	}).Error
	if err != nil {
		log.Println("statusCreate", err.Error())
	}
	if status == 0 {
		err = j.EditTask(j.gormClient.Db, info.Id).
			Select("run_id").
			Updates(jobs_gorm_model.Task{
				RunId: gostring.GetUuId(),
			}).Error
		if err != nil {
			log.Println("statusEdit", err.Error())
		}
		return
	}
	// 任务
	if status == CodeSuccess {
		// 执行成功
		err = j.EditTask(j.gormClient.Db, info.Id).
			Select("status_desc", "number", "run_id", "updated_ip", "result").
			Updates(jobs_gorm_model.Task{
				StatusDesc: "执行成功",
				Number:     info.Number + 1,
				RunId:      gostring.GetUuId(),
				UpdatedIp:  j.config.outsideIp,
				Result:     desc,
			}).Error
		if err != nil {
			log.Println("statusEdit", err.Error())
		}
	}
	if status == CodeEnd {
		// 执行成功、提前结束
		err = j.EditTask(j.gormClient.Db, info.Id).
			Select("status", "status_desc", "number", "updated_ip", "result").
			Updates(jobs_gorm_model.Task{
				Status:     TASK_SUCCESS,
				StatusDesc: "结束执行",
				Number:     info.Number + 1,
				UpdatedIp:  j.config.outsideIp,
				Result:     desc,
			}).Error
		if err != nil {
			log.Println("statusEdit", err.Error())
		}
	}
	if status == CodeError {
		// 执行失败
		err = j.EditTask(j.gormClient.Db, info.Id).
			Select("status_desc", "number", "run_id", "updated_ip", "result").
			Updates(jobs_gorm_model.Task{
				StatusDesc: "执行失败",
				Number:     info.Number + 1,
				RunId:      gostring.GetUuId(),
				UpdatedIp:  j.config.outsideIp,
				Result:     desc,
			}).Error
		if err != nil {
			log.Println("statusEdit", err.Error())
		}
	}
	if info.MaxNumber != 0 {
		if info.Number+1 >= info.MaxNumber {
			// 关闭执行
			err = j.EditTask(j.gormClient.Db, info.Id).
				Select("status").
				Updates(jobs_gorm_model.Task{
					Status: TASK_TIMEOUT,
				}).Error
			if err != nil {
				log.Println("statusEdit", err.Error())
			}
		}
	}
}

// RunAddLog 任务执行日志
func (j *JobsGorm) RunAddLog(id uint, runId string) error {
	return j.gormClient.Db.Create(&jobs_gorm_model.TaskLogRun{
		TaskId:     id,
		RunId:      runId,
		InsideIp:   j.config.insideIp,
		OutsideIp:  j.config.outsideIp,
		Os:         j.config.os,
		Arch:       j.config.arch,
		Gomaxprocs: j.config.maxProCs,
		GoVersion:  j.config.version,
		MacAddrs:   j.config.macAddrS,
	}).Error
}

// ConfigCreateInCustomId 创建正在运行任务
type ConfigCreateInCustomId struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前ip
}

// CreateInCustomId 创建正在运行任务
func (j *JobsGorm) CreateInCustomId(config *ConfigCreateInCustomId) error {
	if config.CurrentIp == "" {
		config.CurrentIp = j.config.outsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}

// ConfigCreateInCustomIdOnly 创建正在运行唯一任务
type ConfigCreateInCustomIdOnly struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前ip
}

// CreateInCustomIdOnly 创建正在运行唯一任务
func (j *JobsGorm) CreateInCustomIdOnly(config *ConfigCreateInCustomIdOnly) error {
	query := j.TaskTypeTakeIn(config.Tx, config.CustomId, config.Type)
	if query.Id != 0 {
		return errors.New(fmt.Sprintf("%d:[%s@%s]任务已存在", query.Id, config.CustomId, config.Type))
	}
	if config.CurrentIp == "" {
		config.CurrentIp = j.config.outsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}

// ConfigCreateInCustomIdMaxNumber 创建正在运行任务并限制数量
type ConfigCreateInCustomIdMaxNumber struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	MaxNumber      int64    // 最大次数
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前ip
}

// CreateInCustomIdMaxNumber 创建正在运行任务并限制数量
func (j *JobsGorm) CreateInCustomIdMaxNumber(config *ConfigCreateInCustomIdMaxNumber) error {
	if config.CurrentIp == "" {
		config.CurrentIp = j.config.outsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		MaxNumber:      config.MaxNumber,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}

// ConfigCreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
type ConfigCreateInCustomIdMaxNumberOnly struct {
	Tx             *gorm.DB // 驱动
	Params         string   // 参数
	Frequency      int64    // 频率(秒单位)
	MaxNumber      int64    // 最大次数
	CustomId       string   // 自定义编号
	CustomSequence int64    // 自定义顺序
	Type           string   // 类型
	TypeName       string   // 类型名称
	SpecifyIp      string   // 指定外网IP
	CurrentIp      string   // 当前ip
}

// CreateInCustomIdMaxNumberOnly 创建正在运行唯一任务并限制数量
func (j *JobsGorm) CreateInCustomIdMaxNumberOnly(config *ConfigCreateInCustomIdMaxNumberOnly) error {
	query := j.TaskTypeTakeIn(config.Tx, config.CustomId, config.Type)
	if query.Id != 0 {
		return errors.New(fmt.Sprintf("%d:[%s@%s]任务已存在", query.Id, config.CustomId, config.Type))
	}
	if config.CurrentIp == "" {
		config.CurrentIp = j.config.outsideIp
	}
	err := config.Tx.Create(&jobs_gorm_model.Task{
		Status:         TASK_IN,
		Params:         config.Params,
		StatusDesc:     "首次添加任务",
		Frequency:      config.Frequency,
		MaxNumber:      config.MaxNumber,
		RunId:          gostring.GetUuId(),
		CustomId:       config.CustomId,
		CustomSequence: config.CustomSequence,
		Type:           config.Type,
		TypeName:       config.TypeName,
		CreatedIp:      config.CurrentIp,
		SpecifyIp:      config.SpecifyIp,
		UpdatedIp:      config.CurrentIp,
	}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("创建[%s@%s]任务失败：%s", config.CustomId, config.Type, err.Error()))
	}
	return nil
}
