package gotime

import "time"

// 时间格式化
const (
	FormatYearMonthDayHourMinuteSeconds = "20060102150405"
	FormatYearMonthDayHourMinute        = "200601021504"
	FormatYearMonthDayHour              = "2006010215"
	FormatYearMonthDay                  = "20060102"
	FormatYearMonth                     = "200601"
	FormatYear                          = "2006"
)

// 时间格式化
const (
	FormatHourMinuteSeconds = "150405"
)

const (
	DateTimeFormat        = time.DateTime // "2006-01-02 15:04:05"
	DateTimeSFormat       = "2006-01-0215:04:05"
	DateTimeShrinkFormat  = "2006-01-02 15:04"
	DateTimeShrinkSFormat = "2006-01-0215:04"
	DateFormat            = time.DateOnly // "2006-01-02"
	TimeFormat            = time.TimeOnly //  "15:04:05"
	TimeShrinkFormat      = "15:04"

	DateYearMonthDayFormat = time.DateOnly //"2006-01-02"
	DateYearMonthFormat    = "2006-01"
)

const (
	DateTimeZhFormat        = "2006年01月02日 15点04分05秒"
	DateTimeZhSFormat       = "2006年01月02日15点04分05秒"
	DateTimeZhShrinkFormat  = "2006年01月02日 15点04分"
	DateTimeZhShrinkSFormat = "2006年01月02日15点04分"
	DateZhFormat            = "2006年01月02日"
	TimeZhFormat            = "15点04分05秒"
	TimeZhShrinkFormat      = "15点04分"

	DateZhYearMonthDayFormat = "2006年01月02日"
	DateZhYearMonthFormat    = "2006年01月"
)
