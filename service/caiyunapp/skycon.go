package caiyunapp

// GetSkyCon https://docs.caiyunapp.com/docs/tables/skycon
func GetSkyCon(skycon string) string {
	switch skycon {
	case "CLEAR_DAY":
		return "晴"
		//return "晴（白天）"
	case "CLEAR_NIGHT":
		return "晴"
		//return "晴（夜间）"
	case "PARTLY_CLOUDY_DAY":
		return "多云"
		//return "多云（白天）"
	case "PARTLY_CLOUDY_NIGHT":
		return "多云"
		//return "多云（夜间）"
	case "CLOUDY":
		return "阴"
	case "LIGHT_HAZE":
		return "轻度雾霾"
	case "MODERATE_HAZE":
		return "中度雾霾"
	case "HEAVY_HAZE":
		return "重度雾霾"
	case "LIGHT_RAIN":
		return "小雨"
	case "MODERATE_RAIN":
		return "中雨"
	case "HEAVY_RAIN":
		return "大雨"
	case "STORM_RAIN":
		return "暴雨"
	case "FOG":
		return "雾"
	case "LIGHT_SNOW":
		return "小雪"
	case "MODERATE_SNOW":
		return "中雪"
	case "HEAVY_SNOW":
		return "大雪"
	case "STORM_SNOW":
		return "暴雪"
	case "DUST":
		return "浮尘"
	case "SAND":
		return "沙尘"
	case "WIND":
		return "大风"
	}
	return skycon
}
