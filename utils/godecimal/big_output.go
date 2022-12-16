package godecimal

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

// String 输出 string
func (d Decimal) String() string {
	return d.floatValue.String()
}

// Int64 输出 int64
func (d Decimal) Int64() int64 {
	i64, _ := d.floatValue.Int64()
	return i64
}

// Float64 输出 float64
func (d Decimal) Float64() float64 {
	rat, _ := new(big.Rat).SetString(d.String())
	f, _ := rat.Float64()
	return f
}

// MoneyFloat64 货币 float64
func (d Decimal) MoneyFloat64() float64 {
	rat, _ := new(big.Rat).SetString(d.floatValue.Text('f', 2))
	f, _ := rat.Float64()
	return f
}

// Float64Point 输出float64带小数点
func (d Decimal) Float64Point(p int) float64 {
	rat, _ := new(big.Rat).SetString(d.floatValue.Text('f', p))
	f, _ := rat.Float64()
	return f
}

// Float64PointAdaptive 输出float64带小数点(自适应)
func (d Decimal) Float64PointAdaptive(maxP int) float64 {
	f, _ := d.floatValue.Float64()
	if maxP > 0 {
		pL := d.pointLength(f)
		if pL > maxP {
			rat, _ := new(big.Rat).SetString(d.floatValue.Text('f', maxP))
			f2, _ := rat.Float64()
			return f2
		} else {
			return f
		}
	} else {
		return f
	}
}

func (Decimal) pointLength(a any) int {
	tmp := strings.Split(fmt.Sprint(a), ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}

// IsInteger 是否为整数
func (d Decimal) IsInteger(d2 float64) bool {
	if d2 > 0 {
		f3 := NewFloat(d.Float64()).QuoFloat(NewFloat(d2).Float64()).Float64()
		if f3 == math.Trunc(f3) {
			return true
		}
		return false
	}
	f3 := NewFloat(d.Float64()).Float64()
	if f3 == math.Trunc(f3) {
		return true
	}
	return false
}
