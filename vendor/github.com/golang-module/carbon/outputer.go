package carbon

import (
	"bytes"
	"strconv"
	"strings"
)

// ToTimestamp outputs a timestamp with second(will be removed from v2.0, only keep Timestamp).
// 输出秒级时间戳(将在v2.0版本移除，只保留 Timestamp)
func (c Carbon) ToTimestamp() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.Timestamp()
}

// ToTimestampWithSecond outputs a timestamp with second(will be removed from v2.0, only keep TimestampWithSecond).
// 输出秒级时间戳(将在v2.0版本移除，只保留 TimestampWithSecond)
func (c Carbon) ToTimestampWithSecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.TimestampWithSecond()
}

// ToTimestampWithMillisecond outputs a timestamp with millisecond(will be removed from v2.0, only keep TimestampWithMillisecond).
// 输出毫秒级时间戳(将在v2.0版本移除，只保留 TimestampWithMillisecond)
func (c Carbon) ToTimestampWithMillisecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.TimestampWithMillisecond()
}

// ToTimestampWithMicrosecond outputs a timestamp with microsecond(will be removed from v2.0, only keep TimestampWithMicrosecond).
// 输出微秒级时间戳(将在v2.0版本移除，只保留 TimestampWithMicrosecond)
func (c Carbon) ToTimestampWithMicrosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.TimestampWithMicrosecond()
}

// ToTimestampWithNanosecond outputs a timestamp with nanosecond(will be removed from v2.0, only keep TimestampWithNanosecond).
// 输出纳秒级时间戳(将在v2.0版本移除，只保留 TimestampWithNanosecond)
func (c Carbon) ToTimestampWithNanosecond() int64 {
	if c.IsInvalid() {
		return 0
	}
	return c.TimestampWithNanosecond()
}

// String outputs a string in date and time format, implement Stringer interface.
// 实现 Stringer 接口
func (c Carbon) String() string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToDateTimeString()
}

// ToString outputs a string in "2006-01-02 15:04:05.999999999 -0700 MST" format.
// 输出"2006-01-02 15:04:05.999999999 -0700 MST"格式字符串
func (c Carbon) ToString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).String()
}

// ToMonthString outputs a string in month format, i18n is supported.
// 输出完整月份字符串，支持i18n
func (c Carbon) ToMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["months"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 12 {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToShortMonthString outputs a string in short month format, i18n is supported.
// 输出缩写月份字符串，支持i18n
func (c Carbon) ToShortMonthString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["short_months"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 12 {
			return slice[c.Month()-1]
		}
	}
	return ""
}

// ToWeekString outputs a string in week format, i18n is supported.
// 输出完整星期字符串，支持i18n
func (c Carbon) ToWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["weeks"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 7 {
			return slice[c.Week()]
		}
	}
	return ""
}

// ToShortWeekString outputs a string in short week format, i18n is supported.
// 输出缩写星期字符串，支持i18n
func (c Carbon) ToShortWeekString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	if months, ok := c.lang.resources["short_weeks"]; ok {
		slice := strings.Split(months, "|")
		if len(slice) == 7 {
			return slice[c.Week()]
		}
	}
	return ""
}

// ToDayDateTimeString outputs a string in day, date and time format.
// 输出天数日期时间字符串
func (c Carbon) ToDayDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(DayDateTimeFormat)
}

// ToDateTimeString outputs a string in date and time format.
// 输出日期时间字符串
func (c Carbon) ToDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(DateTimeFormat)
}

// ToShortDateTimeString outputs a string in short date and time format.
// 输出简写日期时间字符串
func (c Carbon) ToShortDateTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(ShortDateTimeFormat)
}

// ToDateString outputs a string in date format.
// 输出日期字符串
func (c Carbon) ToDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(DateFormat)
}

// ToShortDateString outputs a string in short date format.
// 输出简写日期字符串
func (c Carbon) ToShortDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(ShortDateFormat)
}

// ToTimeString outputs a string in time format.
// 输出时间字符串
func (c Carbon) ToTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(TimeFormat)
}

// ToShortTimeString outputs a string in short time format.
// 输出简写时间字符串
func (c Carbon) ToShortTimeString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(ShortTimeFormat)
}

// ToAtomString outputs a string in ATOM format.
// 输出 ATOM 格式字符串
func (c Carbon) ToAtomString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToRfc3339String(timezone...)
}

// ToAnsicString outputs a string in ANSIC format.
// 输出 ANSIC 格式字符串
func (c Carbon) ToAnsicString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(AnsicFormat)
}

// ToCookieString outputs a string in COOKIE format.
// 输出 COOKIE 格式字符串
func (c Carbon) ToCookieString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(CookieFormat)
}

// ToRssString outputs a string in RSS format.
// 输出 RSS 格式字符串
func (c Carbon) ToRssString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RssFormat)
}

// ToW3cString outputs a string in W3C format.
// 输出 W3C 格式字符串
func (c Carbon) ToW3cString(timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToRfc3339String(timezone...)
}

// ToUnixDateString outputs a string in unix date format.
// 输出 UnixDate 格式字符串
func (c Carbon) ToUnixDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(UnixDateFormat)
}

// ToRubyDateString outputs a string in ruby date format.
// 输出 RubyDate 格式字符串
func (c Carbon) ToRubyDateString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RubyDateFormat)
}

// ToKitchenString outputs a string in KITCHEN format.
// 输出 KITCHEN 格式字符串
func (c Carbon) ToKitchenString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(KitchenFormat)
}

// ToIso8601String outputs a string in ISO8601 format.
// 输出 ISO8601 格式字符串
func (c Carbon) ToIso8601String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(Iso8601Format)
}

// ToRfc822String outputs a string in RFC822 format.
// 输出 RFC822 格式字符串
func (c Carbon) ToRfc822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC822Format)
}

// ToRfc822zString outputs a string in RFC822Z format.
// 输出 RFC822Z 格式字符串
func (c Carbon) ToRfc822zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC822ZFormat)
}

// ToRfc850String outputs a string in RFC850 format.
// 输出 RFC850 格式字符串
func (c Carbon) ToRfc850String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC850Format)
}

// ToRfc1036String outputs a string in RFC1036 format.
// 输出 RFC1036 格式字符串
func (c Carbon) ToRfc1036String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC1036Format)
}

// ToRfc1123String outputs a string in RFC1123 format.
// 输出 RFC1123 格式字符串
func (c Carbon) ToRfc1123String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC1123Format)
}

// ToRfc1123zString outputs a string in RFC1123z format.
// 输出 RFC1123z 格式字符串
func (c Carbon) ToRfc1123zString(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC1123ZFormat)
}

// ToRfc2822String outputs a string in RFC2822 format.
// 输出 RFC2822 格式字符串
func (c Carbon) ToRfc2822String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC2822Format)
}

// ToRfc3339String outputs a string in RFC3339 format.
// 输出 RFC3339 格式字符串
func (c Carbon) ToRfc3339String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC3339Format)
}

// ToRfc7231String outputs a string in RFC7231 format.
// 输出 RFC7231 格式字符串
func (c Carbon) ToRfc7231String(timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(RFC7231Format)
}

// ToLayoutString outputs a string by layout.
// 输出指定布局的时间字符串
func (c Carbon) ToLayoutString(layout string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	return c.Time.In(c.loc).Format(layout)
}

// Layout outputs a string by layout, it is short for ToLayoutString.
// 输出指定布局的时间字符串, 是 ToLayoutString 的简写
func (c Carbon) Layout(layout string, timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToLayoutString(layout, timezone...)
}

// ToFormatString outputs a string by format.
// 输出指定格式的时间字符串
func (c Carbon) ToFormatString(format string, timezone ...string) string {
	if len(timezone) > 0 {
		c.loc, c.Error = getLocationByTimezone(timezone[len(timezone)-1])
	}
	if c.IsInvalid() {
		return ""
	}
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(format); i++ {
		if layout, ok := formats[byte(format[i])]; ok {
			buffer.WriteString(c.Time.In(c.loc).Format(layout))
		} else {
			switch format[i] {
			case '\\': // 原样输出，不解析
				buffer.WriteByte(format[i+1])
				i++
				continue
			case 'W': // ISO-8601 格式数字表示的年份中的第几周，取值范围 1-52
				buffer.WriteString(strconv.Itoa(c.WeekOfYear()))
			case 'N': // ISO-8601 格式数字表示的星期中的第几天，取值范围 1-7
				buffer.WriteString(strconv.Itoa(c.DayOfWeek()))
			case 'S': // 月份中第几天的英文缩写后缀，如st, nd, rd, th
				suffix := "th"
				switch c.Day() {
				case 1, 21, 31:
					suffix = "st"
				case 2, 22:
					suffix = "nd"
				case 3, 23:
					suffix = "rd"
				}
				buffer.WriteString(suffix)
			case 'L': // 是否为闰年，如果是闰年为 1，否则为 0
				if c.IsLeapYear() {
					buffer.WriteString("1")
				} else {
					buffer.WriteString("0")
				}
			case 'G': // 数字表示的小时，24 小时格式，没有前导零，取值范围 0-23
				buffer.WriteString(strconv.Itoa(c.Hour()))
			case 'U': // 秒级时间戳，如 1611818268
				buffer.WriteString(strconv.FormatInt(c.Timestamp(), 10))
			case 'u': // 数字表示的毫秒，如 999
				buffer.WriteString(strconv.Itoa(c.Millisecond()))
			case 'w': // 数字表示的星期中的第几天，取值范围 0-6
				buffer.WriteString(strconv.Itoa(c.DayOfWeek() - 1))
			case 't': // 指定的月份有几天，取值范围 28-31
				buffer.WriteString(strconv.Itoa(c.DaysInMonth()))
			case 'z': // 年份中的第几天，取值范围 0-365
				buffer.WriteString(strconv.Itoa(c.DayOfYear() - 1))
			case 'e': // 当前位置，如 UTC，GMT，Atlantic/Azores
				buffer.WriteString(c.Location())
			case 'Q': // 当前季度，取值范围 1-4
				buffer.WriteString(strconv.Itoa(c.Quarter()))
			case 'C': // 当前世纪数，取值范围 0-99
				buffer.WriteString(strconv.Itoa(c.Century()))
			default:
				buffer.WriteByte(format[i])
			}
		}
	}
	return buffer.String()
}

// Format outputs a string by format, it is short for ToFormatString.
// 输出指定格式的时间字符串, 是 ToFormatString 的简写
func (c Carbon) Format(format string, timezone ...string) string {
	if c.IsInvalid() {
		return ""
	}
	return c.ToFormatString(format, timezone...)
}
