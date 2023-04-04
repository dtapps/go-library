package caiyunapp

// GetColdRiskDesc 感冒指数 https://docs.caiyunapp.com/docs/tables/lifeindex
func GetColdRiskDesc(coldRisk float64) string {
	if coldRisk <= 1 {
		return "少发"
	} else if coldRisk <= 2 {
		return "较易发"
	} else if coldRisk <= 3 {
		return "易发"
	} else if coldRisk <= 4 {
		return "极易发"
	}
	return ""
}
