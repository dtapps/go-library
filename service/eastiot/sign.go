package eastiot

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"sort"
)

func (c *Client) getSign(p gorequest.Params) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gostring.GetString(p.Get(key)))
	}
	signStr += fmt.Sprintf("apiKey=%s", c.GetApiKey())
	return gomd5.ToUpper(signStr)
}
