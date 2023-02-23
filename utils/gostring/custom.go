package gostring

import (
	"github.com/dtapps/go-library/utils/gotime"
)

// GenerateId 生成18位编号
func GenerateId(customId string) string {
	currentTime := gotime.Current().Format()
	one, err := generateIdOne(customId, currentTime, 18)
	if err == nil {
		return one
	}
	two, err := generateIdTwo(customId, currentTime, 18)
	if err == nil {
		return two
	}
	three, err := generateIdThree(customId, currentTime, 18)
	if err == nil {
		return three
	}
	four, err := generateIdFour(customId, currentTime, 18)
	if err == nil {
		return four
	}
	five, err := generateIdFive(customId, currentTime, 18)
	if err == nil {
		return five
	}
	six, err := generateIdSix(customId, 18)
	return six
}

// GenerateIdAndTime 生成18位编号
func GenerateIdAndTime(customId, customTime string) string {
	one, err := generateIdOne(customId, customTime, 18)
	if err == nil {
		return one
	}
	two, err := generateIdTwo(customId, customTime, 18)
	if err == nil {
		return two
	}
	three, err := generateIdThree(customId, customTime, 18)
	if err == nil {
		return three
	}
	four, err := generateIdFour(customId, customTime, 18)
	if err == nil {
		return four
	}
	five, err := generateIdFive(customId, customTime, 18)
	if err == nil {
		return five
	}
	six, err := generateIdSix(customId, 18)
	return six
}
