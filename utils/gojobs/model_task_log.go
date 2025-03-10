package gojobs

import "time"

// ModelTaskLog 任务日志
type ModelTaskLog struct {
	LogID           int64     `json:"log_id,omitempty"`            //【日志】编号
	LogTime         time.Time `json:"log_time,omitempty"`          //【日志】时间
	TaskID          uint      `json:"task_id,omitempty"`           //【任务】编号
	TaskRunID       string    `json:"task_run_id,omitempty"`       //【任务】执行编号
	TaskResultCode  int       `gjson:"task_result_code,omitempty"` //【任务】执行状态码
	TaskResultDesc  string    `json:"task_result_desc,omitempty"`  //【任务】执行结果
	SystemInsideIP  string    `json:"system_inside_ip,omitempty"`  //【系统】内网IP
	SystemOutsideIP string    `json:"system_outside_ip,omitempty"` //【系统】外网IP
}
