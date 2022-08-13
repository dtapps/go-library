package gotime

import (
	"fmt"
	"time"
)

// 时间格式化常量
const (
	RFC3339Format       = time.RFC3339
	Iso8601Format       = "2006-01-02T15:04:05-07:00"
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC7231Format       = "Mon, 02 Jan 2006 15:04:05 GMT"
	DayDateTimeFormat   = "Mon, Jan 2, 2006 3:04 PM"
	DateTimeFormat      = "2006-01-02 15:04:05"
	DateFormat          = "2006-01-02"
	TimeFormat          = "15:04:05"
	ShortDateTimeFormat = "20060102150405"
	ShortDateFormat     = "20060102"
	ShortTimeFormat     = "150405"
)

// Pro 结构体
type Pro struct {
	Time  time.Time
	loc   *time.Location
	Error error
}

// NewPro 初始化结构体
func NewPro() Pro {
	return Pro{
		Time: time.Now(),
	}
}

// BeforeSeconds 获取n秒前的时间
func (p Pro) BeforeSeconds(seconds int64) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("-%ds", seconds))
	p.Time = p.Time.Add(st)
	return p
}

// AfterSeconds 获取n秒后的时间
func (p Pro) AfterSeconds(seconds int64) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("+%ds", seconds))
	p.Time = p.Time.Add(st)
	return p
}

// BeforeMinute 获取n分钟前的时间
func (p Pro) BeforeMinute(seconds int64) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("-%dm", seconds))
	p.Time = p.Time.Add(st)
	return p
}

// AfterMinute 获取n分钟后的时间
func (p Pro) AfterMinute(seconds int64) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("+%dm", seconds))
	p.Time = p.Time.Add(st)
	return p
}

// BeforeHour 获取n小时前的时间
func (p Pro) BeforeHour(hour int64) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("-%dh", hour))
	p.Time = p.Time.Add(st)
	return p
}

// AfterHour 获取n小时后的时间
func (p Pro) AfterHour(hour int64) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("+%dh", hour))
	p.Time = p.Time.Add(st)
	return p
}

// BeforeDay 获取n天前的时间
func (p Pro) BeforeDay(day int) Pro {
	p.Time = p.Time.AddDate(0, 0, -day)
	return p
}

// AfterDay 获取n天后的时间
func (p Pro) AfterDay(day int) Pro {
	p.Time = p.Time.AddDate(0, 0, day)
	return p
}

// SetFormat 格式化
func (p Pro) SetFormat(layout string) string {
	return p.Time.Format(layout)
}

// Month 获取当前月
func (p Pro) Month() int64 {
	return p.MonthOfYear()
}

// MonthOfYear 获取本年的第几月
func (p Pro) MonthOfYear() int64 {
	return int64(p.Time.In(p.loc).Month())
}
