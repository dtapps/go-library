package wechatpayapiv2

import (
	"bytes"
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
	"sort"
	"strings"
)

// GetSign 获取签名
func GetSign(param map[string]any, key string) string {
	sortString := getSortString(param)
	sign := gomd5.Md5(sortString + "&key=" + key)
	return strings.ToUpper(sign)
}

// 支付字符串拼接
func getSortString(m map[string]any) string {
	var buf bytes.Buffer
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := m[k]
		if vs == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(fmt.Sprintf("%v", vs))
	}
	return buf.String()
}
