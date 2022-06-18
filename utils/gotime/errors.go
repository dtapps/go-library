package gotime

import "fmt"

// invalidTimezoneError  无效的时区错误
var invalidTimezoneError = func(timezone string) error {
	return fmt.Errorf("invalid timezone %q, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
}
