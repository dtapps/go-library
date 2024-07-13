package gojobs

import (
	"fmt"
	"net/http"
)

const (
	// CodeAbnormal 异常
	CodeAbnormal = 0
	// CodeError 失败
	CodeError = http.StatusInternalServerError
	// CodeSuccess 成功
	CodeSuccess = http.StatusOK
	// CodeEnd 结束
	CodeEnd = http.StatusCreated
)

// 每隔n秒执行一次
const specSeconds = "*/%d * * * * *"

// Seconds 每隔n秒执行一次
type Seconds struct {
	n int64
}

// GetSeconds 每隔n秒执行一次
func GetSeconds(n int64) *Seconds {
	s := Seconds{}
	s.n = n
	return &s
}

// Spec 每隔n秒执行一次
func (s Seconds) Spec() string {
	if s.n < 0 || s.n > 59 {
		return ""
	}
	return fmt.Sprintf(specSeconds, s.n)
}

// Frequency 每隔n秒执行一次
func (s Seconds) Frequency() int64 {
	if s.n < 0 || s.n > 59 {
		return -1
	}
	return s.n
}

// 每隔n分钟执行一次
const specMinutes = "0 */%d * * * *"

// Minutes 每隔n分钟执行一次
type Minutes struct {
	n int64
}

// GetMinutes 每隔n分钟执行一次
func GetMinutes(n int64) *Minutes {
	s := Minutes{}
	s.n = n
	return &s
}

// Spec 每隔n分钟执行一次
func (s Minutes) Spec() string {
	if s.n < 0 || s.n > 59 {
		return ""
	}
	return fmt.Sprintf(specMinutes, s.n)
}

// Frequency 每隔n分钟执行一次
func (s Minutes) Frequency() int64 {
	if s.n < 0 || s.n > 59 {
		return -1
	}
	return s.n * 60
}

// 每天n点执行一次
const specHour = "0 0 */%d * * *"

// Hour 每天n点执行一次
type Hour struct {
	n int64
}

// GetHour 每天n点执行一次
func GetHour(n int64) *Hour {
	s := Hour{}
	s.n = n
	return &s
}

// Spec 每天n点执行一次
func (s Hour) Spec() string {
	if s.n < 0 || s.n > 23 {
		return ""
	}
	return fmt.Sprintf(specHour, s.n)
}

// Frequency 每天n点执行一次
func (s Hour) Frequency() int64 {
	if s.n < 0 || s.n > 23 {
		return -1
	}
	return s.n * 60 * 60
}

// 每隔n小时执行一次
const specHourInterval = "0 0 %s * * *"

// HourInterval 每隔n小时执行一次
type HourInterval struct {
	n int64
}

// GetHourInterval 每隔n小时执行一次
func GetHourInterval(n int64) *HourInterval {
	s := HourInterval{}
	s.n = n
	return &s
}

// Spec 每隔n小时执行一次
func (s HourInterval) Spec() string {

	if s.n < 0 || s.n > 23 {
		return ""
	}

	// 循环出最近24次执行时间
	var sl []int64
	var i int64
	i = 0
	for {
		if i > 23 {
			break
		}
		sl = append(sl, s.n*i)
		i++
	}

	// TODO 可以合并两个

	// 过滤数据
	str := ""
	for _, v := range sl {
		if v > 23 {
			continue
		}
		str = fmt.Sprintf("%s,%v", str, v)
	}

	if len(str) <= 0 {
		return ""
	}

	return fmt.Sprintf(specHourInterval, str[1:])
}

// Frequency 每隔n小时执行一次
func (s HourInterval) Frequency() int64 {
	if s.n < 0 || s.n > 23 {
		return -1
	}
	return s.n * 60 * 60
}
