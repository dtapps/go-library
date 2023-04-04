package caiyunapp

// GetDressingDesc 穿衣指数 https://docs.caiyunapp.com/docs/tables/lifeindex
func GetDressingDesc(dressing float64) string {
	if dressing <= 0 {
		return "极热"
	} else if dressing <= 1 {
		return "极热"
	} else if dressing <= 2 {
		return "很热"
	} else if dressing <= 3 {
		return "热"
	} else if dressing <= 4 {
		return "温暖"
	} else if dressing <= 5 {
		return "凉爽"
	} else if dressing <= 6 {
		return "冷"
	} else if dressing <= 7 {
		return "寒冷"
	} else if dressing <= 8 {
		return "极冷"
	}
	return ""
}
