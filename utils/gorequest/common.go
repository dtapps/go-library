package gorequest

import (
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gostring"
	"log"
)

// GetParamsString 获取参数字符串
func GetParamsString(src interface{}) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return gostring.ToString(src)
	}
	data, err := gojson.Marshal(src)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
