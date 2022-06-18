package dorm

type SimpleResult struct {
	Result interface{}
	Err    error
}

// NewSimpleResult 构造函数
func NewSimpleResult(result interface{}, err error) *SimpleResult {
	return &SimpleResult{Result: result, Err: err}
}

// Unwrap 空值情况下返回错误
func (r *SimpleResult) Unwrap() interface{} {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

// UnwrapOr 空值情况下设置返回默认值
func (r *SimpleResult) UnwrapOr(defaults interface{}) interface{} {
	if r.Err != nil {
		return defaults
	}
	return r.Result
}

// UnwrapOrElse 空值情况下设置返回其他
func (r *SimpleResult) UnwrapOrElse(f func() interface{}) interface{} {
	if r.Err != nil {
		return f()
	}
	return r.Result
}
