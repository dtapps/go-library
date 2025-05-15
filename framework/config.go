package framework

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type Type string

const (
	Gin   Type = "gin"   // https://gin-gonic.com/zh-cn/
	Hertz Type = "hertz" // https://www.cloudwego.io/zh/docs/hertz/
	Echo  Type = "echo"  // https://echo.labstack.com/
	Fiber Type = "fiber" // https://docs.fiber.org.cn/
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

// IsEcho 是否使用 Echo
func (c *Context) IsEcho() bool {
	if useFramework == Echo && c.echoCtx != nil {
		return true
	}
	return false
}

// GetContext 获取上下文
func (c *Context) GetContext() context.Context {
	return c.ctx
}

// SetContext 设置上下文
func (c *Context) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// GetSafeContext 获取安全上下文
func (c *Context) GetSafeContext() context.Context {
	return context.WithoutCancel(c.ctx)
}

// GetGinContext 获取原始的 Gin 上下文
func (c *Context) GetGinContext() *gin.Context {
	if c.ginCtx == nil {
		return nil
	}
	return c.ginCtx
}

// GetHertzContext 获取原始的 Hertz 上下文
func (c *Context) GetHertzContext() *app.RequestContext {
	if c.hertzCtx == nil {
		return nil
	}
	return c.hertzCtx
}

// GetEchoContext 获取原始的 Echo 上下文
func (c *Context) GetEchoContext() echo.Context {
	if c.echoCtx == nil {
		return nil
	}
	return c.echoCtx
}

// 是否调试
var isDebug = false

// InitUseFramework 初始使用框架
func InitDebug(debug bool) {
	isDebug = debug
}
