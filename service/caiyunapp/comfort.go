package caiyunapp

// GetComfortDesc 舒适度指数 https://docs.caiyunapp.com/docs/tables/lifeindex
func GetComfortDesc(comfort float64) string {
	if comfort <= 0 {
		return "闷热"
	} else if comfort <= 1 {
		return "酷热"
	} else if comfort <= 2 {
		return "很热"
	} else if comfort <= 3 {
		return "热"
	} else if comfort <= 4 {
		return "温暖"
	} else if comfort <= 5 {
		return "舒适"
	} else if comfort <= 6 {
		return "凉爽"
	} else if comfort <= 7 {
		return "冷"
	} else if comfort <= 8 {
		return "很冷"
	} else if comfort <= 9 {
		return "寒冷"
	} else if comfort <= 10 {
		return "极冷"
	} else if comfort <= 11 {
		return "刺骨的冷"
	} else if comfort <= 12 {
		return "湿冷"
	} else if comfort <= 13 {
		return "干冷"
	}
	return ""
}
