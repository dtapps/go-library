package qxwlwagnt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"resty.dev/v3"
)

// PreRequestMiddleware 自动签名中间件
func PreRequestMiddleware(debug bool, userName, appKey, appSecret string) resty.RequestMiddleware {
	return func(_ *resty.Client, r *resty.Request) error {
		// 1. 从 context 获取 bodyMap（作为唯一参数来源）
		params := make(map[string]string)
		if raw, ok := r.Context().Value(bodyMapKey).(map[string]any); ok && raw != nil {
			for k, v := range raw {
				params[k] = fmt.Sprint(v)
			}
		}

		// 2. 添加公共参数（始终存在）
		timestamp := time.Now().Format("20060102150405")
		params["appKey"] = appKey
		params["userName"] = userName
		params["timeStamp"] = timestamp

		// 3. 排序 + 签名（URL 编码）
		keys := make([]string, 0, len(params))
		for k := range params {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var sb strings.Builder
		for _, k := range keys {
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(params[k])))
		}
		sb.WriteString("appSecret=" + url.QueryEscape(appSecret))

		signStr := sb.String()
		hash := md5.Sum([]byte(signStr))
		sign := strings.ToUpper(hex.EncodeToString(hash[:]))

		// 4. 添加参数
		for k, v := range params {
			r.SetQueryParam(k, v)
		}
		r.SetQueryParam("sign", sign)

		if debug {
			fmt.Println("[签名字符串]", signStr)
			fmt.Println("[签名结果]", sign)
			fmt.Println("[最终 QueryParams]", r.QueryParams)
		}

		return nil
	}
}

// Ensure2xxResponseMiddleware 确保响应状态码为 2xx
func Ensure2xxResponseMiddleware(_ *resty.Client, resp *resty.Response) error {
	if !resp.IsSuccess() {
		return fmt.Errorf("请求失败: 状态码 %d, 响应: %s", resp.StatusCode(), resp.String())
	}
	return nil
}
