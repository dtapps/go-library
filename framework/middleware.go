package framework

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
)

// MiddlewareFunc 是兼容 Gin 和 Hertz 的中间件函数签名
type MiddlewareFunc func(ctx *Context)

// GinMiddleware 将统一中间件函数转换为 Gin 的中间件
func GinMiddleware(mw MiddlewareFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		wrapped := &Context{
			ctx:    c.Request.Context(), // 使用 Gin 提供的上下文
			ginCtx: c,                   // 保存 Gin 的上下文
		}
		mw(wrapped)

		// 如果已经中止请求，则不再执行后续中间件
		if !c.IsAborted() {
			c.Next()
		}
	}
}

// HertzMiddleware 将统一中间件函数转换为 Hertz 的中间件
func HertzMiddleware(mw MiddlewareFunc) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		wrapped := &Context{
			ctx:      c,   // 使用 Hertz 提供的上下文
			hertzCtx: ctx, // 保存 Hertz 的上下文
		}
		mw(wrapped)

		// Hertz 的中止机制略不同，不会自动跳过后续中间件，所以需手动判断
		if !ctx.IsAborted() {
			ctx.Next(c)
		}
	}
}

// Set 在上下文中设置键值对
func (c *Context) Set(key string, value any) {
	if c.ginCtx != nil {
		c.ginCtx.Set(key, value)
	}
	if c.hertzCtx != nil {
		c.hertzCtx.Set(key, value)
	}
}

// Get 从上下文中获取键对应的值（返回值和是否存在）
func (c *Context) Get(key string) (value any, exists bool) {
	if c.ginCtx != nil {
		return c.ginCtx.Get(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Get(key)
	}
	return
}

// GetString 从上下文中获取字符串类型的值
func (c *Context) GetString(key string) (s string) {
	if c.ginCtx != nil {
		return c.ginCtx.GetString(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetString(key)
	}
	return
}

// GetBool 从上下文中获取布尔类型的值
func (c *Context) GetBool(key string) (b bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetBool(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetBool(key)
	}
	return
}

// GetInt 从上下文中获取 int 类型的值
func (c *Context) GetInt(key string) (i int) {
	if c.ginCtx != nil {
		return c.ginCtx.GetInt(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetInt(key)
	}
	return
}

// GetInt64 从上下文中获取 int64 类型的值
func (c *Context) GetInt64(key string) (i64 int64) {
	if c.ginCtx != nil {
		return c.ginCtx.GetInt64(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetInt64(key)
	}
	return
}

// GetFloat64 从上下文中获取 float64 类型的值
func (c *Context) GetFloat64(key string) (f64 float64) {
	if c.ginCtx != nil {
		return c.ginCtx.GetFloat64(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetFloat64(key)
	}
	return
}
