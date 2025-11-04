package wechatpayapiv3

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// EncryptOAEPWithPublicKey 使用 OAEP padding方式用公钥进行加密
func EncryptOAEPWithPublicKey(message string, publicKey *rsa.PublicKey) (ciphertext string, err error) {
	if publicKey == nil {
		return "", fmt.Errorf("you should input *rsa.PublicKey")
	}
	ciphertextByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, []byte(message), nil)
	if err != nil {
		return "", fmt.Errorf("encrypt message with public key err:%s", err.Error())
	}
	ciphertext = base64.StdEncoding.EncodeToString(ciphertextByte)
	return ciphertext, nil
}
