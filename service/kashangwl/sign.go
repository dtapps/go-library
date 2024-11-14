package kashangwl

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"io"
	"net/url"
	"sort"
	"strings"
)

// md5(key + 参数1名称 + 参数1值 + 参数2名称 + 参数2值...) 加密源串应为{key}customer_id1192442order_id827669582783timestamp1626845767
func (c *Client) getSign(customerKey string, param *gorequest.Params) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range param.DeepCopy() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	query := bytes.NewBufferString(customerKey)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(gorequest.GetParamsString(param.Get(k)))
	}
	// MD5加密
	h := md5.New()
	io.Copy(h, query)
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

// 获取请求数据
func (c *Client) getRequestData(param *gorequest.Params) string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range param.DeepCopy() {
		args.Set(key, gostring.GetString(val))
	}
	return args.Encode()
}
