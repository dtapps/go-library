package gojson

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// ParseURLQuery 解析 URL 查询字符串为 map 类型，处理 URL 编码和错误情况，返回错误
func ParseURLQuery(input string) (map[string]any, error) {
	paramMap := make(map[string]any)
	keyValuePairs := strings.Split(input, "&")
	for _, pair := range keyValuePairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			return nil, strconv.ErrSyntax
		}
		key, err := url.QueryUnescape(parts[0])
		if err != nil {
			return nil, err
		}
		value, err := url.QueryUnescape(parts[1])
		if err != nil {
			return nil, err
		}
		if existing, ok := paramMap[key]; ok {
			switch v := existing.(type) {
			case string:
				paramMap[key] = []string{v, value}
			case []string:
				paramMap[key] = append(v, value)
			default:
				paramMap[key] = []string{value}
			}
		} else {
			paramMap[key] = value
		}
	}
	return paramMap, nil
}

// ParseURLQueryWithoutError 解析 URL 查询字符串为 map 类型，不返回错误
func ParseURLQueryWithoutError(input string) map[string]any {
	paramMap := make(map[string]any)
	keyValuePairs := strings.Split(input, "&")
	for _, pair := range keyValuePairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, err := url.QueryUnescape(parts[0])
		if err != nil {
			continue
		}
		value, err := url.QueryUnescape(parts[1])
		if err != nil {
			continue
		}
		if existing, ok := paramMap[key]; ok {
			switch v := existing.(type) {
			case string:
				paramMap[key] = []string{v, value}
			case []string:
				paramMap[key] = append(v, value)
			default:
				paramMap[key] = []string{value}
			}
		} else {
			paramMap[key] = value
		}
	}
	return paramMap
}

// BindURLQueryToStruct 将解析后的查询参数绑定到结构体，处理类型转换和错误情况
func BindURLQueryToStruct(params map[string]any, obj any) error {
	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("form")
		if tag == "" {
			tag = strings.ToLower(fieldType.Name)
		}

		if value, ok := params[tag]; ok {
			switch field.Kind() {
			case reflect.String:
				if strValue, ok := value.(string); ok {
					field.SetString(strValue)
				} else if strSlice, ok := value.([]string); ok && len(strSlice) > 0 {
					field.SetString(strSlice[0])
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var numStr string
				if strValue, ok := value.(string); ok {
					numStr = strValue
				} else if strSlice, ok := value.([]string); ok && len(strSlice) > 0 {
					numStr = strSlice[0]
				}
				if numStr != "" {
					num, err := strconv.ParseInt(numStr, 10, 64)
					if err != nil {
						return err
					}
					field.SetInt(num)
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				var numStr string
				if strValue, ok := value.(string); ok {
					numStr = strValue
				} else if strSlice, ok := value.([]string); ok && len(strSlice) > 0 {
					numStr = strSlice[0]
				}
				if numStr != "" {
					num, err := strconv.ParseUint(numStr, 10, 64)
					if err != nil {
						return err
					}
					field.SetUint(num)
				}
			case reflect.Bool:
				var boolStr string
				if strValue, ok := value.(string); ok {
					boolStr = strValue
				} else if strSlice, ok := value.([]string); ok && len(strSlice) > 0 {
					boolStr = strSlice[0]
				}
				if boolStr != "" {
					b, err := strconv.ParseBool(boolStr)
					if err != nil {
						return err
					}
					field.SetBool(b)
				}
			}
		}
	}
	return nil
}
