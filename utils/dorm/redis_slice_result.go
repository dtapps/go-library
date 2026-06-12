package dorm

type SliceResult struct {
	Result []any
	Err    error
}

// NewSliceResult 构造函数
func NewSliceResult(result []any, err error) *SliceResult {
	return &SliceResult{Result: result, Err: err}
}

// Unwrap 空值情况下返回错误
func (r *SliceResult) Unwrap() []any {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

// UnwrapOr 空值情况下设置返回默认值
func (r *SliceResult) UnwrapOr(defaults []any) []any {
	if r.Err != nil {
		return defaults
	}
	return r.Result
}

func (r *SliceResult) Iter() *Iterator {
	return NewIterator(r.Result)
}
