package gotime

import "time"

// Gt 是否大于
func (p Pro) Gt(t2 time.Time) bool {
	return p.Time.After(t2)
}

// Lt 是否小于
func (p Pro) Lt(t2 time.Time) bool {
	return p.Time.Before(t2)
}

// Eq 是否等于
func (p Pro) Eq(t2 time.Time) bool {
	return p.Time.Equal(t2)
}

// Ne 是否不等于
func (p Pro) Ne(t2 time.Time) bool {
	return !p.Eq(t2)
}

// Gte 是否大于等于
func (p Pro) Gte(t2 time.Time) bool {
	return p.Gt(t2) || p.Eq(t2)
}

// Lte 是否小于等于
func (p Pro) Lte(t2 time.Time) bool {
	return p.Lt(t2) || p.Eq(t2)
}

// Between 是否在两个时间之间(不包括这两个时间)
func (p Pro) Between(start time.Time, end time.Time) bool {
	if p.Gt(start) && p.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart 是否在两个时间之间(包括开始时间)
func (p Pro) BetweenIncludedStart(start time.Time, end time.Time) bool {
	if p.Gte(start) && p.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd 是否在两个时间之间(包括结束时间)
func (p Pro) BetweenIncludedEnd(start time.Time, end time.Time) bool {
	if p.Gt(start) && p.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth 是否在两个时间之间(包括这两个时间)
func (p Pro) BetweenIncludedBoth(start time.Time, end time.Time) bool {
	if p.Gte(start) && p.Lte(end) {
		return true
	}
	return false
}
