package gotime

import "time"

// Verification 验证字符串是否为时间
func Verification(str, layout string) (resp time.Time, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}
	location, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return time.Time{}, err
	}
	return location, nil
}
