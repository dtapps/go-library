package wechatpayapiv2

import (
	"bytes"
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
	"sort"
	"strings"
)

// 支付字符串拼接
func (app *App) getSortString(m map[string]interface{}) string {
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
func (app *App) getMd5Sign(paramMap map[string]interface{}) string {
	sortString := app.getSortString(paramMap)
	sign := gomd5.Md5(sortString + "&key=" + app.MchKey)
	return strings.ToUpper(sign)
}

// 验证签名
func (app *App) checkMd5Sign(rspMap map[string]interface{}, sign string) bool {
	calculateSign := app.getMd5Sign(rspMap)
	return calculateSign == sign
}
