package gorandom

import (
	"math/rand/v2"
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

	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	seed := time.Now().UnixNano()                          // rand内部运算的随机数
	r := rand.New(rand.NewPCG(uint64(seed), uint64(seed))) // rand计算得到的随机数

	bytes := make([]byte, n, n)
	l := len(chars)
	var i T = 0
	for {
		if i >= n {
			break
		}
		bytes[i] = chars[r.IntN(l)]
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

func Range(min int64, max int64) int64 {

	var number int64
	for {
		// 生成在范围 [min, max) 内的随机数
		number = int64(rand.IntN(int(max)-int(min)) + int(min))

		// 检查随机数是否大于 min 且小于 max
		if number > min && number < max {
			break
		}
	}

	return number
}
