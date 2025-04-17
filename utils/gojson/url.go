package gojson

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// ParseURLQuery 使用标准库解析 URL 查询字符串为 map[string]any
func ParseURLQuery(input string) (map[string]any, error) {
	values, err := url.ParseQuery(input)
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	for k, v := range values {
		if len(v) == 1 {
			result[k] = v[0]
		} else {
			result[k] = v
		}
	}
	return result, nil
}

// ParseURLQueryWithoutError 忽略错误的版本
func ParseURLQueryWithoutError(input string) map[string]any {
	values, _ := url.ParseQuery(input)
	result := make(map[string]any)
	for k, v := range values {
		if len(v) == 1 {
			result[k] = v[0]
		} else {
			result[k] = v
		}
	}
	return result
}

// 提取第一个参数值（string）
func extractFirst(params map[string]any, tag string) (string, bool) {
	val, ok := params[tag]
	if !ok {
		return "", false
	}
	switch v := val.(type) {
	case string:
		return v, true
	case []string:
		if len(v) > 0 {
			return v[0], true
		}
	}
	return "", false
}

// BindURLQueryToStruct 将 URL 查询参数绑定到结构体字段
func BindURLQueryToStruct(params map[string]any, obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("BindURLQueryToStruct: obj必须是非零指针")
	}
	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue
		}

		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("form")
		if tag == "" {
			tag = strings.ToLower(fieldType.Name)
		}

		strVal, ok := extractFirst(params, tag)
		if !ok || strVal == "" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(strVal)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			num, err := strconv.ParseInt(strVal, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(num)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			num, err := strconv.ParseUint(strVal, 10, 64)
			if err != nil {
				return err
			}
			field.SetUint(num)

		case reflect.Bool:
			b, err := strconv.ParseBool(strVal)
			if err != nil {
				return err
			}
			field.SetBool(b)
		}
	}
	return nil
}
