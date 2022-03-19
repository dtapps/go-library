package carbon

import (
	"time"
)

// IsZero whether is zero time.
// 是否是零值时间
func (c Carbon) IsZero() bool {
	return c.Time.IsZero()
}

// IsInvalid whether is invalid time.
// 是否是无效时间
func (c Carbon) IsInvalid() bool {
	if c.Error != nil || c.IsZero() {
		return true
	}
	return false
}

// IsNow whether is now time.
// 是否是当前时间
func (c Carbon) IsNow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() == c.Now().Timestamp()
}

// IsFuture whether is future time.
// 是否是未来时间
func (c Carbon) IsFuture() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() > c.Now().Timestamp()
}

// IsPast whether is past time.
// 是否是过去时间
func (c Carbon) IsPast() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Timestamp() < c.Now().Timestamp()
}

// IsLeapYear whether is a leap year.
// 是否是闰年
func (c Carbon) IsLeapYear() bool {
	if c.IsInvalid() {
		return false
	}
	year := c.Time.In(c.loc).Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsLongYear whether is a long year, see https://en.wikipedia.org/wiki/ISO_8601#Week_dates.
// 是否是长年
func (c Carbon) IsLongYear() bool {
	if c.IsInvalid() {
		return false
	}
	_, w := time.Date(c.Year(), time.December, 31, 0, 0, 0, 0, c.loc).ISOWeek()
	return w == weeksPerLongYear
}

// IsJanuary whether is January.
// 是否是一月
func (c Carbon) IsJanuary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.January
}

// IsFebruary whether is February.
// 是否是二月
func (c Carbon) IsFebruary() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.February
}

// IsMarch whether is March.
// 是否是三月
func (c Carbon) IsMarch() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.March
}

// IsApril whether is April.
// 是否是四月
func (c Carbon) IsApril() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.April
}

// IsMay whether is May.
// 是否是五月
func (c Carbon) IsMay() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.May
}

// IsJune whether is June.
// 是否是六月
func (c Carbon) IsJune() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.June
}

// IsJuly whether is July.
// 是否是七月
func (c Carbon) IsJuly() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.July
}

// IsAugust whether is August.
// 是否是八月
func (c Carbon) IsAugust() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.August
}

// IsSeptember whether is September.
// 是否是九月
func (c Carbon) IsSeptember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.September
}

// IsOctober whether is October.
// 是否是十月
func (c Carbon) IsOctober() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.October
}

// IsNovember whether is November.
// 是否是十一月
func (c Carbon) IsNovember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.November
}

// IsDecember whether is December.
// 是否是十二月
func (c Carbon) IsDecember() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Month() == time.December
}

// IsMonday whether is Monday.
// 是否是周一
func (c Carbon) IsMonday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Monday
}

// IsTuesday whether is Tuesday.
// 是否是周二
func (c Carbon) IsTuesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Tuesday
}

// IsWednesday whether is Wednesday.
// 是否是周三
func (c Carbon) IsWednesday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Wednesday
}

// IsThursday whether is Thursday.
// 是否是周四
func (c Carbon) IsThursday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Thursday
}

// IsFriday whether is Friday.
// 是否是周五
func (c Carbon) IsFriday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Friday
}

// IsSaturday whether is Saturday.
// 是否是周六
func (c Carbon) IsSaturday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Saturday
}

// IsSunday whether is Sunday.
// 是否是周日
func (c Carbon) IsSunday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.Time.In(c.loc).Weekday() == time.Sunday
}

// IsWeekday whether is weekday.
// 是否是工作日
func (c Carbon) IsWeekday() bool {
	if c.IsInvalid() {
		return false
	}
	return !c.IsSaturday() && !c.IsSunday()
}

// IsWeekend whether is weekend.
// 是否是周末
func (c Carbon) IsWeekend() bool {
	if c.IsInvalid() {
		return false
	}
	return c.IsSaturday() || c.IsSunday()
}

// IsYesterday whether is yesterday.
// 是否是昨天
func (c Carbon) IsYesterday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().SubDay().ToDateString()
}

// IsToday whether is today.
// 是否是今天
func (c Carbon) IsToday() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == c.Now().ToDateString()
}

// IsTomorrow whether is tomorrow.
// 是否是明天
func (c Carbon) IsTomorrow() bool {
	if c.IsInvalid() {
		return false
	}
	return c.ToDateString() == Now().AddDay().ToDateString()
}

// Compare comparison by a operator.
// 时间比较
func (c Carbon) Compare(operator string, t Carbon) bool {
	switch operator {
	case "=":
		return c.Eq(t)
	case "<>", "!=":
		return !c.Eq(t)
	case ">":
		return c.Gt(t)
	case ">=":
		return c.Gte(t)
	case "<":
		return c.Lt(t)
	case "<=":
		return c.Lte(t)
	}
	return false
}

// Gt whether greater than.
// 是否大于
func (c Carbon) Gt(t Carbon) bool {
	return c.Time.After(t.Time)
}

// Lt whether less than.
// 是否小于
func (c Carbon) Lt(t Carbon) bool {
	return c.Time.Before(t.Time)
}

// Eq whether equal.
// 是否等于
func (c Carbon) Eq(t Carbon) bool {
	return c.Time.Equal(t.Time)
}

// Ne whether not equal.
// 是否不等于
func (c Carbon) Ne(t Carbon) bool {
	return !c.Eq(t)
}

// Gte whether greater than or equal.
// 是否大于等于
func (c Carbon) Gte(t Carbon) bool {
	return c.Gt(t) || c.Eq(t)
}

// Lte whether less than or equal.
// 是否小于等于
func (c Carbon) Lte(t Carbon) bool {
	return c.Lt(t) || c.Eq(t)
}

// Between whether between two times, excluded the start and end time.
// 是否在两个时间之间(不包括这两个时间)
func (c Carbon) Between(start Carbon, end Carbon) bool {
	if c.Gt(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedStart whether between two times, included the start time.
// 是否在两个时间之间(包括开始时间)
func (c Carbon) BetweenIncludedStart(start Carbon, end Carbon) bool {
	if c.Gte(start) && c.Lt(end) {
		return true
	}
	return false
}

// BetweenIncludedEnd whether between two times, included the end time.
// 是否在两个时间之间(包括结束时间)
func (c Carbon) BetweenIncludedEnd(start Carbon, end Carbon) bool {
	if c.Gt(start) && c.Lte(end) {
		return true
	}
	return false
}

// BetweenIncludedBoth whether between two times, included the start and end time.
// 是否在两个时间之间(包括这两个时间)
func (c Carbon) BetweenIncludedBoth(start Carbon, end Carbon) bool {
	if c.Gte(start) && c.Lte(end) {
		return true
	}
	return false
}
