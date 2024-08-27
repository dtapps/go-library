package hertz_requestid

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
	"github.com/hertz-contrib/requestid"
)

var HeaderXRequestID = "X-Request-ID"

// NewKj 初始化RequestID中间件
func NewKj(opts ...requestid.Option) app.HandlerFunc {
	return requestid.New(opts...)
}

// GetKj 返回请求标识符
func GetKj(c *app.RequestContext) string {
	return requestid.Get(c)
}

// New 初始化RequestID中间件
func New() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取ID
		rid := c.Request.Header.Get(HeaderXRequestID)
		if rid == "" {
			rid = uuid.New().String()
		}

		// 设置id
		c.Set(HeaderXRequestID, rid)
		c.Header(HeaderXRequestID, rid)
		ctx = context.WithValue(ctx, HeaderXRequestID, rid)

		// 继续
		c.Next(ctx)
	}
}

// Get 返回请求标识符
func Get(c *app.RequestContext) string {
	return c.Response.Header.Get(HeaderXRequestID)
}

// GetX 返回请求标识符
func GetX(c *app.RequestContext) string {
	return c.GetString(HeaderXRequestID)
}
