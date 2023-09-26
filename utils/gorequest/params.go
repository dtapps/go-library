package gorequest

import (
	"sync"
)

// Params 参数
type Params struct {
	mu sync.Mutex // 用于保护 map 的互斥锁
	m  sync.Map   // 使用 sync.Map 存储参数
}

// NewParams 新建参数
func NewParams() *Params {
	p := &Params{}
	return p
}

// NewParamsWith 参数使用
func NewParamsWith(params ...*Params) *Params {
	p := NewParams()

	for _, param := range params {
		p.SetParams(param)
	}

	return p
}

// Set 设置参数
func (p *Params) Set(key string, value any) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.m.Store(key, value)
}

// SetParams 批量设置参数
func (p *Params) SetParams(params *Params) {
	params.m.Range(func(key, value interface{}) bool {
		p.Set(key.(string), value)
		return true
	})
}

// Get 获取参数
func (p *Params) Get(key string) interface{} {
	val, _ := p.m.Load(key)
	return val
}

// ToMap 返回 map[string]interface{}
func (p *Params) ToMap() map[string]interface{} {
	result := make(map[string]interface{})

	p.m.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})

	return result
}

// ToMapAndReset 返回 map[string]interface{} 然后清空原始数据
func (p *Params) ToMapAndReset() map[string]interface{} {
	result := make(map[string]interface{})

	p.m.Range(func(key, value interface{}) bool {
		result[key.(string)] = value
		return true
	})

	p.Reset() // 清空原始数据

	return result
}

// HasData 判断是否有数据
func (p *Params) HasData() bool {
	hasData := false

	p.m.Range(func(_, _ interface{}) bool {
		hasData = true
		// 返回 false 停止遍历
		return false
	})

	return hasData
}

// DeepCopy 深度复制
func (p *Params) DeepCopy() *Params {
	newParams := NewParams()

	p.m.Range(func(key, value interface{}) bool {
		// 深度复制数据并存储到新参数集合
		newValue := deepCopy(value)
		newParams.Set(key.(string), newValue)
		// 清空原始数据
		p.m.Delete(key)
		return true
	})

	p.Reset() // 清空原始数据

	return newParams
}

// Reset 清空结构体
func (p *Params) Reset() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.m = sync.Map{}
}
