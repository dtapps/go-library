package gostring

import (
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gorandom"
	"github.com/dtapps/go-library/utils/gotime"
)

// 生成18位时间[年月日时分]唯一编号
func generateIdOne(customId, setTime string, dataLength int) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
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
func generateIdTwo(customId, setTime string, dataLength int) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
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
func generateIdThree(customId, setTime string, dataLength int) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
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
func generateIdFour(customId, setTime string, dataLength int) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度
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
func generateIdFive(customId, setTime string, dataLength int) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength     = 5               // 随机数据长度`
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
func generateIdSix(customId string, dataLength int) (string, error) {

	var (
		newRandomLength = 0             // 随机数据长度
		customIdLength  = len(customId) // 自定义长度
	)

	const (
		randomLength = 5 // 随机数据长度
	)

	// 重新计算随机数据长度
	newRandomLength = dataLength - customIdLength

	if (customIdLength > dataLength) || (customIdLength == dataLength) || (newRandomLength < randomLength) {
		return "", errors.New("没有满足条件")
	}
	return fmt.Sprintf("%v%s", customId, gorandom.Numeric(newRandomLength)), nil
}
