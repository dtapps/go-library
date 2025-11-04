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

// GetString 安全获取字符串值
// 如果key不存在或类型错误，返回默认值
func (p *Params) GetString(key string, defaultValue ...string) string {
	val := p.Get(key)
	if val == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	
	if str, ok := val.(string); ok {
		return str
	}
	
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetStringOK 安全获取字符串值并返回是否成功
func (p *Params) GetStringOK(key string) (string, bool) {
	val := p.Get(key)
	if val == nil {
		return "", false
	}
	
	str, ok := val.(string)
	return str, ok
}

// GetInt 安全获取整数值
// 如果key不存在或类型错误，返回默认值
func (p *Params) GetInt(key string, defaultValue ...int) int {
	val := p.Get(key)
	if val == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	
	switch v := val.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case int32:
		return int(v)
	case float64:
		return int(v)
	case float32:
		return int(v)
	}
	
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetIntOK 安全获取整数值并返回是否成功
func (p *Params) GetIntOK(key string) (int, bool) {
	val := p.Get(key)
	if val == nil {
		return 0, false
	}
	
	switch v := val.(type) {
	case int:
		return v, true
	case int64:
		return int(v), true
	case int32:
		return int(v), true
	case float64:
		return int(v), true
	case float32:
		return int(v), true
	}
	
	return 0, false
}

// GetBool 安全获取布尔值
// 如果key不存在或类型错误，返回默认值
func (p *Params) GetBool(key string, defaultValue ...bool) bool {
	val := p.Get(key)
	if val == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	
	if b, ok := val.(bool); ok {
		return b
	}
	
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

// GetBoolOK 安全获取布尔值并返回是否成功
func (p *Params) GetBoolOK(key string) (bool, bool) {
	val := p.Get(key)
	if val == nil {
		return false, false
	}
	
	b, ok := val.(bool)
	return b, ok
}

// GetFloat 安全获取浮点数值
// 如果key不存在或类型错误，返回默认值
func (p *Params) GetFloat(key string, defaultValue ...float64) float64 {
	val := p.Get(key)
	if val == nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	
	switch v := val.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	}
	
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// GetFloatOK 安全获取浮点数值并返回是否成功
func (p *Params) GetFloatOK(key string) (float64, bool) {
	val := p.Get(key)
	if val == nil {
		return 0, false
	}
	
	switch v := val.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case int32:
		return float64(v), true
	}
	
	return 0, false
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
