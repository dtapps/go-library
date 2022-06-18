package gojobs

import (
	"fmt"
	"net/http"
)

const (
	CodeAbnormal = 0                              // 异常
	CodeError    = http.StatusInternalServerError // 失败
	CodeSuccess  = http.StatusOK                  // 成功
	CodeEnd      = http.StatusCreated             // 结束
)

// 每隔n秒执行一次
const specSeconds = "*/%d * * * * *"

// GetSpecSeconds 每隔n秒执行一次
var GetSpecSeconds = func(n int64) string {
	if n < 0 || n > 59 {
		return ""
	}
	return fmt.Sprintf(specSeconds, n)
}

// GetFrequencySeconds 每隔n秒执行一次
var GetFrequencySeconds = func(n int64) int64 {
	if n < 0 || n > 59 {
		return -1
	}
	return n
}

// 每隔n分钟执行一次
const specMinutes = "0 */%d * * * *"

// GetSpecMinutes 每隔n分钟执行一次
var GetSpecMinutes = func(n int64) string {
	if n < 0 || n > 59 {
		return ""
	}
	return fmt.Sprintf(specMinutes, n)
}

// GetFrequencyMinutes 每隔n分钟执行一次
var GetFrequencyMinutes = func(n int64) int64 {
	if n < 0 || n > 59 {
		return -1
	}
	return n * 60
}

// 每天n点执行一次
const specHour = "0 0 */%d * * *"

// GetSpecHour 每天n点执行一次
var GetSpecHour = func(n int64) string {
	if n < 0 || n > 23 {
		return ""
	}
	return fmt.Sprintf(specHour, n)
}

// GetFrequencyHour 每天n点执行一次
var GetFrequencyHour = func(n int64) int64 {
	if n < 0 || n > 23 {
		return -1
	}
	return n * 60 * 60
}
