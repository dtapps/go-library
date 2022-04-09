package gotime

import (
	"time"
)

// 通过时区获取 Location 实例
func getLocationByTimezone(timezone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		err = invalidTimezoneError(timezone)
	}
	return loc, err
}
