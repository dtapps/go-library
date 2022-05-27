package gorandom

import (
	"math/rand"
	"time"
)

const numbers string = "0123456789"
const letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const specials = "~!@#$%^*()_+-=[]{}|;:,./<>?"
const alphanumerics string = letters + numbers
const ascii string = alphanumerics + specials

func random[T int | int64](n T, chars string) string {
	if n <= 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n, n)
	l := len(chars)
	var i T = 0
	for {
		if i >= n {
			break
		}
		bytes[i] = chars[r.Intn(l)]
		i++
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
