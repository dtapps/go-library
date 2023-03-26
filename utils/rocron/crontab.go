package rocron

import (
	"fmt"
)

var (
	YearlySpec      = "0 0 0 1 1 *"
	YearlyExplain   = "每年一次，1 月 1 日午夜执行一次"
	AnnuallySpec    = "0 0 0 1 1 *"
	AnnuallyExplain = "每年一次，1 月 1 日午夜执行一次"
	MonthlySpec     = "0 0 0 1 * *"
	MonthlyExplain  = "每月执行，午夜，月初执行一次"
	WeeklySpec      = "0 0 0 * * 0"
	WeeklyExplain   = "每周执行，周六和周日之间的午夜执行一次"
	DailySpec       = "0 0 0 * * *"
	DailyExplain    = "每天午夜执行一次"
	MidnightSpec    = "0 0 0 * * *"
	MidnightExplain = "每天午夜执行一次"
	HourlySpec      = "0 0 * * * *"
	HourlyExplain   = "每小时运行，每小时开始执行一次"
)

/*
*
┌─────────────second 范围 (0 - 60)
│ ┌───────────── min (0 - 59)
│ │ ┌────────────── hour (0 - 23)
│ │ │ ┌─────────────── day of month (1 - 31)
│ │ │ │ ┌──────────────── month (1 - 12)
│ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
│ │ │ │ │ │                  Saturday)
│ │ │ │ │ │
│ │ │ │ │ │
* * * * * *
*/
const (
	secondsSpec = "*/%d * * * * *"
	minutesSpec = "%d */%d * * * *"
	hoursSpec   = "%d %d */%d * * *"
	daySpec     = "%d %d %d */%d * *"
	daysSpec    = "%d %d %d %s * *"
)

// GetSecondSpec 每隔秒
func GetSecondSpec(seconds int) string {
	seconds = filterSeconds(seconds)
	return fmt.Sprintf(secondsSpec, seconds)
}

// GetSecondExplain 每隔秒
func GetSecondExplain(seconds int) string {
	seconds = filterSeconds(seconds)
	return fmt.Sprintf("每隔%d秒执行一次", seconds)
}

// GetMinuteSpec 每隔分钟
func GetMinuteSpec(minutes, seconds int) string {
	minutes = filterMinutes(minutes)
	seconds = filterSeconds(seconds)
	return fmt.Sprintf(minutesSpec, seconds, minutes)
}

// GetMinuteExplain 每隔分钟
func GetMinuteExplain(minutes, seconds int) string {
	minutes = filterMinutes(minutes)
	seconds = filterSeconds(seconds)
	if seconds <= 0 {
		return fmt.Sprintf("每隔%d分钟执行一次", minutes)
	}
	return fmt.Sprintf("每隔%d分钟%d秒执行一次", minutes, seconds)
}

// GetHourSpec 每隔小时
func GetHourSpec(hours, minutes, seconds int) string {
	hours = filterHours(hours)
	minutes = filterMinutes(minutes)
	seconds = filterSeconds(seconds)
	return fmt.Sprintf(hoursSpec, seconds, minutes, hours)
}

// GetHourExplain 每隔小时
func GetHourExplain(hours, minutes, seconds int) string {
	hours = filterHours(hours)
	minutes = filterMinutes(minutes)
	seconds = filterSeconds(seconds)
	if seconds <= 0 {
		if minutes <= 0 {
			return fmt.Sprintf("每隔%d小时执行一次", hours)
		}
		return fmt.Sprintf("每隔%d小时%d分执行一次", hours, minutes)
	}
	return fmt.Sprintf("每隔%d小时%d分%d秒执行一次", hours, minutes, seconds)
}

// GetDaySpec 天
func GetDaySpec(day, hours, minutes, seconds int) string {
	day = filterDays(day)
	hours = filterHours(hours)
	minutes = filterMinutes(minutes)
	seconds = filterSeconds(seconds)
	return fmt.Sprintf(daySpec, seconds, minutes, hours, day)
	//days := concatenateStrings(day)
	//return fmt.Sprintf(daysSpec, seconds, minutes, hours, days)
}

// GetDayExplain 天
func GetDayExplain(days, hours, minutes, seconds int) string {
	days = filterDays(days)
	hours = filterHours(hours)
	minutes = filterMinutes(minutes)
	seconds = filterSeconds(seconds)
	if seconds <= 0 {
		if minutes <= 0 {
			if hours <= 0 {
				return fmt.Sprintf("每隔%d天执行一次", days)
			}
			return fmt.Sprintf("每隔%d天%d时执行一次", days, hours)
		}
		return fmt.Sprintf("每隔%d天%d时%d分执行一次", days, hours, minutes)
	}
	return fmt.Sprintf("每隔%d天%d时%d分%d秒执行一次", days, hours, minutes, seconds)
}
