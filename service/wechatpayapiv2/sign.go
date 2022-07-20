package wechatpayapiv2

import (
	"bytes"
	"fmt"
	"github.com/dtapps/go-library/utils/gomd5"
	"sort"
	"strings"
)

// 支付字符串拼接
func (c *Client) getSortString(m map[string]interface{}) string {
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

// 获取签名
func (c *Client) getMd5Sign(paramMap map[string]interface{}) string {
	sortString := c.getSortString(paramMap)
	sign := gomd5.Md5(sortString + "&key=" + c.GetMchKey())
	return strings.ToUpper(sign)
}

// 验证签名
func (c *Client) checkMd5Sign(rspMap map[string]interface{}, sign string) bool {
	calculateSign := c.getMd5Sign(rspMap)
	return calculateSign == sign
}
