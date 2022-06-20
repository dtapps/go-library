package dorm

import "encoding/json"

// JsonDecodeNoError Json解码，不报错
func JsonDecodeNoError(b []byte) map[string]interface{} {
	var data map[string]interface{}
	_ = json.Unmarshal(b, &data)
	return data
}
