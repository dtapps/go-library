package gotypeop

import "math"

// 检查一个 int64 值是否在 int 范围内
func safeInt64ToInt(v int64) (int, bool) {
	if v > math.MaxInt || v < math.MinInt {
		return 0, false // 超出范围
	}
	return int(v), true
}
