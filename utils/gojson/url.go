package gojson

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var structFieldCache sync.Map // 缓存结构体字段 tag -> index

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

type fieldMeta struct {
	index []int // 支持嵌套字段路径
	tag   string
}

// BindURLQueryToStruct 将 URL 查询参数绑定到结构体字段（支持缓存 + 匿名嵌套结构体）
func BindURLQueryToStruct(params map[string]any, obj any) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("BindURLQueryToStruct: obj must be a non-nil pointer")
	}
	v = v.Elem()
	t := v.Type()

	// 检查缓存
	cacheKey := t.PkgPath() + "." + t.Name()
	cached, ok := structFieldCache.Load(cacheKey)
	var fields []fieldMeta
	if !ok {
		fields = parseFields(t, nil) // 仅当缓存未命中时进行解析
		structFieldCache.Store(cacheKey, fields)
	} else {
		fields = cached.([]fieldMeta)
	}

	// 按照字段路径绑定参数
	for _, meta := range fields {
		field := v.FieldByIndex(meta.index)
		if !field.CanSet() {
			continue
		}

		strVal, ok := extractFirst(params, meta.tag)
		if !ok || strVal == "" {
			continue
		}

		if err := setField(field, strVal); err != nil {
			return err
		}
	}

	return nil
}

// parseFields 递归解析结构体字段，包括匿名嵌套字段
func parseFields(t reflect.Type, parentIndex []int) []fieldMeta {
	var fields []fieldMeta
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}

		index := append([]int{}, parentIndex...)
		index = append(index, i)

		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			fields = append(fields, parseFields(f.Type, index)...)
			continue
		}

		tag := f.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(f.Name)
		}
		fields = append(fields, fieldMeta{index: index, tag: tag})
	}
	return fields
}

// setField 根据值类型设置字段
func setField(field reflect.Value, strVal string) error {
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
	return nil
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
