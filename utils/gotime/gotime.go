package gotime

import (
	"fmt"
	"log"
	"time"
)

const (
	RFC822Format  = "Mon, 02 Jan 2006 15:04:05 MST"
	ISO8601Format = "2006-01-02T15:04:05Z"
)

func NowUTCSeconds() int64 { return time.Now().UTC().Unix() }

func NowUTCNanoSeconds() int64 { return time.Now().UTC().UnixNano() }

// GetCurrentDate 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrentUnix 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// GetCurrentMilliUnix 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetCurrentNanoUnix 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

// GetCurrentWjDate 获取当前的时间 - 字符串 - 没有间隔
func GetCurrentWjDate() string {
	return time.Now().Format("20060102")
}

func FormatISO8601Date(timestampSecond int64) string {
	tm := time.Unix(timestampSecond, 0).UTC()
	return tm.Format(ISO8601Format)
}

// Pro 结构体
type Pro struct {
	Time time.Time
}

// NewPro 初始化结构体
func NewPro() Pro {
	return Pro{
		Time: time.Now(),
	}
}

// Current 获取当前的时间
func Current() Pro {
	p := NewPro()
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// Docker部署golang应用时时区问题 https://www.ddhigh.com/2018/03/01/golang-docker-timezone.html
		log.Printf("时区错误：%v\n", err)
		p.Time = time.Now().Add(time.Hour * 8)
	} else {
		p.Time = time.Now().In(location)
	}
	return p
}

// SetCurrent 设置当前的时间
func SetCurrent(sTime time.Time) Pro {
	p := NewPro()
	p.Time = sTime
	return p
}

// SetCurrentParse 设置当前的时间
func SetCurrentParse(str string) Pro {
	p := NewPro()
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	p.Time = location
	return p
}

// StartOfDay 本日开始时间
func (p Pro) StartOfDay() Pro {
	p.Time = time.Date(p.Time.Year(), p.Time.Month(), p.Time.Day(), 0, 0, 0, 0, p.Time.Location())
	return p
}

// EndOfDay 本日结束时间
func (p Pro) EndOfDay() Pro {
	p.Time = time.Date(p.Time.Year(), p.Time.Month(), p.Time.Day(), 23, 59, 59, 0, p.Time.Location())
	return p
}

// BeforeSeconds 获取n秒前的时间
func (p Pro) BeforeSeconds(seconds int) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("-%ds", seconds))
	p.Time = p.Time.Add(st)
	return p
}

// AfterSeconds 获取n秒后的时间
func (p Pro) AfterSeconds(seconds int) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("+%ds", seconds))
	p.Time = p.Time.Add(st)
	return p
}

// BeforeHour 获取n小时前的时间
func (p Pro) BeforeHour(hour int) Pro {
	st, _ := time.ParseDuration(fmt.Sprintf("-%dh", hour))
	p.Time = p.Time.Add(st)
	return p
}

// AfterHour 获取n小时后的时间
func (p Pro) AfterHour(hour int) Pro {
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

// Now 返回当前时间
func (p Pro) Now() time.Time {
	return p.Time
}

// Format 格式化
func (p Pro) Format() string {
	return p.Time.Format("2006-01-02 15:04:05")
}

// SetFormat 格式化
func (p Pro) SetFormat(layout string) string {
	return p.Time.Format(layout)
}

// Timestamp 秒级时间戳
func (p Pro) Timestamp() int64 {
	return p.Time.Unix()
}

// Unix 时间戳
func (p Pro) Unix() int64 {
	return p.Time.Unix()
}

// UnixNano 纳秒级时间戳
func (p Pro) UnixNano() int64 {
	return p.Time.UnixNano()
}

// TimestampWithMillisecond 毫秒级时间戳
func (p Pro) TimestampWithMillisecond() int64 {
	return p.Time.UnixNano() / 100
}

// Nanosecond 时间戳小数部分 单位：纳秒
func (p Pro) Nanosecond() int {
	return p.Time.Nanosecond()
}

// DiffInHour 相差多少小时
func (p Pro) DiffInHour(t2 time.Time) (hour int64) {
	t2.Before(p.Time)
	diff := p.Time.Unix() - t2.Unix()
	hour = diff / 3600
	return hour
}

// DiffInHourWithAbs 相差多少小时(绝对值)
func (p Pro) DiffInHourWithAbs(t2 time.Time) (hour int64) {
	p.Time.Before(t2)
	diff := t2.Unix() - p.Time.Unix()
	hour = diff / 3600
	if hour > 0 {
		return hour
	}
	t2.Before(p.Time)
	diff = p.Time.Unix() - t2.Unix()
	hour = diff / 3600
	return hour
}

// DiffInMinutes 相差多少分钟
func (p Pro) DiffInMinutes(t2 time.Time) (hour int64) {
	t2.Before(p.Time)
	diff := p.Time.Unix() - t2.Unix()
	hour = diff / 60
	return hour
}

// DiffInMinutesWithAbs 相差多少分钟(绝对值)
func (p Pro) DiffInMinutesWithAbs(t2 time.Time) (hour int64) {
	p.Time.Before(t2)
	diff := t2.Unix() - p.Time.Unix()
	hour = diff / 60
	if hour > 0 {
		return hour
	}
	t2.Before(p.Time)
	diff = p.Time.Unix() - t2.Unix()
	hour = diff / 60
	return hour
}

// DiffInSecond 相差多少秒
func (p Pro) DiffInSecond(t2 time.Time) (hour int64) {
	t2.Before(p.Time)
	diff := p.Time.Unix() - t2.Unix()
	hour = diff
	return hour
}

// DiffInSecondWithAbs 相差多少秒(绝对值)
func (p Pro) DiffInSecondWithAbs(t2 time.Time) (hour int64) {
	p.Time.Before(t2)
	diff := t2.Unix() - p.Time.Unix()
	hour = diff
	if hour > 0 {
		return hour
	}
	t2.Before(p.Time)
	diff = p.Time.Unix() - t2.Unix()
	hour = diff
	return hour
}
