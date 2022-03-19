// @Package carbon
// @Description a simple, semantic and developer-friendly golang package for datetime
// @Page github.com/golang-module/carbon
// @Version v1.5.5
// @Author gouguoyin
// @Blog www.gouguoyin.cn
// @Email contact@gouguoyin.cn

// Package carbon is a simple, semantic and developer-friendly golang package for datetime.
package carbon

import (
	"time"
)

// timezones constant
// 时区常量
const (
	Local = "Local"
	CET   = "CET"
	EET   = "EET"
	EST   = "EST"
	GMT   = "GMT"
	UTC   = "UTC"
	UCT   = "UCT"
	MST   = "MST"

	Cuba      = "Cuba"      // 古巴
	Egypt     = "Egypt"     // 埃及
	Eire      = "Eire"      // 爱尔兰
	Greenwich = "Greenwich" // 格林尼治
	Iceland   = "Iceland"   // 冰岛
	Iran      = "Iran"      // 伊朗
	Israel    = "Israel"    // 以色列
	Jamaica   = "Jamaica"   // 牙买加
	Japan     = "Japan"     // 日本
	Libya     = "Libya"     // 利比亚
	Poland    = "Poland"    // 波兰
	Portugal  = "Portugal"  // 葡萄牙
	PRC       = "PRC"       // 中国
	Singapore = "Singapore" // 新加坡
	Turkey    = "Turkey"    // 土耳其

	Shanghai   = "Asia/Shanghai"       // 上海
	Chongqing  = "Asia/Chongqing"      // 重庆
	Harbin     = "Asia/Harbin"         // 哈尔滨
	HongKong   = "Asia/Hong_Kong"      // 香港
	Macao      = "Asia/Macao"          // 澳门
	Taipei     = "Asia/Taipei"         // 台北
	Tokyo      = "Asia/Tokyo"          // 东京
	Saigon     = "Asia/Saigon"         // 西贡
	Seoul      = "Asia/Seoul"          // 首尔
	Bangkok    = "Asia/Bangkok"        // 曼谷
	Dubai      = "Asia/Dubai"          // 迪拜
	NewYork    = "America/New_York"    // 纽约
	LosAngeles = "America/Los_Angeles" // 洛杉矶
	Chicago    = "America/Chicago"     // 芝加哥
	Moscow     = "Europe/Moscow"       // 莫斯科
	London     = "Europe/London"       // 伦敦
	Berlin     = "Europe/Berlin"       // 柏林
	Paris      = "Europe/Paris"        // 巴黎
	Rome       = "Europe/Rome"         // 罗马
)

// months constant
// 月份常量
const (
	January   = "January"   // 一月
	February  = "February"  // 二月
	March     = "March"     // 三月
	April     = "April"     // 四月
	May       = "May"       // 五月
	June      = "June"      // 六月
	July      = "July"      // 七月
	August    = "August"    // 八月
	September = "September" // 九月
	October   = "October"   // 十月
	November  = "November"  // 十一月
	December  = "December"  // 十二月
)

// weeks constant
// 星期常量
const (
	Monday    = "Monday"    // 周一
	Tuesday   = "Tuesday"   // 周二
	Wednesday = "Wednesday" // 周三
	Thursday  = "Thursday"  // 周四
	Friday    = "Friday"    // 周五
	Saturday  = "Saturday"  // 周六
	Sunday    = "Sunday"    // 周日
)

// numbers constant
// 数字常量
const (
	YearsPerMillennium         = 1000    // 每千年1000年
	YearsPerCentury            = 100     // 每世纪100年
	YearsPerDecade             = 10      // 每十年10年
	QuartersPerYear            = 4       // 每年4季度
	MonthsPerYear              = 12      // 每年12月
	MonthsPerQuarter           = 3       // 每季度3月
	WeeksPerNormalYear         = 52      // 每常规年52周
	weeksPerLongYear           = 53      // 每长年53周
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

// formats constant
// 时间格式化常量
const (
	AnsicFormat         = time.ANSIC
	UnixDateFormat      = time.UnixDate
	RubyDateFormat      = time.RubyDate
	RFC822Format        = time.RFC822
	RFC822ZFormat       = time.RFC822Z
	RFC850Format        = time.RFC850
	RFC1123Format       = time.RFC1123
	RFC1123ZFormat      = time.RFC1123Z
	RssFormat           = time.RFC1123Z
	RFC2822Format       = time.RFC1123Z
	RFC3339Format       = time.RFC3339
	KitchenFormat       = time.Kitchen
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

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	Time         time.Time
	weekStartsAt time.Weekday
	loc          *time.Location
	lang         *Language
	Error        error
}

// NewCarbon returns a new Carbon instance.
// 初始化 Carbon 结构体
func NewCarbon() Carbon {
	return Carbon{weekStartsAt: time.Sunday, loc: time.Local, lang: NewLanguage()}
}

// Time2Carbon converts time.Time to Carbon.
// 将 time.Time 转换成 Carbon
func Time2Carbon(tt time.Time) Carbon {
	c := NewCarbon()
	c.Time = tt
	return c
}

// Carbon2Time converts Carbon to time.Time.
// 将 Carbon 转换成 time.Time
func (c Carbon) Carbon2Time() time.Time {
	return c.Time.In(c.loc)
}

// Now returns a Carbon instance for now.
// 当前
func (c Carbon) Now(timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	c.Time = time.Now().In(c.loc)
	return c
}

// Now returns a Carbon instance for now.
// 当前
func Now(timezone ...string) Carbon {
	return NewCarbon().Now(timezone...)
}

// Tomorrow returns a Carbon instance for tomorrow.
// 明天
func (c Carbon) Tomorrow(timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	if c.IsZero() {
		c.Time = time.Now().In(c.loc).AddDate(0, 0, 1)
	} else {
		c.Time = c.Time.In(c.loc).AddDate(0, 0, 1)
	}
	return c
}

// Tomorrow returns a Carbon instance for tomorrow.
// 明天
func Tomorrow(timezone ...string) Carbon {
	return NewCarbon().Tomorrow(timezone...)
}

// Yesterday returns a Carbon instance for yesterday.
// 昨天
func (c Carbon) Yesterday(timezone ...string) Carbon {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.Error != nil {
		return c
	}
	if c.IsZero() {
		c.Time = time.Now().In(c.loc).AddDate(0, 0, -1)
	} else {
		c.Time = c.Time.In(c.loc).AddDate(0, 0, -1)
	}
	return c
}

// Yesterday returns a Carbon instance for yesterday.
// 昨天
func Yesterday(timezone ...string) Carbon {
	return NewCarbon().Yesterday(timezone...)
}
