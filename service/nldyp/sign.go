package nldyp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"github.com/dtapps/go-library/utils/gotime"
	"sort"
	"strings"
)

func (c *Client) Sign(p gorequest.Params) gorequest.Params {
	p.Set("vendor", c.GetVendor())
	p.Set("ts", gotime.Current().Timestamp())
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += key + gostring.GetString(p.Get(key))
	}
	p.Set("sign", createSign(fmt.Sprintf("%s%s%s", c.GetVendor(), p.Get("ts"), signStr)))
	return p
}

// 签名
func createSign(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
