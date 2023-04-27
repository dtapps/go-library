package mq135

// 根据传感器电阻计算CO2 ppm的参数
var (
	PARA float64 = 116.6020682
	PARB float64 = 2.769034857
)

// 建模温度和湿度相关性的参数
var (
	CORA float64 = .00035
	CORB float64 = .02718
	CORC float64 = 1.39538
	CORD float64 = .0018
	CORE float64 = -.003333333
	CORF float64 = -.001923077
	CORG float64 = 1.130128205
)

// 用于校准目的的大气CO2水平
var (
	ATMOCO2 float64 = 414.47 // Global CO2 Aug 2021
)

// GetCorrectionFactor 获取校正系数以校正温度和湿度
func GetCorrectionFactor(t float64, h float64) float64 {
	if t < 20 {
		return CORA*t*t - CORB*t + CORC - (h-33.)*CORD
	} else {
		return CORE*t + CORF*h + CORG
	}
}

func GetCorrectedResistance(r float64, t float64, h float64) float64 {
	return r / GetCorrectionFactor(t, h)
}
