package alipayopen

import (
	"context"
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"log"
	"sort"
	"strings"
)

func (c *Client) rsaSign(signContent string, hash crypto.Hash) (string, error) {

	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)

	signByte, err := c.privateKey.Sign(rand.Reader, hashed, crypto.SHA256)
	if err != nil {
		return "", err
	}

	sign := base64.StdEncoding.EncodeToString(signByte)

	return sign, nil
}

func (c *Client) sign(ctx context.Context, param gorequest.Params) gorequest.Params {
	// 排序
	var keys []string
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 转换成字符串
	var signStr = ""
	for _, key := range keys {
		if key != "" {
			signStr += fmt.Sprintf("%s=%s&", key, gostring.GetString(param.Get(key)))
		}
	}

	signStr = strings.TrimRight(signStr, "&")

	// 签名
	sign, err := c.rsaSign(signStr, crypto.SHA256)
	if err != nil {
		log.Printf("签名失败：%s\n", err)
		return nil
	}
	param.Set("sign", sign)

	return param
}
