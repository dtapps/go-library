package gostring

import (
	"github.com/spf13/cast"
)

// ToString 转 string
func ToString(v any) string {
	return cast.ToString(v)
}

// ToFloat64 转 float64
func ToFloat64(v any) float64 {
	return cast.ToFloat64(v)
}

// ToInt 转 int
func ToInt(v any) int {
	return cast.ToInt(v)
}

// ToInt64 转 int64
func ToInt64(v any) int64 {
	return cast.ToInt64(v)
}

// ToUint 转 uint64
func ToUint(v any) uint {
	return cast.ToUint(v)
}

// ToUint64 转 uint64
func ToUint64(v any) uint64 {
	return cast.ToUint64(v)
}
