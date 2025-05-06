package wikeyun

import (
	"crypto/md5"
	"encoding/hex"
	"go.dtapp.net/library/utils/gorequest"
	"sort"
	"strconv"
	"strings"
	"time"
)

type respSign struct {
	AppKey    int64
	Timestamp string
	Client    string
	V         string
	Format    string
	Sign      string
}

// 签名
func (c *Client) sign(param *gorequest.Params) respSign {
	// 默认参数
	v := "1.0"
	format := "json"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	param.Set("v", v)                   // 客户端接口版本，目前是1.0
	param.Set("format", format)         // 默认json
	param.Set("app_key", c.GetAppKey()) // 应用唯一表示
	param.Set("client", c.clientIP)     // 客户端请求ip
	param.Set("timestamp", timestamp)   // unix时间戳（秒单位）
	// 排序所有的 key
	var keys []string
	for key := range param.DeepGetAny() {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := c.GetAppSecret()
	for _, key := range keys {
		signStr += key + gorequest.GetString(param.Get(key))
	}
	signStr += c.GetAppSecret()
	return respSign{
		AppKey:    c.GetAppKey(),
		Timestamp: timestamp,
		Client:    c.clientIP,
		V:         v,
		Format:    format,
		Sign:      c.createSign(signStr),
	}
}

// 签名
func (c *Client) createSign(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
