package gostring

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ToString 转换成string
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

// ToFloat64 string到float64
func ToFloat64(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

// ToInt string到int
func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// ToInt64 string到int64
func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	return int64(ToFloat64(s))
}

// ToUint string到uint64
func ToUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return uint(i)
	}
	return 0
}

// ToUint64 string到uint64
func ToUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return i
	}
	return 0
}

// Replace 字符串替换
func Replace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

func HmacSha256Hex(key, strToSign string) string {
	hasHer := hmac.New(sha256.New, []byte(key))
	hasHer.Write([]byte(strToSign))
	return hex.EncodeToString(hasHer.Sum(nil))
}

// Space 去除空格
func Space(str string) string {
	return strings.Replace(str, " ", "", -1)
}

// LineBreak 去除换行符
func LineBreak(str string) string {
	return strings.Replace(str, "\n", "", -1)
}

// SpaceAndLineBreak 去除空格和去除换行符
func SpaceAndLineBreak(str string) string {
	return LineBreak(Space(str))
}

// TrimLastChar 删除字符串中的最后一个
func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

// Split 字符串分隔
func Split(s string, sep string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}

// Contains 判断字符串是否包含某个字符
func Contains(s, sep string) bool {
	return strings.Contains(s, sep)
}

func NumericalToString(value interface{}) (string, bool) {
	var val string

	switch value.(type) {
	default:
		return "0", false
	case int:
		intVal, _ := value.(int)
		val = strconv.FormatInt(int64(intVal), 10)
	case int8:
		intVal, _ := value.(int8)
		val = strconv.FormatInt(int64(intVal), 10)
	case int16:
		intVal, _ := value.(int16)
		val = strconv.FormatInt(int64(intVal), 10)
	case int32:
		intVal, _ := value.(int32)
		val = strconv.FormatInt(int64(intVal), 10)
	case int64:
		intVal, _ := value.(int64)
		val = strconv.FormatInt(int64(intVal), 10)
	case uint:
		intVal, _ := value.(uint)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint8:
		intVal, _ := value.(uint8)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint16:
		intVal, _ := value.(uint16)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint32:
		intVal, _ := value.(uint32)
		val = strconv.FormatUint(uint64(intVal), 10)
	case uint64:
		intVal, _ := value.(uint64)
		val = strconv.FormatUint(intVal, 10)
	case float32:
		floatVal, _ := value.(float32)
		val = strconv.FormatFloat(float64(floatVal), 'f', -1, 32)
	case float64:
		floatVal, _ := value.(float64)
		val = strconv.FormatFloat(floatVal, 'f', -1, 64)
	}
	return val, true
}
