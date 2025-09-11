package gostring

import (
	"strings"
	"unicode"
)

type WechatTruncate struct {
}

func NewWechatTruncate() *WechatTruncate {
	return &WechatTruncate{}
}

// Thing 事物 - 20个以内字符（汉字、数字、字母、符号）
func (t *WechatTruncate) Thing(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) > maxLen {
		return string(runes[:maxLen])
	}
	return s
}

// Number 数字 - 32位以内数字（可带小数）
func (t *WechatTruncate) Number(s string, maxLen int) string {
	// 只保留数字和小数点
	allowed := "0123456789."
	var filtered []rune
	for _, r := range s {
		if strings.ContainsRune(allowed, r) {
			filtered = append(filtered, r)
		}
	}
	result := string(filtered)
	if len([]rune(result)) > maxLen {
		result = string([]rune(result)[:maxLen])
	}
	return result
}

// Letter 字母 - 32位以内字母
func (t *WechatTruncate) Letter(s string, maxLen int) string {
	var filtered []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			filtered = append(filtered, r)
		}
	}
	result := string(filtered)
	if len([]rune(result)) > maxLen {
		result = string([]rune(result)[:maxLen])
	}
	return result
}

// Symbol 符号 - 5位以内符号（非字母、非数字、非汉字）
func (t *WechatTruncate) Symbol(s string, maxLen int) string {
	var filtered []rune
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !t.isChinese(r) {
			filtered = append(filtered, r)
		}
	}
	result := string(filtered)
	if len([]rune(result)) > maxLen {
		result = string([]rune(result)[:maxLen])
	}
	return result
}

// CharacterString 字符串 - 32位以内数字、字母或符号
func (t *WechatTruncate) CharacterString(s string, maxLen int) string {
	var filtered []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || !t.isChinese(r) && !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			filtered = append(filtered, r)
		}
	}
	result := string(filtered)
	if len([]rune(result)) > maxLen {
		result = string([]rune(result)[:maxLen])
	}
	return result
}

// Time 时间 - 不截断，只校验格式（你可自定义格式化）
func (t *WechatTruncate) Time(s string) string {
	// 微信支持如 "15:01" 或 "2019年10月1日 15:01"
	// 这里不截断，因为时间格式固定，超长说明数据异常
	return s
}

// Date 日期 - 同上，不截断
func (t *WechatTruncate) Date(s string) string {
	return s
}

// Amount 金额 - 1币种符号+10位数字（可小数）+可选“元”
func (t *WechatTruncate) Amount(s string, maxDigits int) string {
	// 示例: "¥123.45元" 或 "$999.99"
	// 策略：保留第一个非数字字符（币种），然后取最多 maxDigits 位数字+小数点，最后可加“元”
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}

	var result []rune
	// 1. 取第一个可能是币种的符号
	if !unicode.IsDigit(runes[0]) && runes[0] != '.' {
		result = append(result, runes[0])
		runes = runes[1:]
	}

	// 2. 取数字和小数点，最多 maxDigits 个字符
	count := 0
	for _, r := range runes {
		if count >= maxDigits {
			break
		}
		if unicode.IsDigit(r) || r == '.' {
			result = append(result, r)
			count++
		}
	}

	// 3. 如果原字符串以“元”结尾，且结果还没加，可以补上
	if strings.HasSuffix(s, "元") && !strings.HasSuffix(string(result), "元") {
		result = append(result, '元')
	}

	return string(result)
}

// PhoneNumber 电话 - 17位以内，数字和符号
func (t *WechatTruncate) PhoneNumber(s string, maxLen int) string {
	allowed := "0123456789+-() "
	var filtered []rune
	for _, r := range s {
		if strings.ContainsRune(allowed, r) {
			filtered = append(filtered, r)
		}
	}
	result := string(filtered)
	if len([]rune(result)) > maxLen {
		result = string([]rune(result)[:maxLen])
	}
	return result
}

// CarNumber 车牌 - 8位以内，首尾可汉字，中间字母数字
func (t *WechatTruncate) CarNumber(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) > maxLen {
		runes = runes[:maxLen]
	}

	// 校验：只允许首尾是汉字，其他位置只能字母或数字
	for i, r := range runes {
		if i == 0 || i == len(runes)-1 {
			// 首位或末位：允许汉字、字母、数字
			if !(t.isChinese(r) || unicode.IsLetter(r) || unicode.IsDigit(r)) {
				runes[i] = 'X' // 替换为占位符，或直接跳过
			}
		} else {
			// 中间位置：只允许字母或数字
			if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
				runes[i] = 'X'
			}
		}
	}
	return string(runes)
}

// Name 姓名 - 10汉字内 或 20字母内，混合按汉字算
func (t *WechatTruncate) Name(s string, maxChinese int, maxAlpha int) string {
	runes := []rune(s)
	chineseCount := 0
	alphaCount := 0

	for _, r := range runes {
		if t.isChinese(r) {
			chineseCount++
		} else if unicode.IsLetter(r) {
			alphaCount++
		}
		// 忽略其他字符（如空格、符号）
	}

	if chineseCount > 0 {
		// 有汉字 → 按汉字限制（10字内）
		if len(runes) > maxChinese {
			return string(runes[:maxChinese])
		}
	} else {
		// 纯字母 → 按字母限制（20字内）
		if len(runes) > maxAlpha {
			return string(runes[:maxAlpha])
		}
	}
	return s
}

// Phrase 汉字 - 5个以内纯汉字
func (t *WechatTruncate) Phrase(s string, maxLen int) string {
	var filtered []rune
	for _, r := range s {
		if t.isChinese(r) {
			filtered = append(filtered, r)
		}
	}
	if len(filtered) > maxLen {
		filtered = filtered[:maxLen]
	}
	return string(filtered)
}

// isChinese 判断是否为汉字
func (t *WechatTruncate) isChinese(r rune) bool {
	return unicode.Is(unicode.Han, r)
}
