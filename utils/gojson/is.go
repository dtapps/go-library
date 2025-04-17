package gojson

import "encoding/json"

// IsValidJSON 检查给定字符串是否为有效的 JSON 格式。
func IsValidJSON(s string) bool {
	var js map[string]any
	return json.Unmarshal([]byte(s), &js) == nil
}
