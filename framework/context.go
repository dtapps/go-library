package framework

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
)

// Context 统一的 Context 封装
type Context struct {
	ginCtx   *gin.Context        // Gin 上下文
	hertzCtx *app.RequestContext // Hertz 上下文
}

// JSON 方法：统一返回 JSON 响应
func (c *Context) JSON(code int, obj any) {
	if c.ginCtx != nil {
		c.ginCtx.JSON(code, obj)
	} else if c.hertzCtx != nil {
		c.hertzCtx.JSON(code, obj)
	}
}

// Param 方法：统一获取路径参数
func (c *Context) Param(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Param(key)
	} else if c.hertzCtx != nil {
		return c.hertzCtx.Param(key)
	}
	return ""
}

// BindAndValidate 方法：统一绑定和验证请求数据
func (c *Context) BindAndValidate(obj any) error {
	if c.ginCtx != nil {
		// Gin 的绑定和验证
		if err := c.ginCtx.ShouldBind(obj); err != nil {
			return err
		}
	} else if c.hertzCtx != nil {
		// Hertz 的绑定和验证
		if err := c.hertzCtx.BindAndValidate(obj); err != nil {
			return err
		}
	}
	return nil
}

// GetGinContext 获取原始的 Gin 上下文
func (c *Context) GetGinContext() *gin.Context {
	return c.ginCtx
}

// GetHertzContext 获取原始的 Hertz 上下文
func (c *Context) GetHertzContext() *app.RequestContext {
	return c.hertzCtx
}

// HandlerFunc 统一的处理函数签名
type HandlerFunc func(ctx *Context)

// HertzHandler 封装 Hertz
func HertzHandler(handler HandlerFunc) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wrapperCtx := &Context{hertzCtx: ctx}
		handler(wrapperCtx)
	}
}

// GinHandler 封装 Gin
func GinHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		wrapperCtx := &Context{ginCtx: c}
		handler(wrapperCtx)
	}
}

// GetApis 通用的 GetApis 方法
func GetApis(ctx *Context) {
	name := ctx.Param("name")
	ctx.JSON(200, map[string]string{"message": "Hello, " + name})
}
