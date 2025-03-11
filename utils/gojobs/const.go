package gojobs

import "net/http"

const (
	SpecifyIpNull = "0.0.0.0"
)

const (
	// CodeAbnormal 异常
	CodeAbnormal = 0
	// CodeConflict 冲突
	CodeConflict = http.StatusConflict // http.StatusBadRequest 改 http.StatusConflict
	// CodeObservation 观查
	CodeObservation = http.StatusProcessing // http.StatusMultipleChoices 改 http.StatusProcessing
	// CodeError 失败
	CodeError = http.StatusInternalServerError
	// CodeSuccess 成功
	CodeSuccess = http.StatusOK
	// CodeEnd 完成
	CodeEnd = http.StatusNoContent // http.StatusCreated 改 http.StatusNoContent
	// CodeTimeout 超时
	CodeTimeout = http.StatusGatewayTimeout
	// CodeWait 等待
	CodeWait = http.StatusAccepted
)

const (
	// TASK_CONFLICT 冲突
	TASK_CONFLICT = "CONFLICT"
	// TASK_OBSERVATION 观查
	TASK_OBSERVATION = "OBSERVATION"
	// TASK_ERROR 失败
	TASK_ERROR = "ERROR"
	// TASK_IN 成功
	TASK_IN = "IN"
	// TASK_SUCCESS 完成
	TASK_SUCCESS = "SUCCESS"
	// TASK_TIMEOUT 超时
	TASK_TIMEOUT = "TIMEOUT"
	// TASK_WAIT 等待
	TASK_WAIT = "WAIT"
)
