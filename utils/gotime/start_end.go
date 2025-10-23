package gotime

import "time"

// 数字常量
const (
	YearsPerMillennium         = 1000    // 每千年1000年
	YearsPerCentury            = 100     // 每世纪100年
	YearsPerDecade             = 10      // 每十年10年
	QuartersPerYear            = 4       // 每年4季度
	MonthsPerYear              = 12      // 每年12月
	MonthsPerQuarter           = 3       // 每季度3月
	WeeksPerNormalYear         = 52      // 每常规年52周
	WeeksPerLongYear           = 53      // 每长年53周
	WeeksPerMonth              = 4       // 每月4周
	DaysPerLeapYear            = 366     // 每闰年366天
	DaysPerNormalYear          = 365     // 每常规年365天
	DaysPerWeek                = 7       // 每周7天
	HoursPerWeek               = 168     // 每周168小时
	HoursPerDay                = 24      // 每天24小时
	MinutesPerDay              = 1440    // 每天1440分钟
	MinutesPerHour             = 60      // 每小时60分钟
	SecondsPerWeek             = 604800  // 每周604800秒
	SecondsPerDay              = 86400   // 每天86400秒
	SecondsPerHour             = 3600    // 每小时3600秒
	SecondsPerMinute           = 60      // 每分钟60秒
	MillisecondsPerSecond      = 1000    // 每秒1000毫秒
	MicrosecondsPerMillisecond = 1000    // 每毫秒1000微秒
	MicrosecondsPerSecond      = 1000000 // 每秒1000000微秒
)

// StartOfCentury 本世纪开始时间
func (p Pro) StartOfCentury() Pro {
	p.Time = time.Date(p.Time.Year()/YearsPerCentury*YearsPerCentury, 1, 1, 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfCentury 本世纪结束时间
func (p Pro) EndOfCentury() Pro {
	p.Time = time.Date(p.Time.Year()/YearsPerCentury*YearsPerCentury+99, 12, 31, 23, 59, 59, 999999999, p.Time.Location())
	return p
}

// StartOfDecade 本年代开始时间
func (p Pro) StartOfDecade() Pro {
	p.Time = time.Date(p.Time.Year()/YearsPerDecade*YearsPerDecade, 1, 1, 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfDecade 本年代结束时间
func (p Pro) EndOfDecade() Pro {
	p.Time = time.Date(p.Time.Year()/YearsPerDecade*YearsPerDecade+9, 12, 31, 23, 59, 59, 999999999, p.Time.Location())
	return p
}

// StartOfYear 本年开始时间
func (p Pro) StartOfYear() Pro {
	p.Time = time.Date(p.Time.Year(), 1, 1, 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfYear 本年结束时间
func (p Pro) EndOfYear() Pro {
	p.Time = time.Date(p.Time.Year(), 12, 31, 23, 59, 59, 999999999, p.Time.Location())
	return p
}

// Quarter 获取当前季度（1~4）
func (p Pro) Quarter() int {
	return (int(p.Time.Month()) + 2) / 3
}

// StartOfQuarter 本季度开始时间
func (p Pro) StartOfQuarter() Pro {
	p.Time = time.Date(p.Time.Year(), time.Month(3*p.Quarter()-2), 1, 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfQuarter 本季度结束时间
func (p Pro) EndOfQuarter() Pro {
	next := time.Date(p.Time.Year(), time.Month(3*p.Quarter()+1), 1, 0, 0, 0, 0, p.Time.Location())
	p.Time = next.Add(-time.Nanosecond)
	return p
}

// StartOfMonth 本月开始时间
func (p Pro) StartOfMonth() Pro {
	p.Time = time.Date(p.Time.Year(), time.Month(p.Month()), 1, 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfMonth 本月结束时间
func (p Pro) EndOfMonth() Pro {
	next := time.Date(p.Time.Year(), time.Month(p.Month())+1, 1, 0, 0, 0, 0, p.Time.Location())
	p.Time = next.Add(-time.Nanosecond)
	return p
}

// StartOfWeek 本周开始时间
func (p Pro) StartOfWeek() Pro {
	t := p.Time
	offset := (int(t.Weekday()) + 6) % 7
	p.Time = time.Date(t.Year(), t.Month(), t.Day()-offset, 0, 0, 0, 0, t.Location())
	return p
}

// EndOfWeek 本周结束时间
func (p Pro) EndOfWeek() Pro {
	start := p.StartOfWeek().Time
	p.Time = time.Date(start.Year(), start.Month(), start.Day()+6, 23, 59, 59, 999999999, start.Location())
	return p
}

// StartOfDay 本日开始时间
func (p Pro) StartOfDay() Pro {
	p.Time = time.Date(p.Time.Year(), p.Time.Month(), p.Time.Day(), 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfDay 本日结束时间
func (p Pro) EndOfDay() Pro {
	p.Time = time.Date(p.Time.Year(), p.Time.Month(), p.Time.Day(), 23, 59, 59, 999999999, p.Time.Location())
	return p
}
