package gostring

import (
	"math/rand/v2"
	"time"
)

// GenerateRandom 生成count个长度length不重复的随机数
func GenerateRandom(length, count int) []int {
	return GenerateRandomFunc(length, count, func(num int) bool {
		return false
	})
}

// GenerateRandomFunc 生成count个长度length不重复的随机数，支持外部查询
func GenerateRandomFunc(length, count int, dFun func(num int) bool) []int {

	var fI int = 0
	startStr := "1"
	endStr := "9"

	for {
		if fI+2 > length {
			break
		}
		startStr += "0"
		endStr += "9"

		fI = fI + 1
	}

	return GenerateRandomNumber(ToInt(startStr), ToInt(endStr), count, dFun)
}

// GenerateRandomNumber 生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start, end, count int, dFun func(num int) bool) []int {

	// 范围检查
	if end < start || (end-start) < count {
		return nil
	}

	// 存放结果的slice
	nums := make([]int, 0)

	// 随机数生成器，加入时间戳保证每次生成的随机数不一样
	seed := time.Now().UnixNano()                          // rand内部运算的随机数
	r := rand.New(rand.NewPCG(uint64(seed), uint64(seed))) // rand计算得到的随机数
	for len(nums) < count {

		// 生成随机数
		num := r.IntN(end-start) + start

		// 查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			isExist := dFun(num)
			if isExist == false {
				nums = append(nums, num)
			}
		}
	}

	return nums
}
