package float

import "math"

// ToInt64 将float64转成精确的int64
func ToInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// ToFloat64 精准float64
func ToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}
