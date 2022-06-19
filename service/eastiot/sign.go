package eastiot

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
	"sort"
	"strconv"
)

func (c *Client) getSign(p map[string]interface{}) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, c.getString(p[key]))
	}
	signStr += fmt.Sprintf("apiKey=%s", c.config.ApiKey)
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
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
}
