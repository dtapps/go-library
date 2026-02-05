package gotime

import (
	"log/slog"
	"strconv"
	"strings"
	"time"
)

// Current 获取当前的时间
func Current() Pro {
	p := NewPro()
	p.loc = shangHaiLoc
	p.Time = time.Now().In(p.loc)
	return p
}

// SetCurrent 设置当前的时间
func SetCurrent(sTime time.Time) Pro {
	p := NewPro()
	p.loc = shangHaiLoc
	p.Time = sTime.In(p.loc)
	return p
}

// SetCurrentParse 设置当前的时间
func SetCurrentParse(str string) Pro {
	p := NewPro()
	p.loc = shangHaiLoc

	layout := DateTimeFormat
	if str == "" || str == "0" || str == "0000-00-00 00:00:00" || str == "0000-00-00" || str == "00:00:00" {
		return p
	}
	if len(str) == 10 && strings.Count(str, "-") == 2 {
		layout = DateFormat
	}
	if strings.Index(str, "T") == 10 {
		layout = RFC3339Format
	}
	if _, err := strconv.ParseInt(str, 10, 64); err == nil {
		switch len(str) {
		case 8:
			layout = FormatYearMonthDay
		case 14:
			layout = FormatYearMonthDayHourMinuteSeconds
		}
	}
	t, err := time.ParseInLocation(layout, str, p.loc)
	if err != nil {
		slog.Warn("时间解析失败",
			slog.String("input", str),
			slog.String("layout", layout),
			slog.Any("err", err),
		)
	}

	p.Time = t
	return p
}

// SetCurrentUnix 设置当前的时间 Unix时间戳
func SetCurrentUnix(ts int64) Pro {
	p := NewPro()
	p.loc = shangHaiLoc
	p.Time = time.Unix(ts, 0).In(p.loc)
	return p
}

// SetCurrentMillisecondUnix 设置当前的时间 毫秒Unix时间戳
func SetCurrentMillisecondUnix(ts int64) Pro {
	p := NewPro()
	p.loc = shangHaiLoc
	sec := ts / 1000
	nsec := (ts % 1000) * int64(time.Millisecond) // 保留毫秒
	p.Time = time.Unix(sec, nsec).In(p.loc)
	return p
}

// Now 今天此刻
func (p Pro) Now() time.Time {
	return p.Time
}

// Format 今天此刻格式化
func (p Pro) Format() string {
	return p.Time.Format(DateTimeFormat)
}

// FormatFilter 今天此刻格式化 带 过滤无效时间
func (p Pro) FormatFilter() string {
	if strings.Contains(p.Time.Format(DateTimeFormat), "0001-01-01") {
		return ""
	} else {
		return p.Time.Format(DateTimeFormat)
	}
}

// ToDateFormat 今天此刻日期
func (p Pro) ToDateFormat() string {
	return p.Time.Format(DateFormat)
}

// ToDateFormatTime 今天此刻日期
func (p Pro) ToDateFormatTime() time.Time {
	return SetCurrentParse(p.Time.Format(DateFormat)).Time
}

// ToTimeFormat 今天此刻时间
func (p Pro) ToTimeFormat() string {
	return p.Time.Format(TimeFormat)
}

// Timestamp 今天此刻时间戳
func (p Pro) Timestamp() int64 {
	return p.Time.Unix()
}

// TimestampWithSecond 今天此刻时间戳
func (p Pro) TimestampWithSecond() int64 {
	return p.Time.Unix()
}

// TimestampWithMillisecond 今天毫秒级时间戳
func (p Pro) TimestampWithMillisecond() int64 {
	return p.Time.UnixNano() / int64(time.Millisecond)
}

// TimestampWithMicrosecond 今天微秒级时间戳
func (p Pro) TimestampWithMicrosecond() int64 {
	return p.Time.UnixNano() / int64(time.Microsecond)
}

// TimestampWithNanosecond 今天纳秒级时间戳
func (p Pro) TimestampWithNanosecond() int64 {
	return p.Time.UnixNano()
}
