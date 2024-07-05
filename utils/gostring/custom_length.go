package gostring

import "go.dtapp.net/library/utils/gotime"

// GenerateIdLength 生成自定义长度编号
func GenerateIdLength(customId string, dataLength int) string {
	currentTime := gotime.Current().Format()
	one, err := generateIdOne(customId, currentTime, dataLength)
	if err == nil {
		return one
	}
	two, err := generateIdTwo(customId, currentTime, dataLength)
	if err == nil {
		return two
	}
	three, err := generateIdThree(customId, currentTime, dataLength)
	if err == nil {
		return three
	}
	four, err := generateIdFour(customId, currentTime, dataLength)
	if err == nil {
		return four
	}
	five, err := generateIdFive(customId, currentTime, dataLength)
	if err == nil {
		return five
	}
	six, err := generateIdSix(customId, dataLength)
	return six
}

// GenerateIdAndTimeLength 生成自定义长度编号
func GenerateIdAndTimeLength(customId, customTime string, dataLength int) string {
	one, err := generateIdOne(customId, customTime, dataLength)
	if err == nil {
		return one
	}
	two, err := generateIdTwo(customId, customTime, dataLength)
	if err == nil {
		return two
	}
	three, err := generateIdThree(customId, customTime, dataLength)
	if err == nil {
		return three
	}
	four, err := generateIdFour(customId, customTime, dataLength)
	if err == nil {
		return four
	}
	five, err := generateIdFive(customId, customTime, dataLength)
	if err == nil {
		return five
	}
	six, err := generateIdSix(customId, dataLength)
	return six
}
