package gorequest

// OldParams 参数
type OldParams map[string]any

// NewOldParams 新建参数
func NewOldParams() OldParams {
	P := make(OldParams)
	return P
}

// NewOldParamsWith 参数使用
func NewOldParamsWith(oldParams ...OldParams) OldParams {
	p := make(OldParams)
	for _, v := range oldParams {
		p.SetParams(v)
	}
	return p
}

// Set 设置参数
func (p OldParams) Set(key string, value interface{}) {
	p[key] = value
}

// SetParams 批量设置参数
func (p OldParams) SetParams(OldParams OldParams) {
	for key, value := range OldParams {
		p[key] = value
	}
}

// Get 获取参数
func (p OldParams) Get(key string) interface{} {
	return p[key]
}

// DeepCopy 深度复制
func (p *OldParams) DeepCopy() map[string]any {
	targetMap := make(map[string]any)

	// 从原始复制到目标
	for key, value := range *p {
		targetMap[key] = value
	}

	// 重新申请一个新的map
	*p = map[string]any{}
	return targetMap
}
