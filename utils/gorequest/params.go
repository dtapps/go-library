package gorequest

import "sync"

// Params 参数
type Params struct {
	sync.Mutex
	m map[string]interface{}
}

// NewParams 新建参数
func NewParams() *Params {
	return &Params{
		m: make(map[string]interface{}),
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
func (p *Params) Set(key string, value interface{}) {
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
func (p *Params) Get(key string) interface{} {
	p.Lock()
	defer p.Unlock()
	return p.m[key]
}

// DeepCopy 深度复制
func (p *Params) DeepCopy() map[string]interface{} {
	p.Lock()
	defer p.Unlock()
	targetMap := make(map[string]interface{})
	for key, value := range p.m {
		targetMap[key] = value
	}
	return targetMap
}
