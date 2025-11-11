package wechatopen

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
)

// Sign 微信公众号 url 签名.
func Sign(token, timestamp, nonce string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce}
	strs.Sort()

	buf := make([]byte, 0, len(token)+len(timestamp)+len(nonce))
	buf = append(buf, strs[0]...)
	buf = append(buf, strs[1]...)
	buf = append(buf, strs[2]...)

	hashsum := sha1.Sum(buf)
	return hex.EncodeToString(hashsum[:])
}

// MsgSign 微信公众号/企业号 消息体签名.
func MsgSign(token, timestamp, nonce, encryptedMsg string) (signature string) {
	strs := sort.StringSlice{token, timestamp, nonce, encryptedMsg}
	strs.Sort()

	h := sha1.New()

	bufw := bufio.NewWriterSize(h, 128) // sha1.BlockSize 的整数倍
	bufw.WriteString(strs[0])
	bufw.WriteString(strs[1])
	bufw.WriteString(strs[2])
	bufw.WriteString(strs[3])
	bufw.Flush()

	hashsum := h.Sum(nil)
	return hex.EncodeToString(hashsum)
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timeStamp, nonce string, token string) bool {
	paramsArray := []string{token, timeStamp, nonce}
	// 字典序排序
	sort.Strings(paramsArray)
	paramsMsg := ""
	for _, value := range paramsArray {
		//fmt.Println(value)
		paramsMsg += value
	}
	//sha1
	sha1Param := sha1.New()
	sha1Param.Write([]byte(paramsMsg))
	msg := hex.EncodeToString(sha1Param.Sum([]byte("")))
	return msg == signature
}

func AesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7

	if len(cipherData)%k != 0 {
		return nil, fmt.Errorf("crypto/cipher: 密文大小不是aes密钥长度的倍数")
	}

	// 创建加密算法实例
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	// 创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)

	return plainData, nil
}
