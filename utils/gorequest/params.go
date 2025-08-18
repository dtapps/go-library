package gorequest

import (
	"fmt"
	"sort"
	"sync"
)

// Params 参数
type Params struct {
	sync.Mutex
	m map[string]any
}

// NewParams 新建参数
func NewParams() *Params {
	return &Params{
		m: make(map[string]any),
	}
}

// NewParamsWith 参数使用
func NewParamsWith(params ...*Params) *Params {
	p := NewParams()
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

// Set 设置参数
func (p *Params) Set(key string, value any) {
	p.Lock()
	defer p.Unlock()

	p.m[key] = value
}

// SetParams 批量设置参数
func (p *Params) SetParams(params *Params) {
	p.Lock()
	defer p.Unlock()
	for key, value := range params.m {

		p.m[key] = value
	}
}

// Get 获取参数
func (p *Params) Get(key string) any {
	p.Lock()
	defer p.Unlock()
	return p.m[key]
}

// Deprecated: 请使用 DeepGetString / DeepGetAny 方法代替。
// DeepGet 方法已被弃用，因为它返回的类型不够明确，可能导致类型转换问题。
func (p *Params) DeepGet() map[string]any {
	p.Lock()
	defer p.Unlock()

	targetMap := make(map[string]any)
	for key, value := range p.m {
		targetMap[key] = value
	}
	return targetMap
}

// DeepCopy 深度复制
func (p *Params) DeepCopy() *Params {
	p.Lock()
	defer p.Unlock()

	// 深度复制数据
	targetParam := NewParams()
	for key, value := range p.m {
		targetParam.Set(key, value)
	}

	// 清空原始数据
	p.m = make(map[string]any)

	return targetParam
}

// DeepGetString 深度获取
func (p *Params) DeepGetString() map[string]string {
	p.Lock()
	defer p.Unlock()

	targetMap := make(map[string]string)
	for key, value := range p.m {
		targetMap[key] = fmt.Sprintf("%v", value)
	}
	return targetMap
}

// DeepGetAny 深度获取
func (p *Params) DeepGetAny() map[string]any {
	p.Lock()
	defer p.Unlock()

	targetMap := make(map[string]any)
	for key, value := range p.m {
		targetMap[key] = value
	}
	return targetMap
}

// SortByKeyString 按 key 排序，返回 map[string]string
func (p *Params) SortByKeyString(order SortOrder) map[string]string {
	p.Lock()
	defer p.Unlock()

	// 收集 key 并排序
	keys := make([]string, 0, len(p.m))
	for k := range p.m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if order == Asc {
			return keys[i] < keys[j]
		}
		return keys[i] > keys[j]
	})

	// 生成新的 map[string]string
	ordered := make(map[string]string, len(p.m))
	for _, k := range keys {
		ordered[k] = fmt.Sprintf("%v", p.m[k])
	}
	return ordered
}

// SortByKeyAny 按 key 排序，返回 map[string]any
func (p *Params) SortByKeyAny(order SortOrder) map[string]any {
	p.Lock()
	defer p.Unlock()

	// 收集 key 并排序
	keys := make([]string, 0, len(p.m))
	for k := range p.m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if order == Asc {
			return keys[i] < keys[j]
		}
		return keys[i] > keys[j]
	})

	// 生成新的 map[string]any
	ordered := make(map[string]any, len(p.m))
	for _, k := range keys {
		ordered[k] = p.m[k]
	}
	return ordered
}
