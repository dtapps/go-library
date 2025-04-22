package gorequest

import (
	"net/url"
	"sync"
)

// Headers 头部信息
type Headers struct {
	sync.Mutex
	m map[string]string
}

// NewHeaders 新建头部信息
func NewHeaders() *Headers {
	return &Headers{
		m: make(map[string]string),
	}
}

// NewNewHeadersWith 头部信息使用
func NewNewHeadersWith(headers ...*Headers) *Headers {
	h := NewHeaders()
	for _, v := range headers {
		h.SetHeaders(v)
	}
	return h
}

// Set 设置头部信息
func (h *Headers) Set(key string, value string) {
	h.Lock()
	defer h.Unlock()
	h.m[key] = value
}

// SetHeaders 批量设置头部信息
func (h *Headers) SetHeaders(headers *Headers) {
	h.Lock()
	defer h.Unlock()
	for key, value := range headers.m {
		h.m[key] = value
	}
}

// Get 获取头部信息
func (h *Headers) Get(key string) string {
	h.Lock()
	defer h.Unlock()
	return h.m[key]
}

// GetQuery 获取头部信息
func (h *Headers) GetQuery() string {
	u := url.Values{}
	for k, v := range h.m {
		u.Set(k, v)
	}
	return u.Encode()
}

// Deprecated: 请使用 DeepGetString 方法代替。
// DeepGet 方法已被弃用，因为它返回的类型不够明确
func (h *Headers) DeepGet() map[string]string {
	h.Lock()
	defer h.Unlock()

	targetMap := make(map[string]string)
	for key, value := range h.m {
		targetMap[key] = value
	}
	return targetMap
}

// DeepGetString 深度获取
func (h *Headers) DeepGetString() map[string]string {
	h.Lock()
	defer h.Unlock()

	targetMap := make(map[string]string)
	for key, value := range h.m {
		targetMap[key] = value
	}
	return targetMap
}

// DeepCopy 深度复制
func (h *Headers) DeepCopy() *Headers {
	h.Lock()
	defer h.Unlock()

	// 深度复制数据
	targetHeader := NewHeaders()
	for key, value := range h.m {
		targetHeader.Set(key, value)
	}

	// 清空原始数据
	h.m = make(map[string]string)

	return targetHeader
}
