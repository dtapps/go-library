package caiyunapp

import "fmt"

// GetWindSpeed https://docs.caiyunapp.com/docs/tables/wind
func GetWindSpeed(speed float64) string {
	if speed <= 0 {
		return "0级"
	} else if speed <= 5 {
		return "1级"
	} else if speed <= 11 {
		return "2级"
	} else if speed <= 19 {
		return "3级"
	} else if speed <= 28 {
		return "4级"
	} else if speed <= 38 {
		return "5级"
	} else if speed <= 49 {
		return "6级"
	} else if speed <= 61 {
		return "7级"
	} else if speed <= 74 {
		return "8级"
	} else if speed <= 88 {
		return "9级"
	} else if speed <= 102 {
		return "10级"
	} else if speed <= 117 {
		return "11级"
	} else if speed <= 133 {
		return "12级"
	} else if speed <= 149 {
		return "13级"
	} else if speed <= 166 {
		return "14级"
	} else if speed <= 183 {
		return "15级"
	} else if speed <= 201 {
		return "16级"
	} else if speed <= 220 {
		return "17级"
	}
	return fmt.Sprintf("%v", speed)
}

// GetWindSpeedDesc https://docs.caiyunapp.com/docs/tables/wind
func GetWindSpeedDesc(speed float64) string {
	if speed <= 0 {
		return "无风"
	} else if speed <= 5 {
		return "微风徐徐"
	} else if speed <= 11 {
		return "清风"
	} else if speed <= 19 {
		return "树叶摇摆"
	} else if speed <= 28 {
		return "树枝摇动"
	} else if speed <= 38 {
		return "风力强劲"
	} else if speed <= 49 {
		return "风力强劲"
	} else if speed <= 61 {
		return "风力超强"
	} else if speed <= 74 {
		return "狂风大作"
	} else if speed <= 88 {
		return "狂风呼啸"
	} else if speed <= 102 {
		return "暴风毁树"
	} else if speed <= 117 {
		return "暴风毁树"
	} else if speed <= 133 {
		return "飓风"
	} else if speed <= 149 {
		return "台风"
	} else if speed <= 166 {
		return "强台风"
	} else if speed <= 183 {
		return "强台风"
	} else if speed <= 201 {
		return "超强台风"
	} else if speed <= 220 {
		return "超强台风"
	}
	return fmt.Sprintf("%v", speed)
}

// GetWindDirectionDesc https://docs.caiyunapp.com/docs/tables/wind
func GetWindDirectionDesc(direction float64) string {
	if direction <= 11.26 {
		return "北"
	} else if direction <= 33.75 {
		return "北东北"
	} else if direction <= 56.25 {
		return "东北"
	} else if direction <= 78.75 {
		return "东东北"
	} else if direction <= 101.25 {
		return "东"
	} else if direction <= 123.75 {
		return "东东南"
	} else if direction <= 146.25 {
		return "东南"
	} else if direction <= 168.75 {
		return "南东南"
	} else if direction <= 191.25 {
		return "南"
	} else if direction <= 213.75 {
		return "南西南"
	} else if direction <= 236.25 {
		return "西南"
	} else if direction <= 258.75 {
		return "西西南"
	} else if direction <= 281.25 {
		return "西"
	} else if direction <= 303.75 {
		return "西西北"
	} else if direction <= 326.25 {
		return "西北"
	} else if direction <= 348.75 {
		return "北西北"
	}
	return "北"
}
