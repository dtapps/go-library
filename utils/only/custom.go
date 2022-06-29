package only

import (
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gotime"
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
	four, err := generateIdFour(customId)
	return four
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
	four, err := generateIdFour(customId)
	return four
}

// 生成18位时间[年月日时分]唯一编号
func generateIdOne(customId, setTime string) (string, error) {

	var (
		randomLength   = 4             // 随机数据长度
		customIdLength = len(customId) // 自定义长度
	)

	const (
		dataLength       = 18              // 默认数据长度
		dateFormat       = "200601021504"  // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	randomLength = dataLength - (dateFormatLength + customIdLength)

	if dateFormatLength+customIdLength > dataLength {
		return "", errors.New("超出长度")
	} else if dateFormatLength+customIdLength == dataLength {
		return "", errors.New("无法使用时间方法")
	} else if randomLength < 3 {
		return "", errors.New("自定义的数据过长")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(randomLength)), nil
}

// 生成18位时间[年月日时]唯一编号
func generateIdTwo(customId, setTime string) (string, error) {

	var (
		randomLength   = 4             // 随机数据长度
		customIdLength = len(customId) // 自定义长度
	)

	const (
		dataLength       = 18              // 默认数据长度
		dateFormat       = "2006010215"    // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	randomLength = dataLength - (dateFormatLength + customIdLength)

	if dateFormatLength+customIdLength > dataLength {
		return "", errors.New("超出长度")
	} else if dateFormatLength+customIdLength == dataLength {
		return "", errors.New("无法使用时间方法")
	} else if randomLength < 3 {
		return "", errors.New("自定义的数据过长")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(randomLength)), nil
}

// 生成18位时间[年月日]唯一编号
func generateIdThree(customId, setTime string) (string, error) {

	var (
		randomLength   = 4             // 随机数据长度
		customIdLength = len(customId) // 自定义长度
	)

	const (
		dataLength       = 18              // 默认数据长度
		dateFormat       = "20060102"      // 时间格式
		dateFormatLength = len(dateFormat) // 时间格式长度
	)

	// 重新计算随机数据长度
	randomLength = dataLength - (dateFormatLength + customIdLength)

	if dateFormatLength+customIdLength > dataLength {
		return "", errors.New("超出长度")
	} else if dateFormatLength+customIdLength == dataLength {
		return "", errors.New("无法使用时间方法")
	} else if randomLength < 3 {
		return "", errors.New("自定义的数据过长")
	}
	return fmt.Sprintf("%v%s%s", customId, gotime.SetCurrentParse(setTime).SetFormat(dateFormat), gorandom.Numeric(randomLength)), nil
}

// 生成18位随机唯一编号
func generateIdFour(customId string) (string, error) {

	var (
		randomLength   = 4             // 随机数据长度
		customIdLength = len(customId) // 自定义长度
	)

	const (
		dataLength = 18 // 默认数据长度
	)

	// 重新计算随机数据长度
	randomLength = dataLength - customIdLength

	if customIdLength >= dataLength {
		return "", errors.New("超出长度")
	} else if randomLength < 3 {
		return "", errors.New("自定义的数据过长")
	}
	return fmt.Sprintf("%v%s", customId, gorandom.Numeric(randomLength)), nil
}
