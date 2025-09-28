package gostring

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// TrimSpace 去除前后空格
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
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

// SpaceAndLineBreak 去除空格和换行符（仅处理 ASCII 空格 ' ' 和换行 '\n'，不处理制表符 \t 等）。
// ⚠️ 注意：此函数不会移除制表符（\t）、回车（\r）等其他空白字符，如需完整清理请使用 CleanWhitespace。
func SpaceAndLineBreak(str string) string {
	return LineBreak(Space(str))
}

// CleanWhitespace 移除字符串中所有常见的空白字符，包括：
// 空格(' ')、制表符('\t')、换行符('\n')、回车符('\r')。
// 适用于清理用户输入需要严格格式化的场景。
func CleanWhitespace(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	return str
}

// RemoveAllWhitespace 移除字符串中所有 Unicode 定义的空白字符（包括中文全角空格、不间断空格等）。
// 使用 unicode.IsSpace 判断，覆盖范围最广，适用于国际化文本处理。
// 性能略低于 CleanWhitespace，但在处理普通 ASCII 数据时差异可忽略。
func RemoveAllWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1 // 删除该字符
		}
		return r
	}, str)
}

// TrimLastChar 删除字符串中的最后一个
func TrimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

// SplitNoEmpty 将字符串按 sep 分割，并过滤空值
func SplitNoEmpty(s string, sep string) []string {
	ts := TrimSpace(s)
	if len(ts) <= 0 {
		return []string{}
	}
	// 如果不包含分隔符
	if !Contains(ts, sep) {
		return []string{ts}
	}
	// 开始分割并清理
	parts := strings.Split(ts, sep)
	var result []string
	for _, part := range parts {
		trimmed := TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return strings.Split(s, sep)
}

// Split 字符串分隔（保留空值）
func Split(s string, sep string) []string {
	ts := TrimSpace(s)
	if len(ts) <= 0 {
		return []string{}
	}
	// 如果不包含分隔符
	if !Contains(ts, sep) {
		return []string{ts}
	}
	return strings.Split(ts, sep)
}

// Contains 判断字符串是否包含某个字符
func Contains(s, sep string) bool {
	return strings.Contains(s, sep)
}

func NumericalToString(value any) (string, bool) {
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

func GetString(i any) string {
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
		bytes, _ := json.Marshal(v)
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

// GetQueryParamSimple 从给定的URL字符串中提取指定名称的查询参数。
// 如果无法找到或解析出该参数，则返回空字符串。
// 假设查询参数的位置是固定的，不进行完整的查询字符串解析。
func GetQueryParamSimple(urlStr, paramName string) string {
	// 找到查询字符串部分
	queryStart := strings.IndexByte(urlStr, '?')
	if queryStart == -1 {
		return ""
	}

	// 查找参数名
	paramPattern := paramName + "="
	queryString := urlStr[queryStart+1:]
	pairs := strings.Split(queryString, "&")

	for _, pair := range pairs {
		if strings.HasPrefix(pair, paramPattern) {
			// 截取参数值
			return pair[len(paramPattern):]
		}
	}

	return ""
}

// EqualSlices 比较两个字符串切片是否相等
func EqualSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
