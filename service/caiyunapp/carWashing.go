package caiyunapp

// GetCarWashingDesc 洗车指数 https://docs.caiyunapp.com/docs/tables/lifeindex
func GetCarWashingDesc(carWashing float64) string {
	if carWashing <= 1 {
		return "适宜"
	} else if carWashing <= 2 {
		return "较适宜"
	} else if carWashing <= 3 {
		return "较不适宜"
	} else if carWashing <= 4 {
		return "不适应"
	}
	return ""
}
