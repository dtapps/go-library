package gorequest

// Params 参数
type Params map[string]interface{}

// NewParams 新建参数
func NewParams() Params {
	P := make(Params)
	return P
}

// NewParamsWith 参数使用
func NewParamsWith(params ...Params) Params {
	p := make(Params)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

// Set 设置参数
func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

// SetParams 批量设置参数
func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

// Get 获取参数
func (p Params) Get(key string) interface{} {
	return p[key]
}

// DeepCopy 深度复制
func (p *Params) DeepCopy() map[string]interface{} {
	targetMap := make(map[string]interface{})

	// 从原始复制到目标
	for key, value := range *p {
		targetMap[key] = value
	}

	// 重新申请一个新的map
	*p = map[string]interface{}{}
	return targetMap
}
