package gostring

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go.dtapp.net/library/utils/gojson"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
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

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

// ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

func GetString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:

		bytes, _ := gojson.Marshal(v)
		return string(bytes)
	}
}

// IsNotChineseOrDigit 检查字符串包含中文字符或数字字符 则返回 true；否则返回 false。
func IsNotChineseOrDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// IsNotChinese 检查字符串包含中文字符 则返回 true；否则返回 false。
func IsNotChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// IsNotDigit 判断字符串是否包含数字字符
//
// 参数：
//
//	s string: 要判断的字符串
//
// 返回值：
//
//	bool: 如果字符串中包含数字字符，则返回true；否则返回false
//
// 说明：
//
//	该函数遍历字符串s中的每个字符，使用unicode.IsDigit函数判断字符是否为数字。
//	如果找到任何一个数字字符，则返回true；如果遍历完整个字符串后仍未找到数字字符，则返回false。
func IsNotDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// TruncateStringRune 截断字符串
//
// 参数：
//
//	s string: 要截断的字符串
//	maxLength int: 截断后的最大长度
//
// 返回值：
//
//	string: 截断后的字符串
//
// 说明：
//
//	该函数将字符串s截断为最多maxLength个字符，如果字符串s的长度小于等于maxLength，则返回原字符串；
//	否则返回前maxLength个字符组成的字符串。注意这里是按rune（Unicode码点）来截断，而不是按字节。
func TruncateStringRune(s string, maxLength int) string {
	runes := []rune(s)
	if len(runes) > maxLength {
		return string(runes[:maxLength])
	}
	return s
}
