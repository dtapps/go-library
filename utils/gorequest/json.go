package gorequest

import "encoding/json"

// IsValidJSON 检查给定字符串是否为有效的 JSON 格式。
func IsValidJSON(s string) bool {
	var js map[string]any
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsValidJSONWith 检查给定字符串是否为有效的 JSON 格式。
func IsValidJSONWith[T any](s string) bool {
	var dat T
	return json.Unmarshal([]byte(s), &dat) == nil
}

// JsonDecodeNoError 同 JsonDecode，但忽略错误。
func JsonDecodeNoError(data string) map[string]any {
	var dat map[string]any
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

// JsonDecodeNoErrorWith 同 JsonDecode，但忽略错误。
func JsonDecodeNoErrorWith[T any](data string) T {
	var dat T
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

// JsonEncodeNoError 同 JsonEncode，但忽略错误。
func JsonEncodeNoError(data any) string {
	jsons, _ := json.Marshal(data)
	return string(jsons)
}
