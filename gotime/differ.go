package gotime

import "time"

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
