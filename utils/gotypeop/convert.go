package gotypeop

import (
	"log/slog"
	"math"
)

// IntSliceToInt64Slice 将 []int 转换为 []int64
func IntSliceToInt64Slice(is []int) []int64 {
	if is == nil {
		return nil
	}
	result := make([]int64, len(is))
	for i, v := range is {
		result[i] = int64(v)
	}
	return result
}

// Int64SliceToIntSlice 将 []int64 转换为 []int
func Int64SliceToIntSlice(is []int64) []int {
	if is == nil {
		return nil
	}
	result := make([]int, len(is))
	for i, v := range is {
		// 检查是否超出 int 范围（如 math.MaxInt）
		if v > math.MaxInt || v < math.MinInt {
			slog.Error("[Int64SliceToIntSlice] 超出 int 范围了", slog.Int64("value", v))
		}
		result[i] = int(v)
	}
	return result
}
