package godecimal

import (
	"fmt"
	"strconv"
)

// Decimal 四舍五入
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
