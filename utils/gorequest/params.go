package gorequest

import "sync"

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

// DeepGet 深度获取
func (p *Params) DeepGet() map[string]any {
	p.Lock()
	defer p.Unlock()

	targetMap := make(map[string]interface{})
	for key, value := range p.m {
		targetMap[key] = value
	}
	return targetMap
}

// DeepCopy 深度复制
func (p *Params) DeepCopy() map[string]any {
	p.Lock()
	defer p.Unlock()

	// 深度复制数据
	targetMap := make(map[string]any)
	for key, value := range p.m {
		targetMap[key] = value
	}

	// 清空原始数据
	p.m = make(map[string]any)

	return targetMap
}
