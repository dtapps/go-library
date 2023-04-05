package caiyunapp

// GetRealtimeUltravioletDesc 紫外线 https://docs.caiyunapp.com/docs/tables/lifeindex
func GetRealtimeUltravioletDesc(ultraviolet float64) string {
	if ultraviolet <= 0 {
		return "无"
	} else if ultraviolet <= 1 {
		return "很弱"
	} else if ultraviolet <= 2 {
		return "很弱"
	} else if ultraviolet <= 3 {
		return "弱"
	} else if ultraviolet <= 4 {
		return "弱"
	} else if ultraviolet <= 5 {
		return "中等"
	} else if ultraviolet <= 6 {
		return "中等"
	} else if ultraviolet <= 7 {
		return "强"
	} else if ultraviolet <= 8 {
		return "强"
	} else if ultraviolet <= 9 {
		return "强"
	} else if ultraviolet <= 10 {
		return "很强"
	} else if ultraviolet <= 11 {
		return "极强"
	}
	return "无"
}

// GetDailyUltravioletDesc 紫外线 https://docs.caiyunapp.com/docs/tables/lifeindex
func GetDailyUltravioletDesc(ultraviolet string) string {
	if ultraviolet <= "1" {
		return "最弱"
	} else if ultraviolet <= "2" {
		return "弱"
	} else if ultraviolet <= "3" {
		return "中等"
	} else if ultraviolet <= "4" {
		return "强"
	} else if ultraviolet <= "5" {
		return "很强"
	}
	return "无"
}
