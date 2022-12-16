package gojobs

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gotime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TaskLog 任务日志模型
type TaskLog struct {
	LogId   primitive.ObjectID `json:"log_id,omitempty" bson:"_id,omitempty"` //【记录】编号
	LogTime primitive.DateTime `json:"log_time,omitempty" bson:"log_time"`    //【记录】时间
	Task    struct {
		Id         uint   `json:"id" bson:"id"`                   //【任务】编号
		RunId      string `json:"run_id" bson:"run_id"`           //【任务】执行编号
		ResultCode int    `json:"result_code" bson:"result_code"` //【任务】执行状态码
		ResultDesc string `json:"result_desc" bson:"result_desc"` //【任务】执行结果
		ResultTime string `json:"result_time" bson:"result_time"` //【任务】执行时间
	} `json:"task,omitempty" bson:"task,omitempty"` //【任务】信息
	System struct {
		Hostname        string  `json:"hostname" bson:"hostname"`                                       //【系统】主机名
		Os              string  `json:"os" bson:"os"`                                                   //【系统】系统类型
		Version         string  `json:"version" bson:"version"`                                         //【系统】系统版本
		Kernel          string  `json:"kernel" bson:"kernel"`                                           //【系统】系统内核
		KernelVersion   string  `json:"kernel_version" bson:"kernel_version"`                           //【系统】系统内核版本
		BootTime        string  `json:"boot_time" bson:"boot_time"`                                     //【系统】系统开机时间
		CpuCores        int     `json:"cpu_cores,omitempty" bson:"cpu_cores,omitempty"`                 //【系统】CPU核数
		CpuModelName    string  `json:"cpu_model_name,omitempty" bson:"cpu_model_name,omitempty"`       //【系统】CPU型号名称
		CpuMhz          float64 `json:"cpu_mhz,omitempty" bson:"cpu_mhz,omitempty"`                     //【系统】CPU兆赫
		InsideIp        string  `json:"inside_ip" bson:"inside_ip"`                                     //【系统】内网ip
		OutsideIp       string  `json:"outside_ip" bson:"outside_ip"`                                   //【系统】外网ip
		GoVersion       string  `json:"go_version,omitempty" bson:"go_version,omitempty"`               //【系统】go版本
		SdkVersion      string  `json:"sdk_version,omitempty" bson:"sdk_version,omitempty"`             //【系统】sdk版本
		MongoVersion    string  `json:"mongo_version,omitempty" bson:"mongo_version,omitempty"`         //【系统】mongo版本
		MongoSdkVersion string  `json:"mongo_sdk_version,omitempty" bson:"mongo_sdk_version,omitempty"` //【系统】mongo sdk版本
		RedisVersion    string  `json:"redis_version,omitempty" bson:"redis_version,omitempty"`         //【系统】redis版本
		RedisSdkVersion string  `json:"redis_sdk_version,omitempty" bson:"redis_sdk_version,omitempty"` //【系统】redis sdk版本
		LogVersion      string  `json:"log_version,omitempty" bson:"log_version,omitempty"`             //【系统】log版本
	} `json:"system,omitempty" bson:"system,omitempty"` //【系统】信息
}

func (TaskLog) CollectionName() string {
	return "task_log"
}

// 创建时间序列集合
func (TaskLog) createCollection(ctx context.Context, zapLog *golog.ZapLog, db *dorm.MongoClient, databaseName string) {
	err := db.Database(databaseName).CreateCollection(ctx, TaskLog{}.CollectionName(), options.CreateCollection().SetTimeSeriesOptions(options.TimeSeries().SetTimeField("log_time")))
	if err != nil {
		zapLog.WithTraceId(ctx).Sugar().Errorf("创建时间序列集合：%s", err)
	}
}

// 创建索引
func (TaskLog) createIndexes(ctx context.Context, zapLog *golog.ZapLog, db *dorm.MongoClient, databaseName string) {
	_, err := db.Database(databaseName).Collection(TaskLog{}.CollectionName()).CreateManyIndexes(ctx, []mongo.IndexModel{{
		Keys: bson.D{{
			Key:   "log_time",
			Value: -1,
		}},
	}})
	if err != nil {
		zapLog.WithTraceId(ctx).Sugar().Errorf("创建索引：%s", err)
	}
}

// MongoTaskLogRecord 记录
func (c *Client) MongoTaskLogRecord(ctx context.Context, task jobs_gorm_model.Task, runId string, taskResultCode int, taskResultDesc string) {

	taskLog := TaskLog{
		LogId:   primitive.NewObjectID(),
		LogTime: primitive.NewDateTimeFromTime(gotime.Current().Time),
	}

	taskLog.Task.Id = task.Id
	taskLog.Task.RunId = runId
	taskLog.Task.ResultCode = taskResultCode
	taskLog.Task.ResultDesc = taskResultDesc
	taskLog.Task.ResultTime = gotime.Current().Format()

	taskLog.System.Hostname = c.config.systemHostname                                                                //【系统】主机名
	taskLog.System.Os = c.config.systemOs                                                                            //【系统】系统类型
	taskLog.System.Version = c.config.systemVersion                                                                  //【系统】系统版本
	taskLog.System.Kernel = c.config.systemKernel                                                                    //【系统】系统内核
	taskLog.System.KernelVersion = c.config.systemKernelVersion                                                      //【系统】系统内核版本
	taskLog.System.BootTime = gotime.SetCurrent(gotime.SetCurrentUnix(int64(c.config.systemBootTime)).Time).Format() //【系统】系统开机时间
	taskLog.System.CpuCores = c.config.cpuCores                                                                      //【系统】CPU核数
	taskLog.System.CpuModelName = c.config.cpuModelName                                                              //【程序】CPU型号名称
	taskLog.System.CpuMhz = c.config.cpuMhz                                                                          //【系统】CPU兆赫
	taskLog.System.InsideIp = c.config.systemInsideIp                                                                //【系统】内网ip
	taskLog.System.OutsideIp = c.config.systemOutsideIp                                                              //【系统】外网ip
	taskLog.System.GoVersion = c.config.goVersion                                                                    //【系统】Go版本
	taskLog.System.SdkVersion = c.config.sdkVersion                                                                  //【系统】Sdk版本
	taskLog.System.MongoVersion = c.config.mongoVersion                                                              //【系统】mongo版本
	taskLog.System.MongoSdkVersion = c.config.mongoSdkVersion                                                        //【系统】mongo sdk版本
	taskLog.System.RedisVersion = c.config.redisVersion                                                              //【系统】redis版本
	taskLog.System.RedisSdkVersion = c.config.redisSdkVersion                                                        //【系统】redis sdk版本
	taskLog.System.LogVersion = c.config.logVersion                                                                  //【系统】log版本

	_, err := c.mongoClient.Database(c.mongoConfig.databaseName).Collection(TaskLog{}.CollectionName()).InsertOne(ctx, taskLog)
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("记录失败：%s", err)
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("记录数据：%+v", taskLog)
	}

}

// MongoTaskLogDelete 删除
func (c *Client) MongoTaskLogDelete(ctx context.Context, hour int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{"log_time", bson.D{{"$lt", primitive.NewDateTimeFromTime(gotime.Current().BeforeHour(hour).Time)}}}}
	return c.mongoClient.Database(c.mongoConfig.databaseName).Collection(TaskLog{}.CollectionName()).DeleteMany(ctx, filter)
}
