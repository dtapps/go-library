package rocron

func filterSeconds(seconds int) int {
	if seconds <= 0 {
		return 0
	}
	if seconds > 60 {
		return 60
	}
	return seconds
}

func filterMinutes(minutes int) int {
	if minutes <= 0 {
		return 0
	}
	if minutes > 59 {
		return 59
	}
	return minutes
}

func filterHours(hours int) int {
	if hours <= 0 {
		return 0
	}
	if hours > 23 {
		return 23
	}
	return hours
}

func filterDays(days int) int {
	if days <= 0 {
		return 1
	}
	if days > 31 {
		return 31
	}
	return days
}
