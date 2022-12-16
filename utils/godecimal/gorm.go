package godecimal

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type DFloat struct {
	FValue float64
	Point  int
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (f DFloat) Value() (driver.Value, error) {
	return NewFloat(f.FValue).Float64PointAdaptive(f.Point), nil
}

// Scan 方法实现了 sql.Scanner 接口
func (f *DFloat) Scan(value interface{}) error {
	f1, _ := value.(float64)
	*f = DFloat{
		FValue: f1,
		Point:  pointLength(f1),
	}
	return nil
}

func pointLength(a any) int {
	tmp := strings.Split(fmt.Sprint(a), ".")
	if len(tmp) <= 1 {
		return 0
	}
	return len(tmp[1])
}
