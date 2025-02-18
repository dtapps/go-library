package gorequest

import (
	"encoding/json"
	"fmt"
)

// GetParamsString 获取参数字符串
func GetParamsString(src any) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return fmt.Sprint(src)
	}
	data, _ := json.Marshal(src)
	return string(data)
}
