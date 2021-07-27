package dredis

type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{Result: result, Err: err}
}

// Unwrap 空值情况下返回错误
func (r *StringResult) Unwrap() string {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Result
}

// UnwrapOr 空值情况下设置返回默认值
func (r *StringResult) UnwrapOr(defaults string) string {
	if r.Err != nil {
		return defaults
	}
	return r.Result
}

func (r *StringResult) UnwrapOrElse(f func() string) string {
	if r.Err != nil {
		return f()
	}
	return r.Result
}
