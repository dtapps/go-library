package goint

import (
	"math"
	"strconv"
)

// ToString int到string
func ToString(n int) string {
	return strconv.Itoa(n)
}

// ToFloat64 int到float64
func ToFloat64(n int) float64 {
	return float64(n) / math.Pow10(0)
}
