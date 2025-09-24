package gotime

import (
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
	ShortDateTimeFormat = "20060102150405"
	ShortDateFormat     = "20060102"
	ShortTimeFormat     = "150405"
)

// Pro 结构体
type Pro struct {
	Time time.Time
	loc  *time.Location
}

// NewPro 初始化结构体
func NewPro() Pro {
	return Pro{
		Time: time.Now(),
	}
}

// BeforeSeconds 获取 n 秒前的时间
func (p Pro) BeforeSeconds(n int) Pro {
	p.Time = p.Time.Add(-time.Duration(n) * time.Second)
	return p
}

// AfterSeconds 获取 n 秒后的时间
func (p Pro) AfterSeconds(n int) Pro {
	p.Time = p.Time.Add(time.Duration(n) * time.Second)
	return p
}

// BeforeMinute 获取 n 分钟前的时间
func (p Pro) BeforeMinute(n int) Pro {
	p.Time = p.Time.Add(-time.Duration(n) * time.Minute)
	return p
}

// AfterMinute 获取 n 分钟后的时间
func (p Pro) AfterMinute(n int) Pro {
	p.Time = p.Time.Add(time.Duration(n) * time.Minute)
	return p
}

// BeforeHour 获取 n 小时前的时间
func (p Pro) BeforeHour(n int) Pro {
	p.Time = p.Time.Add(-time.Duration(n) * time.Hour)
	return p
}

// AfterHour 获取 n 小时后的时间
func (p Pro) AfterHour(n int) Pro {
	p.Time = p.Time.Add(time.Duration(n) * time.Hour)
	return p
}

// BeforeDay 获取 n 天前的时间
func (p Pro) BeforeDay(day int) Pro {
	p.Time = p.Time.AddDate(0, 0, -day)
	return p
}

// AfterDay 获取 n 天后的时间
func (p Pro) AfterDay(day int) Pro {
	p.Time = p.Time.AddDate(0, 0, day)
	return p
}

// BeforeWeek 获取 n 周前的时间
func (p Pro) BeforeWeek(weeks int) Pro {
	p.Time = p.Time.AddDate(0, 0, -weeks*7)
	return p
}

// AfterWeek 获取 n 周后的时间
func (p Pro) AfterWeek(weeks int) Pro {
	p.Time = p.Time.AddDate(0, 0, weeks*7)
	return p
}

// BeforeMonth 获取 n 个月前的时间
func (p Pro) BeforeMonth(months int) Pro {
	p.Time = p.Time.AddDate(0, -months, 0)
	return p
}

// AfterMonth 获取 n 个月后的时间
func (p Pro) AfterMonth(months int) Pro {
	p.Time = p.Time.AddDate(0, months, 0)
	return p
}

// BeforeQuarter 获取 n 个季度前的时间（1 季度 = 3 个月）
func (p Pro) BeforeQuarter(quarters int) Pro {
	p.Time = p.Time.AddDate(0, -quarters*3, 0)
	return p
}

// AfterQuarter 获取 n 个季度后的时间（1 季度 = 3 个月）
func (p Pro) AfterQuarter(quarters int) Pro {
	p.Time = p.Time.AddDate(0, quarters*3, 0)
	return p
}

// BeforeYear 获取 n 年前的时间
func (p Pro) BeforeYear(years int) Pro {
	p.Time = p.Time.AddDate(-years, 0, 0)
	return p
}

// AfterYear 获取 n 年后的时间
func (p Pro) AfterYear(years int) Pro {
	p.Time = p.Time.AddDate(years, 0, 0)
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
