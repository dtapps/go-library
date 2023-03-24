package nldyp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gotime"
	"sort"
	"strconv"
	"strings"
)

func (c *Client) Sign(p map[string]interface{}) map[string]interface{} {
	p["vendor"] = c.GetVendor()
	p["ts"] = gotime.Current().Timestamp()
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += key + getString(p[key])
	}
	p["sign"] = createSign(fmt.Sprintf("%s%s%s", c.GetVendor(), p["ts"], signStr))
	return p
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := gojson.Marshal(v)
		return string(bytes)
	}
}

// 签名
func createSign(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
