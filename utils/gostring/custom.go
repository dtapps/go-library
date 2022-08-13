package gostring

import (
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gorandom"
	"github.com/dtapps/go-library/utils/gotime"
)

// GenerateId 生成18一编号
func GenerateId(customId string) string {
	currentTime := gotime.Current().Format()
	one, err := generateIdOne(customId, currentTime)
	if err == nil {
		return one
	}
	two, err := generateIdTwo(customId, currentTime)
	if err == nil {
		return two
	}
	three, err := generateIdThree(customId, currentTime)
	if err == nil {
		return three
	}
	four, err := generateIdFour(customId, currentTime)
	if err == nil {
		return four
	}
	five, err := generateIdFive(customId, currentTime)
	if err == nil {
		return five
	}
	six, err := generateIdSix(customId)
	return six
}

// GenerateIdAndTime 生成18一编号
func GenerateIdAndTime(customId, customTime string) string {
	one, err := generateIdOne(customId, customTime)
	if err == nil {
		return one
	}
	two, err := generateIdTwo(customId, customTime)
	if err == nil {
		return two
	}
	three, err := generateIdThree(customId, customTime)
	if err == nil {
		return three
	}
	four, err := generateIdFour(customId, customTime)
	if err == nil {
		return four
	}
	five, err := generateIdFive(customId, customTime)
	if err == nil {
		return five
	}
	six, err := generateIdSix(customId)
	return six
}

// 生成18位时间[年月日时分]唯一编号
func generateIdOne(customId, setTime string) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
		dataLength       = 18              // 默认数据长度
		dateFormat       = "200601021504"  // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - (dateFormatLength + customIdLength)

	if (dateFormatLength+customIdLength > dataLength) || (dateFormatLength+customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(newRandomLength)), nil
}

// 生成18位时间[年月日时]唯一编号
func generateIdTwo(customId, setTime string) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
		dataLength       = 18              // 默认数据长度
		dateFormat       = "2006010215"    // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - (dateFormatLength + customIdLength)

	if (dateFormatLength+customIdLength > dataLength) || (dateFormatLength+customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(newRandomLength)), nil
}

// 生成18位时间[年月日]唯一编号
func generateIdThree(customId, setTime string) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
		dataLength       = 18              // 默认数据长度
		dateFormat       = "20060102"      // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - (dateFormatLength + customIdLength)

	if (dateFormatLength+customIdLength > dataLength) || (dateFormatLength+customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(newRandomLength)), nil
}

// 生成18位时间[年月]唯一编号
func generateIdFour(customId, setTime string) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
		dataLength       = 18              // 默认数据长度
		dateFormat       = "200601"        // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - (dateFormatLength + customIdLength)

	if (dateFormatLength+customIdLength > dataLength) || (dateFormatLength+customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(newRandomLength)), nil
}

// 生成18位时间[年]唯一编号
func generateIdFive(customId, setTime string) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度`
		dataLength       = 18              // 默认数据长度
		dateFormat       = "2006"          // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - (dateFormatLength + customIdLength)

	if (dateFormatLength+customIdLength > dataLength) || (dateFormatLength+customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(newRandomLength)), nil
}

// 生成18位随机唯一编号
func generateIdSix(customId string) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength = 5  // 随机数据长度
		dataLength   = 18 // 默认数据长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - customIdLength

	if (customIdLength > dataLength) || (customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s", customId, gorandom.Numeric(newRandomLength)), nil
}
