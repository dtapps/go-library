package goint64

import (
	"math"
	"strconv"
)

// ToFloat64 int64到float64
func ToFloat64(n int64) float64 {
	return float64(n) / math.Pow10(0)
}

// ToUnwrap 将int64恢复成正常的float64
func ToUnwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// ToUnwrapToInt64 精准int64
func ToUnwrapToInt64(num int64, retain int) int64 {
	return int64(ToUnwrap(num, retain))
}

// ToFloat64NewWiFi 返回float64
func ToFloat64NewWiFi(num int64) float64 {
	return float64(num / 100)
}

// ToString int到string
func ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}
