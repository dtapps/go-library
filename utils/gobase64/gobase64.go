package gobase64

import (
	"encoding/base64"
	"github.com/dtapps/go-library/utils/gophp"
)

// Encode base64编码
func Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// Decode base64解码
func Decode(input string) string {
	decodeString, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(decodeString)
}

// EncodeUrl Base64安全URL编码
func EncodeUrl(input string) string {
	return gophp.Rtrim(gophp.Strtr(base64.StdEncoding.EncodeToString([]byte(input)), "+/", "-_"), "=")
}

// DecodeUrl Base64安全URL解码
func DecodeUrl(input string) string {
	decodeString, err := base64.StdEncoding.DecodeString(gophp.StrPad(gophp.Strtr(input, "-_", "+/"), len(input)/4, "="))
	if err != nil {
		return ""
	}
	return string(decodeString)
}
