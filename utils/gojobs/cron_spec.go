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

// 每隔n秒执行一次
type seconds struct {
	n int64
}

// GetSeconds 每隔n秒执行一次
func GetSeconds(n int64) *seconds {
	s := seconds{}
	s.n = n
	return &s
}

// Spec 每隔n秒执行一次
func (s seconds) Spec() string {
	if s.n < 0 || s.n > 59 {
		return ""
	}
	return fmt.Sprintf(specSeconds, s.n)
}

// Frequency 每隔n秒执行一次
func (s seconds) Frequency() int64 {
	if s.n < 0 || s.n > 59 {
		return -1
	}
	return s.n
}

// 每隔n分钟执行一次
const specMinutes = "0 */%d * * * *"

// 每隔n分钟执行一次
type minutes struct {
	n int64
}

// GetMinutes 每隔n分钟执行一次
func GetMinutes(n int64) *minutes {
	s := minutes{}
	s.n = n
	return &s
}

// Spec 每隔n分钟执行一次
func (s minutes) Spec() string {
	if s.n < 0 || s.n > 59 {
		return ""
	}
	return fmt.Sprintf(specMinutes, s.n)
}

// Frequency 每隔n分钟执行一次
func (s minutes) Frequency() int64 {
	if s.n < 0 || s.n > 59 {
		return -1
	}
	return s.n * 60
}

// 每天n点执行一次
const specHour = "0 0 */%d * * *"

// 每天n点执行一次
type hour struct {
	n int64
}

// GetHour 每天n点执行一次
func GetHour(n int64) *hour {
	s := hour{}
	s.n = n
	return &s
}

// Spec 每天n点执行一次
func (s hour) Spec() string {
	if s.n < 0 || s.n > 23 {
		return ""
	}
	return fmt.Sprintf(specHour, s.n)
}

// Frequency 每天n点执行一次
func (s hour) Frequency() int64 {
	if s.n < 0 || s.n > 23 {
		return -1
	}
	return s.n * 60 * 60
}
