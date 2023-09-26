package eastiot

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gorequest"
	"sort"
	"strconv"
)

func (c *Client) getSign(p gorequest.Params) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, c.getString(p.Get(key)))
	}
	signStr += fmt.Sprintf("apiKey=%s", c.GetApiKey())
	return gomd5.ToUpper(signStr)
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
		bytes, _ := gojson.Marshal(v)
		return string(bytes)
	}
}
