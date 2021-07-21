package md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func Md5ToUpper(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}

// GetMD5Encode 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Get16MD5Encode 返回一个16位md5加密后的字符串
func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}
