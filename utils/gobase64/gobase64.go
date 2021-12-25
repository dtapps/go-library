package gobase64

import "encoding/base64"

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
