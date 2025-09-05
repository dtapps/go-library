package gorandom

import (
	"math/rand/v2"
	"time"
)

const numbers = "0123456789"
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const specials = "~!@#$%^*()_+-=[]{}|;:,./<>?"
const alphanumerics = letters + numbers
const ascii = alphanumerics + specials

func random[T int | int64](n T, chars string) string {
	if n <= 0 {
		return ""
	}

	bytes := make([]byte, n)
	charsLen := len(chars)

	for i := T(0); i < n; i++ {
		bytes[i] = chars[rand.IntN(charsLen)]
	}
	return string(bytes)
}

// Alphanumeric 随机字母数字
func Alphanumeric[T int | int64](n T) string {
	return random(n, alphanumerics)
}

// Alphabetic 随机字母
func Alphabetic[T int | int64](n T) string {
	return random(n, letters)
}

// Numeric 随机数字
func Numeric[T int | int64](n T) string {
	return random(n, numbers)
}

// Ascii 随机ASCII
func Ascii[T int | int64](n T) string {
	return random(n, ascii)
}

// Range 返回一个在 [min, max) 范围内的随机整数
func Range(min int64, max int64) int64 {
	if min >= max {
		return min
	}
	return int64(rand.IntN(int(max-min))) + min
}

// RangeTime 返回一个在 [min, max) 范围内的随机时间间隔
func RangeTime(min int64, max int64) time.Duration {
	return time.Duration(Range(min, max))
}
