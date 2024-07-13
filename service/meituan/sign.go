package meituan

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"io"
	"sort"
)

// 签名(sign)生成逻辑（新版）
// https://union.meituan.com/v2/apiDetail?id=27
func (c *Client) getSign(Secret string, param gorequest.Params) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := bytes.NewBufferString(Secret)
	for _, k := range keys {
		signStr.WriteString(k)
		signStr.WriteString(gostring.GetString(param.Get(k)))
	}
	signStr.WriteString(Secret)
	// md5加密
	has := md5.New()
	io.Copy(has, signStr)
	return hex.EncodeToString(has.Sum(nil))
}
