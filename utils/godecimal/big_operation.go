package godecimal

// Add 加 (d+d2)
func (d Decimal) Add(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Add(d.floatValue, d2.floatValue)
	return mul
}

// AddFloat 加
func (d Decimal) AddFloat(d2 float64) Decimal {
	mul := New()
	mul.floatValue.Add(d.floatValue, NewFloat(d2).floatValue)
	return mul
}

// AddInt 加
func (d Decimal) AddInt(d2 int64) Decimal {
	mul := New()
	mul.floatValue.Add(d.floatValue, NewInt(d2).floatValue)
	return mul
}

// AddString 加
func (d Decimal) AddString(d2 string) Decimal {
	mul := New()
	mul.floatValue.Add(d.floatValue, NewString(d2).floatValue)
	return mul
}

// Sub 减 (d-d2)
func (d Decimal) Sub(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Sub(d.floatValue, d2.floatValue)
	return mul
}

// SubFloat 减
func (d Decimal) SubFloat(d2 float64) Decimal {
	mul := New()
	mul.floatValue.Sub(d.floatValue, NewFloat(d2).floatValue)
	return mul
}

// SubInt 减
func (d Decimal) SubInt(d2 int64) Decimal {
	mul := New()
	mul.floatValue.Sub(d.floatValue, NewInt(d2).floatValue)
	return mul
}

// SubString 减
func (d Decimal) SubString(d2 string) Decimal {
	mul := New()
	mul.floatValue.Sub(d.floatValue, NewString(d2).floatValue)
	return mul
}

// Mul 乘 (d*d2)
func (d Decimal) Mul(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Mul(d.floatValue, d2.floatValue)
	return mul
}

// MulFloat 乘
func (d Decimal) MulFloat(d2 float64) Decimal {
	mul := New()
	mul.floatValue.Mul(d.floatValue, NewFloat(d2).floatValue)
	return mul
}

// MulInt 乘
func (d Decimal) MulInt(d2 int64) Decimal {
	mul := New()
	mul.floatValue.Mul(d.floatValue, NewInt(d2).floatValue)
	return mul
}

// MulString 乘
func (d Decimal) MulString(d2 string) Decimal {
	mul := New()
	mul.floatValue.Mul(d.floatValue, NewString(d2).floatValue)
	return mul
}

// Quo 除 (d/d2)
func (d Decimal) Quo(d2 Decimal) Decimal {
	mul := New()
	mul.floatValue.Quo(d.floatValue, d2.floatValue)
	return mul
}

// QuoFloat 除
func (d Decimal) QuoFloat(d2 float64) Decimal {
	mul := New()
	mul.floatValue.Quo(d.floatValue, NewFloat(d2).floatValue)
	return mul
}

// QuoInt 除
func (d Decimal) QuoInt(d2 int64) Decimal {
	mul := New()
	mul.floatValue.Quo(d.floatValue, NewInt(d2).floatValue)
	return mul
}

// QuoString 除
func (d Decimal) QuoString(d2 string) Decimal {
	mul := New()
	mul.floatValue.Quo(d.floatValue, NewString(d2).floatValue)
	return mul
}
