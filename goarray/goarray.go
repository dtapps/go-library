package goarray

func Grouping() {

}

// TurnString []string 转 string
func TurnString(ss []string) (s string) {
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

// RemoveDuplicateElement 去重
func RemoveDuplicateElement(ss []string) []string {
	result := make([]string, 0, len(ss))
	temp := map[string]struct{}{}
	for _, item := range ss {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
