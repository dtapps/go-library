package hertz_requestid

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

var headerXRequestID = "X-Request-ID"

// New 初始化RequestID中间件
func New() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取ID
		requestId := c.Request.Header.Get(headerXRequestID)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(headerXRequestID, requestId)
		// 设置id以确保请求id在响应中
		c.Header(headerXRequestID, requestId)
		ctx = context.WithValue(ctx, headerXRequestID, requestId)
		c.Next(ctx)
	}
}

// Get 返回请求标识符
func Get(c *app.RequestContext) string {
	return c.Response.Header.Get(headerXRequestID)
}
