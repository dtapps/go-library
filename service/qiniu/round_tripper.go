package qiniu

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SkipRule 定义不需要签名的规则
type SkipRule struct {
	Host       string   // 匹配域名，如 "api.qiniu.com"
	PathPrefix string   // 匹配路径前缀，如 "/v1/"
	Methods    []string // 匹配的方法，如 ["POST", "PUT"]，为空则匹配所有
}

type SignTransport struct {
	Transport http.RoundTripper // 被拦截的 Transport
	SkipRules []SkipRule        // 过滤规则列表
	AccessKey string            // AccessKey
	SecretKey string            // SecretKey
}

func NewSignTransport(base http.RoundTripper, accessKey string, secretKey string) *SignTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &SignTransport{
		Transport: base,
		SkipRules: []SkipRule{},
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

// baseTransport 获取基础 Transport
func (t *SignTransport) baseTransport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// 配置 SkipRules 过滤规则列表
func (t *SignTransport) SetSkipRules(rules []SkipRule) {
	t.SkipRules = rules
}

// Clone 复制一个新的 SignTransport
func (t *SignTransport) Clone() *SignTransport {
	return &SignTransport{
		Transport: t.baseTransport(), // 被拦截的 Transport
		SkipRules: t.SkipRules,       // 过滤规则列表
		AccessKey: t.AccessKey,       // AccessKey
		SecretKey: t.SecretKey,       // SecretKey
	}
}

// Middleware 返回一个 http.RoundTripper 方法
func (t *SignTransport) Middleware() func(http.RoundTripper) http.RoundTripper {
	return func(next http.RoundTripper) http.RoundTripper {
		if next == nil {
			next = t.baseTransport()
		}
		return &SignTransport{
			Transport: next,        // 被拦截的 Transport
			SkipRules: t.SkipRules, // 过滤规则列表
			AccessKey: t.AccessKey, // AccessKey
			SecretKey: t.SecretKey, // SecretKey
		}
	}
}

// Instance 返回一个 http.RoundTripper 实例
func (t *SignTransport) Instance(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = t.baseTransport()
	}
	return &SignTransport{
		Transport: next,        // 被拦截的 Transport
		SkipRules: t.SkipRules, // 过滤规则列表
		AccessKey: t.AccessKey, // AccessKey
		SecretKey: t.SecretKey, // SecretKey
	}
}

// RoundTrip 实现了 http.RoundTripper 接口
func (t *SignTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	// 1. 如果命中了过滤规则，直接调用底层 Transport，跳过签名
	if t.shouldSkip(req) {
		return t.baseTransport().RoundTrip(req)
	}

	// 2. 待签名字符串构造
	signingStr := req.URL.Path
	if req.URL.RawQuery != "" {
		signingStr += "?" + req.URL.RawQuery
	}
	signingStr += "\n"

	// 3. 处理 Body
	// 只有在满足以下条件时才读取 Body 加入签名：
	// - 不是 GET 请求
	// - Content-Type 为 application/x-www-form-urlencoded
	ct := req.Header.Get("Content-Type")
	if req.Method != http.MethodGet && ct == "application/x-www-form-urlencoded" {
		if req.Body != nil {
			bodyBytes, err := io.ReadAll(req.Body)
			if err == nil {
				req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				signingStr += string(bodyBytes)
			}
		}
	}

	// 4. 计算 HMAC-SHA1
	h := hmac.New(sha1.New, []byte(t.SecretKey))
	h.Write([]byte(signingStr))

	// 5. 使用 URLSafe Base64
	sign := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// 6. 格式: QBox <AccessKey>:<Sign>
	token := fmt.Sprintf("QBox %s:%s", t.AccessKey, sign)
	req.Header.Set("Authorization", token)

	return t.baseTransport().RoundTrip(req)
}

// shouldSkip 判断是否需要跳过签名
func (t *SignTransport) shouldSkip(req *http.Request) bool {

	for _, rule := range t.SkipRules {
		// 如果规则指定了 Host 且不匹配，继续看下一条
		if rule.Host != "" && rule.Host != req.URL.Host {
			continue
		}
		// 如果规则指定了前缀且不匹配，继续
		if rule.PathPrefix != "" && !strings.HasPrefix(req.URL.Path, rule.PathPrefix) {
			continue
		}
		// 如果规则指定了方法且不匹配，继续
		if len(rule.Methods) > 0 {
			matchedMethod := false
			for _, m := range rule.Methods {
				if strings.ToUpper(m) == req.Method {
					matchedMethod = true
					break
				}
			}
			if !matchedMethod {
				continue
			}
		}

		// 只要有一条规则完全命中，就代表需要“过滤（不签名）”
		return true
	}
	return false
}
