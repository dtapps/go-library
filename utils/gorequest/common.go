package gorequest

import (
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gostring"
	"log"
)

// deepCopy 用于深度复制数据
func deepCopy(data interface{}) interface{} {
	// 这里根据数据类型进行适当的深度复制操作，包括切片、映射等数据结构
	// 返回深度复制后的数据
	// 示例中只支持切片类型的深度复制，你可以扩展支持其他数据类型
	if srcSlice, ok := data.([]interface{}); ok {
		dstSlice := make([]interface{}, len(srcSlice))
		copy(dstSlice, srcSlice)
		srcSlice = srcSlice[:0]
		return dstSlice
	}

	// 如果没有特殊处理的数据类型，直接返回原始数据
	return data
}

// GetParamsString 获取参数字符串
func GetParamsString(src interface{}) string {
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return gostring.ToString(src)
	}
	data, err := gojson.Marshal(src)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
