package wechatpayapiv2

import (
	"bytes"
	"fmt"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gorequest"
	"sort"
	"strings"
)

// 支付字符串拼接
func (c *Client) getSortString(m *gorequest.Params) string {
	var buf bytes.Buffer
	keys := make([]string, 0)
	for k := range m.ToMap() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := m.Get(k)
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
func (c *Client) getMd5Sign(paramMap *gorequest.Params) string {
	sortString := c.getSortString(paramMap)
	sign := gomd5.Md5(sortString + "&key=" + c.GetMchKey())
	return strings.ToUpper(sign)
}

// 验证签名
func (c *Client) checkMd5Sign(rspMap *gorequest.Params, sign string) bool {
	calculateSign := c.getMd5Sign(rspMap)
	return calculateSign == sign
}
