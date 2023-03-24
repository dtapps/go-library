package kashangwl

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// md5(key + 参数1名称 + 参数1值 + 参数2名称 + 参数2值...) 加密源串应为{key}customer_id1192442order_id827669582783timestamp1626845767
func (c *Client) getSign(customerKey string, params map[string]interface{}) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	query := bytes.NewBufferString(customerKey)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(gorequest.GetParamsString(params[k]))
	}
	// MD5加密
	h := md5.New()
	io.Copy(h, query)
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

// 获取请求数据
func (c *Client) getRequestData(params map[string]interface{}) string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range params {
		args.Set(key, c.getString(val))
	}
	return args.Encode()
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
