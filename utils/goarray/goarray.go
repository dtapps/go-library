package goarray

func Grouping() {

}

// TurnString 字符串切片 转 字符串
func TurnString[T string](ss []T) (s T) {
	sl := len(ss)
	for k, v := range ss {
		if k+1 == sl {
			s = s + v
		} else {
			s = s + v + ","
		}
	}
	return s
}

// SplitSliceIntoChunks 将一个字符串切片分割成多个子切片，每个子切片的长度不超过指定的最大长度。
func SplitSliceIntoChunks[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

// RemoveDuplicateElement 去重
func RemoveDuplicateElement[T string | int | int16 | int32](ss []T) []T {
	result := make([]T, 0, len(ss))
	temp := map[T]struct{}{}
	for _, item := range ss {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
