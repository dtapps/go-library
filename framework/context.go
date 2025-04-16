package framework

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
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

// Param 方法：统一获取路径参数
func (c *Context) Param(key string) string {
	if c.ginCtx != nil {
		return c.ginCtx.Param(key)
	}
	if c.hertzCtx != nil {
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
	}
	if c.hertzCtx != nil {
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

// GetTest 示例方法
func GetTest(ctx *Context) {

	// 参数
	type User struct {
		Name string `query:"name" binding:"required"`
		Age  int    `json:"age" binding:"gte=0,lte=100"`
		ID   string `path:"id" binding:"required"`
	}

	// 调用 Validator 方法进行验证
	var user User
	if err := ctx.Validator(&user); err != nil {
		errors := ctx.ValidatorError(err)
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"error":   "Validation failed",
			"details": errors,
		})
		return
	}

}
