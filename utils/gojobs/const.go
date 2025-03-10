package gojobs

import "net/http"

const (
	SpecifyIpNull = "0.0.0.0"
)

const (
	// CodeAbnormal 异常
	CodeAbnormal = 0
	// CodeConflict 冲突
	CodeConflict = http.StatusBadRequest
	// CodeObservation 观查
	CodeObservation = http.StatusMultipleChoices
	// CodeError 失败
	CodeError = http.StatusInternalServerError
	// CodeSuccess 成功
	CodeSuccess = http.StatusOK
	// CodeEnd 结束
	CodeEnd = http.StatusCreated
)

const (
	// TASK_IN 任务运行
	TASK_IN = "IN"
	// TASK_CONFLICT 冲突
	TASK_CONFLICT = "CONFLICT"
	// TASK_OBSERVATION 观查
	TASK_OBSERVATION = "OBSERVATION"
	// TASK_SUCCESS 任务完成
	TASK_SUCCESS = "SUCCESS"
	// TASK_ERROR 任务异常
	TASK_ERROR = "ERROR"
	// TASK_TIMEOUT 任务超时
	TASK_TIMEOUT = "TIMEOUT"
	// TASK_WAIT 任务等待
	TASK_WAIT = "WAIT"
)
