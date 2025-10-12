package gorequest

// 定义请求类型
const (
	httpParamsModeJson = "JSON"
	httpParamsModeXml  = "XML"
	httpParamsModeForm = "FORM"
)

// SortByKey 排序方式
type SortOrder int

const (
	Asc  SortOrder = iota // 升序
	Desc                  // 降序
)
