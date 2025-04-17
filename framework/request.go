package framework

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
)

type RequestWrapper struct {
	ctx      context.Context     // 统一的上下文
	ginCtx   *gin.Context        // Gin 上下文
	hertzCtx *app.RequestContext // Hertz 上下文
}

// 返回请求相关的封装方法
func (c *Context) Request() *RequestWrapper {
	return &RequestWrapper{
		ctx:      c.ctx,
		ginCtx:   c.ginCtx,
		hertzCtx: c.hertzCtx,
	}
}
