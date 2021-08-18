package random

import (
	"math/rand"
	"time"
)

const numbers string = "0123456789"
const letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const specials = "~!@#$%^*()_+-=[]{}|;:,./<>?"
const alphanumerics string = letters + numbers
const ascii string = alphanumerics + specials

func random(n int, chars string) string {
	if n <= 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n, n)
	l := len(chars)
	for i := 0; i < n; i++ {
		bytes[i] = chars[r.Intn(l)]
	}
	return string(bytes)
}

// Alphanumeric 随机字母数字
func Alphanumeric(n int) string {
	return random(n, alphanumerics)
}

// Alphabetic 随机字母
func Alphabetic(n int) string {
	return random(n, letters)
}

// Numeric 随机数字
func Numeric(n int) string {
	return random(n, numbers)
}

// Ascii 随机ASCII
func Ascii(n int) string {
	return random(n, ascii)
}
