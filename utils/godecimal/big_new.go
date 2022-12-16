package godecimal

// NewString 从字符串创建
func NewString(s string) Decimal {
	d := New()
	d.floatValue.SetString(s)
	return d
}

// NewFloat 从浮点数创建
func NewFloat(f float64) Decimal {
	d := New()
	d.floatValue.SetFloat64(f)
	return d
}

// NewInt 从整数创建
func NewInt(i int64) Decimal {
	d := New()
	d.floatValue.SetInt64(i)
	return d
}

// NewUint 从无符合整数创建
func NewUint(i uint64) Decimal {
	d := New()
	d.floatValue.SetUint64(i)
	return d
}
