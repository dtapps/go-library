package framework

import "fmt"

type Type string

const (
	Gin   Type = "gin"
	Hertz Type = "hertz"
)

// 使用框架
var useFramework Type = ""

// InitUseFramework 初始使用框架
func InitUseFramework(use Type) error {
	if use != Gin && use != Hertz {
		return fmt.Errorf("不支持的框架类型: %s", use)
	}
	useFramework = use
	return nil
}

// IsGin 是否使用 Gin
func (c *Context) IsGin() bool {
	if useFramework == Gin && c.ginCtx != nil {
		return true
	}
	return false
}

// IsHertz 是否使用 Hertz
func (c *Context) IsHertz() bool {
	if useFramework == Hertz && c.hertzCtx != nil {
		return true
	}
	return false
}
