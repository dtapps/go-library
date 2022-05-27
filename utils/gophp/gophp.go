package gophp

import (
	"go.dtapp.net/library/utils/gophp/serialize"
	"strconv"
)

// Serialize 序列
func Serialize(value any) ([]byte, error) {
	return serialize.Marshal(value)
}

// Unserialize 反序列
func Unserialize(data []byte) (any, error) {
	return serialize.UnMarshal(data)
}

func BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// ArrayColumn array_column()
func ArrayColumn(input map[string]map[string]any, columnKey string) []any {
	columns := make([]any, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}
