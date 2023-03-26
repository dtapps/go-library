package rocron

import "fmt"

var (
	EarlyInTheMorningSpec    = "0 0 0 * * *"
	EarlyInTheMorningExplain = "每天0点执行一次"
)

// GetSecondsSpec 秒
func GetSecondsSpec(seconds int) string {
	if seconds <= 0 {
		seconds = 0
	}
	if seconds > 60 {
		seconds = 60
	}
	return fmt.Sprintf("*/%d * * * * *", seconds)
}

// GetSecondsExplain 秒
func GetSecondsExplain(seconds int) string {
	if seconds <= 0 {
		seconds = 0
	}
	if seconds > 60 {
		seconds = 60
	}
	return fmt.Sprintf("每隔%d秒执行一次", seconds)
}

// GetMinutesSpec 分
func GetMinutesSpec(minutes, seconds int) string {
	if minutes <= 0 {
		minutes = 0
	}
	if seconds > 59 {
		seconds = 59
	}
	return fmt.Sprintf("%d */%d * * * *", seconds, minutes)
}

// GetMinutesExplain 分
func GetMinutesExplain(minutes, seconds int) string {
	if minutes <= 0 {
		minutes = 0
	}
	if seconds <= 0 {
		return fmt.Sprintf("每隔%d分钟执行一次", minutes)
	}
	if seconds > 59 {
		seconds = 59
	}
	return fmt.Sprintf("每隔%d分钟%d秒执行一次", minutes, seconds)
}

// GetHoursSpec 小时
func GetHoursSpec(hours, minutes, seconds int) string {
	if hours <= 0 {
		hours = 0
	}
	if minutes > 59 {
		minutes = 59
	}
	if seconds > 59 {
		seconds = 59
	}
	return fmt.Sprintf("%d %d */%d * * *", seconds, minutes, hours)
}

// GetHoursExplain 小时
func GetHoursExplain(hours, minutes, seconds int) string {
	if hours <= 0 {
		hours = 0
	}
	if seconds <= 0 {
		if minutes <= 0 {
			return fmt.Sprintf("每隔%d小时执行一次", hours)
		}
		if minutes > 59 {
			minutes = 59
		}
		return fmt.Sprintf("每隔%d小时%d分执行一次", hours, minutes)
	}
	if seconds > 59 {
		seconds = 59
	}
	return fmt.Sprintf("每隔%d小时%d分%d秒执行一次", hours, minutes, seconds)
}
