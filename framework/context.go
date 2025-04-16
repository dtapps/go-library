package framework

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
)

// Context 统一的 Context 封装
type Context struct {
	Ctx      context.Context     // 统一的上下文
	ginCtx   *gin.Context        // Gin 上下文
	hertzCtx *app.RequestContext // Hertz 上下文
}

// Next 继续处理请求
func (c *Context) Next() {
	if c.ginCtx != nil {
		c.ginCtx.Next()
	}
	if c.hertzCtx != nil {
		c.hertzCtx.Next(c.Ctx)
	}
}

// Abort 中止请求
func (c *Context) Abort() {
	if c.ginCtx != nil {
		c.ginCtx.Abort()
	}
	if c.hertzCtx != nil {
		c.hertzCtx.Abort()
	}
}

// AbortWithStatus 中止请求并设置状态码
func (c *Context) AbortWithStatus(code int) {
	if c.ginCtx != nil {
		c.ginCtx.AbortWithStatus(code)
	}
	if c.hertzCtx != nil {
		c.hertzCtx.AbortWithStatus(code)
	}
}

// AbortWithStatusJSON 中止请求并设置状态码和响应体
func (c *Context) AbortWithStatusJSON(code int, jsonObj any) {
	if c.ginCtx != nil {
		c.ginCtx.AbortWithStatusJSON(code, jsonObj)
	}
	if c.hertzCtx != nil {
		c.hertzCtx.AbortWithStatusJSON(code, jsonObj)
	}
}

// JSON 方法：统一返回 JSON 响应
func (c *Context) JSON(code int, obj any) {
	if c.ginCtx != nil {
		c.ginCtx.JSON(code, obj)
	}
	if c.hertzCtx != nil {
		c.hertzCtx.JSON(code, obj)
	}
}

// String 方法：统一返回 JSON 响应
func (c *Context) String(code int, format string, values ...any) {
	if c.ginCtx != nil {
		c.ginCtx.String(code, format, values)
	}
	if c.hertzCtx != nil {
		c.hertzCtx.String(code, format, values)
	}
}

// Param 获取路径参数
func (c *Context) Param(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Param(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Param(key)
	}
	return ""
}

func (c *Context) Query(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Query(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.Query(key)
	}
	return ""
}

func (c *Context) DefaultQuery(key, defaultValue string) string {
	if c.ginCtx != nil {
		return c.ginCtx.DefaultQuery(key, defaultValue)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.DefaultQuery(key, defaultValue)
	}
	return ""
}
func (c *Context) GetQuery(key string) (string, bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetQuery(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetQuery(key)
	}
	return "", false
}

//func (c *Context) QueryArray(key string) (values []string) {
//	if c.ginCtx != nil {
//		return c.ginCtx.QueryArray(key)
//	}
//	if c.hertzCtx != nil {
//		return c.hertzCtx.QueryArgs(key)
//	}
//	return
//}

func (c *Context) PostForm(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.PostForm(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.PostForm(key)
	}
	return ""
}

func (c *Context) DefaultPostForm(key, defaultValue string) string {
	if c.ginCtx != nil {
		return c.ginCtx.DefaultPostForm(key, defaultValue)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.DefaultPostForm(key, defaultValue)
	}
	return ""
}

func (c *Context) PostFormArray(key string) (values []string) {
	if c.ginCtx != nil {
		return c.ginCtx.PostFormArray(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.PostFormArray(key)
	}
	return
}

func (c *Context) GetPostForm(key string) (string, bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetPostForm(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetPostForm(key)
	}
	return "", false
}

func (c *Context) GetPostFormArray(key string) (values []string, ok bool) {
	if c.ginCtx != nil {
		return c.ginCtx.GetPostFormArray(key)
	}
	if c.hertzCtx != nil {
		return c.hertzCtx.GetPostFormArray(key)
	}
	return
}

// BindAndValidate 方法：统一绑定和验证请求数据
//func (c *Context) BindAndValidate(obj any) error {
//	if c.ginCtx != nil {
//		// Gin 的绑定和验证
//		if err := c.ginCtx.ShouldBind(obj); err != nil {
//			return err
//		}
//	}
//	if c.hertzCtx != nil {
//		// Hertz 的绑定和验证
//		if err := c.hertzCtx.BindAndValidate(obj); err != nil {
//			return err
//		}
//	}
//	return nil
//}

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
		wrapperCtx := &Context{
			Ctx:      c,   // 使用 Hertz 提供的上下文
			hertzCtx: ctx, // 保存 Hertz 的上下文
		}
		handler(wrapperCtx)
	}
}

// GinHandler 封装 Gin
func GinHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		wrapperCtx := &Context{
			Ctx:    c.Request.Context(), // 使用 Gin 提供的上下文
			ginCtx: c,                   // 保存 Gin 的上下文
		}
		handler(wrapperCtx)
	}
}
