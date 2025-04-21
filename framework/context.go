package framework

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// Context 统一的 Context 封装
type Context struct {
	ctx      context.Context     // 统一的上下文
	ginCtx   *gin.Context        // Gin 上下文
	hertzCtx *app.RequestContext // Hertz 上下文
	echoCtx  echo.Context        // Echo 上下文
}

// Next 继续处理请求
func (c *Context) Next() {
	if c.IsGin() {
		c.ginCtx.Next()
	}
	if c.IsHertz() {
		c.hertzCtx.Next(c.ctx)
	}
	if c.IsEcho() {

	}
}

// Abort 中止请求
func (c *Context) Abort() {
	if c.IsGin() {
		c.ginCtx.Abort()
	}
	if c.IsHertz() {
		c.hertzCtx.Abort()
	}
	if c.IsEcho() {
	}
}

// AbortWithStatus 中止请求并设置状态码
func (c *Context) AbortWithStatus(code int) {
	if c.IsGin() {
		c.ginCtx.AbortWithStatus(code)
	}
	if c.IsHertz() {
		c.hertzCtx.AbortWithStatus(code)
	}
	if c.IsEcho() {
	}
}

// AbortWithStatusJSON 中止请求并设置状态码和响应体
func (c *Context) AbortWithStatusJSON(code int, jsonObj any) {
	if c.IsGin() {
		c.ginCtx.AbortWithStatusJSON(code, jsonObj)
	}
	if c.IsHertz() {
		c.hertzCtx.AbortWithStatusJSON(code, jsonObj)
	}
	if c.IsEcho() {
	}
}

// JSON 方法：统一返回 JSON 响应
func (c *Context) JSON(code int, obj any) {
	if c.IsGin() {
		c.ginCtx.JSON(code, obj)
	}
	if c.IsHertz() {
		c.hertzCtx.JSON(code, obj)
	}
	if c.IsEcho() {
	}
}

// String 方法：统一返回 JSON 响应
func (c *Context) String(code int, format string, values ...any) {
	if c.IsGin() {
		c.ginCtx.String(code, format, values)
	}
	if c.IsHertz() {
		c.hertzCtx.String(code, format, values)
	}
	if c.IsEcho() {
	}
}

func (c *Context) PostForm(key string) string {
	if c.IsGin() {
		return c.ginCtx.PostForm(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.PostForm(key)
	}
	if c.IsEcho() {
	}
	return ""
}

func (c *Context) DefaultPostForm(key, defaultValue string) string {
	if c.IsGin() {
		return c.ginCtx.DefaultPostForm(key, defaultValue)
	}
	if c.IsHertz() {
		return c.hertzCtx.DefaultPostForm(key, defaultValue)
	}
	if c.IsEcho() {
	}
	return ""
}

func (c *Context) PostFormArray(key string) (values []string) {
	if c.IsGin() {
		return c.ginCtx.PostFormArray(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.PostFormArray(key)
	}
	if c.IsEcho() {
	}
	return
}

func (c *Context) GetPostForm(key string) (string, bool) {
	if c.IsGin() {
		return c.ginCtx.GetPostForm(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.GetPostForm(key)
	}
	if c.IsEcho() {
	}
	return "", false
}

func (c *Context) GetPostFormArray(key string) (values []string, ok bool) {
	if c.IsGin() {
		return c.ginCtx.GetPostFormArray(key)
	}
	if c.IsHertz() {
		return c.hertzCtx.GetPostFormArray(key)
	}
	if c.IsEcho() {
	}
	return
}

// HandlerFunc 统一的处理函数签名
type HandlerFunc func(ctx *Context)

// GinHandler 封装 Gin
func GinHandler(handler HandlerFunc) gin.HandlerFunc {
	if useFramework != Gin {
		return nil
	}
	return func(c *gin.Context) {
		wrapperCtx := &Context{
			ctx:    c.Request.Context(), // 使用 Gin 提供的上下文
			ginCtx: c,                   // 保存 Gin 的上下文
		}
		handler(wrapperCtx)
	}
}

// HertzHandler 封装 Hertz
func HertzHandler(handler HandlerFunc) app.HandlerFunc {
	if useFramework != Hertz {
		return nil
	}
	return func(c context.Context, ctx *app.RequestContext) {
		wrapperCtx := &Context{
			ctx:      c,   // 使用 Hertz 提供的上下文
			hertzCtx: ctx, // 保存 Hertz 的上下文
		}
		handler(wrapperCtx)
	}
}

// EchoHandler 封装 Echo
func EchoHandler(handler HandlerFunc) echo.HandlerFunc {
	if useFramework != Echo {
		return nil
	}
	return func(c echo.Context) error {
		wrapperCtx := &Context{
			ctx:     c.Request().Context(), // 使用 Echo 提供的上下文
			echoCtx: c,                     // 保存 Echo 的上下文
		}
		handler(wrapperCtx)
		return nil
	}
}
