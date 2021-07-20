package string

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256Hex(key, strToSign string) string {
	hasher := hmac.New(sha256.New, []byte(key))
	hasher.Write([]byte(strToSign))
	return hex.EncodeToString(hasher.Sum(nil))
}
