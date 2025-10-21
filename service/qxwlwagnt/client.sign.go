package qxwlwagnt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func (c *Client) getSign(params map[string]string) string {
	// 1. 获取所有 key
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 2. 按英文字典序排序
	sort.Strings(keys)

	// 3. 拼接 key=value&
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(params[k])
		sb.WriteString("&")
	}

	// 4. 拼接 appSecret
	sb.WriteString("appSecret=")
	sb.WriteString(c.config.appSecret)

	signStr := sb.String()

	// 5. 计算 MD5（32位小写）
	h := md5.Sum([]byte(signStr))
	sign := hex.EncodeToString(h[:])

	fmt.Println("[签名字符串]", signStr)
	fmt.Println("[签名结果]", sign)

	return sign
}
