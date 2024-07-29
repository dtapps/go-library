package gojson

import (
	"encoding/json"
	"io"
	"strings"
)

// Marshal 将 Go 数据结构转换为 JSON 字符串。
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal 将 JSON 字符串解析并映射到 Go 数据结构中。
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// NewDecoder 创建一个新的 JSON 解码器（应返回解码器）。
func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

// NewEncoder 创建一个新的 JSON 编码器（应返回编码器）。
func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

// Encode 将 Go 数据结构编码为 JSON 字符串。
func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// MarshalToString 同 Encode 函数，将 Go 数据结构编码为 JSON 字符串。
func MarshalToString(msg interface{}) (string, error) {
	j, e := json.Marshal(msg)
	if e != nil {
		return "", e
	}
	return string(j), nil
}

// JsonDecode 将 JSON 字符串解析为 map 类型。
func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

// JsonDecodeNoError 同 JsonDecode，但忽略错误。
func JsonDecodeNoError(data string) map[string]interface{} {
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

// JsonEncode 将 Go 数据结构编码为 JSON 字符串。
func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

// JsonEncodeNoError 同 JsonEncode，但忽略错误。
func JsonEncodeNoError(data interface{}) string {
	jsons, _ := json.Marshal(data)
	return string(jsons)
}

// JsonDecodesNoError 将 JSON 字符串解析为字符串数组，忽略错误。
func JsonDecodesNoError(data string) []string {
	var dat []string
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

// ParseQueryString 解析 URL 查询字符串为 map 类型。
func ParseQueryString(input string) map[string]interface{} {
	paramMap := make(map[string]interface{})
	keyValuePairs := strings.Split(input, "&")
	for _, pair := range keyValuePairs {
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			paramMap[key] = value
		}
	}
	return paramMap
}

// IsValidJSON 检查给定字符串是否为有效的 JSON 格式。
func IsValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
