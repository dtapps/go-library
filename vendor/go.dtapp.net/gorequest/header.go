package gorequest

import (
	"net/url"
)

// Headers 头部信息
type Headers map[string]string

// NewHeaders 新建头部信息
func NewHeaders() Headers {
	P := make(Headers)
	return P
}

// NewNewHeadersWith 头部信息使用
func NewNewHeadersWith(headers ...Headers) Headers {
	p := make(Headers)
	for _, v := range headers {
		p.SetHeaders(v)
	}
	return p
}

// Set 设置头部信息
func (p Headers) Set(key, value string) {
	p[key] = value
}

// SetHeaders 批量设置头部信息
func (p Headers) SetHeaders(headers Headers) {
	for key, value := range headers {
		p[key] = value
	}
}

// GetQuery 获取头部信息
func (p Headers) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, v)
	}
	return u.Encode()
}

// DeepCopy 深度复制
func (p *Headers) DeepCopy() map[string]string {
	targetMap := make(map[string]string)

	// 从原始复制到目标
	for key, value := range *p {
		targetMap[key] = value
	}

	// 重新申请一个新的map
	*p = map[string]string{}
	return targetMap
}
