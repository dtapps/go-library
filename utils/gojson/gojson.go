package gojson

import (
	"github.com/dtapps/go-library/utils/gojson/json"
	"strings"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func MarshalToString(msg interface{}) (string, error) {
	j, e := json.Marshal(msg)
	if e != nil {
		return "", e
	}
	return string(j), nil
}

func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

func JsonDecodeNoError(data string) map[string]interface{} {
	var dat map[string]interface{}
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

func JsonEncodeNoError(data interface{}) string {
	jsons, _ := json.Marshal(data)
	return string(jsons)
}

func JsonDecodesNoError(data string) []string {
	var dat []string
	_ = json.Unmarshal([]byte(data), &dat)
	return dat
}

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

func IsValidJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
