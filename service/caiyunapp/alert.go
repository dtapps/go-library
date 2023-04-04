package caiyunapp

func GetCodeType(code string) string {
	firsTwo := code[:2]
	switch firsTwo {
	case "01":
		return "台风"
	case "02":
		return "暴雨"
	case "03":
		return "暴雪"
	case "04":
		return "寒潮"
	case "05":
		return "大风"
	case "06":
		return "沙尘暴"
	case "07":
		return "高温"
	case "08":
		return "干旱"
	case "09":
		return "雷电"
	case "10":
		return "冰雹"
	case "11":
		return "霜冻"
	case "12":
		return "大雾"
	case "13":
		return "霾"
	case "14":
		return "道路结冰"
	case "15":
		return "森林火险"
	case "16":
		return "雷雨大风"
	case "17":
		return "春季沙尘天气趋势预警"
	case "18":
		return "沙尘"
	}
	return ""
}
func GetCodeId(code string) string {
	lastTwo := code[len(code)-2 : 0]
	switch lastTwo {
	case "00":
		return "白色"
	case "01":
		return "蓝色"
	case "02":
		return "黄色"
	case "03":
		return "橙色"
	case "04":
		return "红色"
	}
	return ""
}

func GetCodeIdColor(code string) string {
	lastTwo := code[len(code)-2 : 0]
	switch lastTwo {
	case "00":
		return "white"
	case "01":
		return "blue"
	case "02":
		return "yellow"
	case "03":
		return "orange"
	case "04":
		return "red"
	}
	return ""
}
