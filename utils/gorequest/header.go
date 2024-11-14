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

// DeepCopy 深度复制
func (h *Headers) DeepCopy() map[string]string {
	h.Lock()
	defer h.Unlock()
	targetMap := make(map[string]string)
	for key, value := range h.m {
		targetMap[key] = value
	}
	return targetMap
}
