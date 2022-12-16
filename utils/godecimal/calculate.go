package godecimal

// Float64Add 加 (f1+f2)
func Float64Add(f1, f2 float64) float64 {
	return NewFloat(f1).Add(NewFloat(f2)).Float64()
}

// Float64Sub 减 (f1-f2)
func Float64Sub(f1, f2 float64) float64 {
	return NewFloat(f1).Sub(NewFloat(f2)).Float64()
}

// Float64Mul 乘 (f1*f2)
func Float64Mul(f1, f2 float64) float64 {
	return NewFloat(f1).Mul(NewFloat(f2)).Float64()
}

// Float64Quo 除 (f1/f2)
func Float64Quo(f1, f2 float64) float64 {
	return NewFloat(f1).Quo(NewFloat(f2)).Float64()
}
