package goorder

import (
	"fmt"
	"go.dtapp.net/library/utils/gorandom"
	"go.dtapp.net/library/utils/gotime"
)

// GenerateOrDate 生成订单号
func GenerateOrDate[T string | int | int8 | int16 | int32 | int64](customId T) string {
	return fmt.Sprintf("%v%s%s", customId, gotime.Current().SetFormat("200601021504"), gorandom.Numeric(3))
}
