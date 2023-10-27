package gostring

import (
	"math/rand"
	"time"
)

// GenerateRandomString 生成count个长度length不重复的随机数
func GenerateRandomString(length, count int) []string {
	return GenerateRandomStringFunc(length, count, func(num string) bool {
		return false
	})
}

// GenerateRandomStringFunc 生成count个长度length不重复的随机数，支持外部查询
func GenerateRandomStringFunc(length, count int, dFun func(num string) bool) []string {

	var fI int = 0
	startStr := "1"
	endStr := "9"
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	for {
		if fI+2 > length {
			break
		}
		startStr += "0"
		endStr += "9"

		fI = fI + 1
	}

	return GenerateRandomStringNumber(ToInt(startStr), ToInt(endStr), count, alphabet, length, dFun)
}

// GenerateRandomStringNumber 生成count个[start,end)结束的不重复的随机数
func GenerateRandomStringNumber(start, end, count int, alphabet string, length int, dFun func(num string) bool) []string {

	// 范围检查
	if end < start || len(alphabet) == 0 || (end-start) < count {
		return nil
	}

	// 存放结果的slice
	results := make([]string, 0)

	for len(results) < count {

		// 生成随机字符串
		str := make([]byte, length)
		for i := range str {
			str[i] = alphabet[rand.Intn(len(alphabet))]
		}

		// 查重
		exist := false
		for _, v := range results {
			if v == string(str) {
				exist = true
				break
			}
		}

		if !exist {
			isExist := dFun(string(str))
			if !isExist {
				results = append(results, string(str))
			}
		}
	}

	return results
}
