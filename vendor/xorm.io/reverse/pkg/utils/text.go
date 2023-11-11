package utils

import "strings"

func UnTitle(src string) string {
	if src == "" {
		return ""
	}

	if len(src) == 1 {
		return strings.ToLower(string(src[0]))
	}
	return strings.ToLower(string(src[0])) + src[1:]
}

func UpTitle(src string) string {
	if src == "" {
		return ""
	}

	return strings.ToUpper(src)
}
