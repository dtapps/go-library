package gorequest

import (
	"sync"
)

// Headers 头部信息
type Headers struct {
	mu sync.Mutex // 用于保护 map 的互斥锁
	m  sync.Map   // 使用 sync.Map 存储参数
}

// NewHeaders 新建头部信息
func NewHeaders() *Headers {
	P := &Headers{}
	return P
}

// NewNewHeadersWith 头部信息使用
func NewNewHeadersWith(headers ...*Headers) *Headers {
	p := NewHeaders()
	for _, v := range headers {
		p.SetHeaders(v)
	}
	return p
}

// Set 设置头部信息
func (p *Headers) Set(key string, value any) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.m.Store(key, value)
}

// SetHeaders 批量设置头部信息
func (p *Headers) SetHeaders(headers *Headers) {
	headers.m.Range(func(key, value interface{}) bool {
		p.Set(key.(string), value.(string))
		return true
	})
}

// ToMap 返回 map[string]interface{}
func (p *Headers) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	p.m.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})

	return result
}

// HasData 判断是否有数据
func (p *Headers) HasData() bool {
	hasData := false

	p.m.Range(func(_, _ interface{}) bool {
		hasData = true
		// 返回 false 停止遍历
		return false
	})

	return hasData
}

// DeepCopy 深度复制
func (p *Headers) DeepCopy() *Headers {
	newHeaders := NewHeaders()

	p.m.Range(func(key, value interface{}) bool {
		newHeaders.Set(key.(string), value.(string))
		return true
	})

	return newHeaders
}
