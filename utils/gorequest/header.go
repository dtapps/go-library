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
	h := &Headers{}
	return h
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
func (h *Headers) Set(key string, value any) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.m.Store(key, value)
}

// SetHeaders 批量设置头部信息
func (h *Headers) SetHeaders(headers *Headers) {
	headers.m.Range(func(key, value interface{}) bool {
		h.Set(key.(string), value.(string))
		return true
	})
}

// ToMap 返回 map[string]interface{}
func (h *Headers) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	h.m.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})

	return result
}

// ToMapAndReset 返回 map[string]interface{} 然后清空原始数据
func (h *Headers) ToMapAndReset() map[string]interface{} {
	result := make(map[string]interface{})

	h.m.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})

	h.Reset() // 清空原始数据

	return result
}

// HasData 判断是否有数据
func (h *Headers) HasData() bool {
	hasData := false

	h.m.Range(func(_, _ interface{}) bool {
		hasData = true
		// 返回 false 停止遍历
		return false
	})

	return hasData
}

// DeepCopy 深度复制
func (h *Headers) DeepCopy() *Headers {
	newHeaders := NewHeaders()

	h.m.Range(func(key, value interface{}) bool {
		// 深度复制数据并存储到新参数集合
		newValue := deepCopy(value)
		newHeaders.Set(key.(string), newValue)
		// 清空原始数据
		h.m.Delete(key)
		return true
	})

	h.Reset() // 清空原始数据

	return newHeaders
}

// Reset 清空结构体
func (h *Headers) Reset() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.m = sync.Map{}
}
