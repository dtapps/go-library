package wikeyun

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"
)

type respSign struct {
	AppKey    int
	Timestamp string
	Client    string
	V         string
	Format    string
	Sign      string
}

// 签名
func (c *Client) sign(params map[string]interface{}) respSign {
	// 默认参数
	v := "1.0"
	format := "json"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params["v"] = v                     // 客户端接口版本，目前是1.0
	params["format"] = format           // 默认json
	params["app_key"] = c.config.AppKey // 应用唯一表示
	params["client"] = c.clientIp       // 客户端请求ip
	params["timestamp"] = timestamp     // unix时间戳（秒单位）
	// 排序所有的 key
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := c.config.AppSecret
	for _, key := range keys {
		signStr += key + getString(params[key])
	}
	signStr += c.config.AppSecret
	return respSign{
		AppKey:    c.config.AppKey,
		Timestamp: timestamp,
		Client:    c.clientIp,
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
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
}
