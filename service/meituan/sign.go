package meituan

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"io"
	"sort"
	"strconv"
)

// 签名(sign)生成逻辑（新版）
// https://union.meituan.com/v2/apiDetail?id=27
func (c *Client) getSign(Secret string, param *gorequest.Params) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range param.ToMap() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := bytes.NewBufferString(Secret)
	for _, k := range keys {
		signStr.WriteString(k)
		signStr.WriteString(c.getString(param.Get(k)))
	}
	signStr.WriteString(Secret)
	// md5加密
	has := md5.New()
	io.Copy(has, signStr)
	return hex.EncodeToString(has.Sum(nil))
}

func (c *Client) getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		marshal, _ := gojson.Marshal(v)
		return string(marshal)
	}
}
