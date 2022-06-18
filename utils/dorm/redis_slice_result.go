package dorm

type SliceResult struct {
	Result []interface{}
	Err    error
}

// NewSliceResult 构造函数
func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

// Unwrap 空值情况下返回错误
func (r *SliceResult) Unwrap() []interface{} {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

// UnwrapOr 空值情况下设置返回默认值
func (r *SliceResult) UnwrapOr(defaults []interface{}) []interface{} {
	if r.Err != nil {
		return defaults
	}
	return r.Result
}

func (r *SliceResult) Iter() *Iterator {
	return NewIterator(r.Result)
}
