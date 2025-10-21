package qxwlwagnt

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"resty.dev/v3"
)

// signMD5 返回小写 md5
func signMD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// PreRequestMiddleware 自动签名中间件
func PreRequestMiddleware(debug bool, userName, appKey, appSecret string) resty.RequestMiddleware {
	return func(_ *resty.Client, r *resty.Request) error {
		method := strings.ToUpper(r.RawRequest.Method)

		// 1. 统一收集 query 和 body 参数
		allParams := map[string]string{}

		// 收集 query 参数
		for k, v := range r.QueryParams {
			if len(v) > 0 {
				allParams[k] = v[0]
			}
		}

		// 收集 body 参数
		if r.Body != nil {
			switch body := r.Body.(type) {
			case map[string]any:
				for k, v := range body {
					allParams[k] = fmt.Sprint(v)
				}
			case map[string]string:
				for k, v := range body {
					allParams[k] = v
				}
			default:
				// 尝试解析 JSON
				b, _ := json.Marshal(r.Body)
				var m map[string]any
				if err := json.Unmarshal(b, &m); err == nil {
					for k, v := range m {
						allParams[k] = fmt.Sprint(v)
					}
				}
			}
		}

		// 2. 公共参数
		allParams["appKey"] = appKey
		allParams["userName"] = userName

		// 3. 按 key 排序
		keys := make([]string, 0, len(allParams))
		for k := range allParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		// 4. 拼接字符串
		var sb strings.Builder
		for _, k := range keys {
			sb.WriteString(fmt.Sprintf("%s=%s&", k, allParams[k]))
		}
		sb.WriteString("appSecret=" + appSecret)

		// 5. 计算 MD5 签名
		signStr := sb.String()
		hash := md5.Sum([]byte(signStr))
		sign := strings.ToUpper(hex.EncodeToString(hash[:]))

		// 6. 设置签名参数
		if method == "GET" {
			r.SetQueryParam("sign", sign)
			r.SetQueryParam("appKey", appKey)
			r.SetQueryParam("userName", userName)
		} else { // POST、PUT等
			bodyMap := map[string]any{}
			if r.Body != nil {
				b, _ := json.Marshal(r.Body)
				_ = json.Unmarshal(b, &bodyMap)
			}
			bodyMap["sign"] = sign
			bodyMap["appKey"] = appKey
			bodyMap["userName"] = userName
			r.SetBody(bodyMap)
		}

		if debug {
			fmt.Println("[签名字符串]", signStr)
			fmt.Println("[签名结果]", sign)
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
