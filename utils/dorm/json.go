package dorm

import (
	"github.com/dtapps/go-library/utils/gojson"
)

// JsonDecodeNoError json字符串转结构体，不报错
func JsonDecodeNoError(b []byte) map[string]interface{} {
	var data map[string]interface{}
	_ = gojson.Unmarshal(b, &data)
	return data
}

// JsonEncodeNoError 结构体转json字符串，不报错
func JsonEncodeNoError(data interface{}) string {
	jsons, _ := gojson.Marshal(data)
	return string(jsons)
}
