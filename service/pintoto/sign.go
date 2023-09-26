package pintoto

import (
	"crypto/md5"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"net/url"
	"sort"
)

func (c *Client) getSign(appSecret string, p gorequest.Params) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gostring.GetString(p.Get(key)))
	}
	signStr += fmt.Sprintf("appSecret=%s", appSecret)
	// md5加密
	data := []byte(signStr)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// 获取请求数据
func (c *Client) getRequestData(param gorequest.Params) string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range param {
		args.Set(key, gostring.GetString(val))
	}
	return args.Encode()
}
