package gorequest

import (
	"net/url"
)

// OldHeaders 头部信息
type OldHeaders map[string]string

// NewOldHeaders 新建头部信息
func NewOldHeaders() OldHeaders {
	P := make(OldHeaders)
	return P
}

// NewNewOldHeadersWith 头部信息使用
func NewNewOldHeadersWith(oldHeaders ...OldHeaders) OldHeaders {
	p := make(OldHeaders)
	for _, v := range oldHeaders {
		p.SetHeaders(v)
	}
	return p
}

// Set 设置头部信息
func (h OldHeaders) Set(key, value string) {
	h[key] = value
}

// SetHeaders 批量设置头部信息
func (h OldHeaders) SetHeaders(OldHeaders OldHeaders) {
	for key, value := range OldHeaders {
		h[key] = value
	}
}

// GetQuery 获取头部信息
func (h OldHeaders) GetQuery() string {
	u := url.Values{}
	for k, v := range h {
		u.Set(k, v)
	}
	return u.Encode()
}

// DeepCopy 深度复制
func (h *OldHeaders) DeepCopy() map[string]string {
	targetMap := make(map[string]string)

	// 从原始复制到目标
	for key, value := range *h {
		targetMap[key] = value
	}

	// 重新申请一个新的map
	*h = map[string]string{}
	return targetMap
}
