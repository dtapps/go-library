package godecimal

import "math"

// Abs 取绝对值
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Floor 向下取整
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Ceil 向上取整
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Round 就近取整
func Round(x float64) float64 {
	return math.Round(x)
}

// RoundPoint 就近取整并保留小数点
func RoundPoint(x float64) float64 {
	return math.Round(x*100) / 100
}

// Max 取较大值
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min 取较小值
func Min(x, y float64) float64 {
	return math.Min(x, y)
}
