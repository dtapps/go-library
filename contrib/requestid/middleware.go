package requestid

import (
	"context"
	"github.com/google/uuid"
	"go.dtapp.net/library/framework"
)

const __headerXRequestID = "X-Request-ID"

func New() framework.MiddlewareFunc {
	return func(c *framework.Context) {
		// 从请求中获取ID
		rid := c.GetHeader(__headerXRequestID)
		if rid == "" {
			rid = uuid.Must(uuid.NewV7()).String()
		}

		// 设置id
		c.Set(__headerXRequestID, rid)
		c.Header(__headerXRequestID, rid)
		c.SetContext(context.WithValue(c.GetContext(), __headerXRequestID, rid))

		// 继续
		c.Next()
	}
}

func Get(c *framework.Context) string {
	return c.Response().GetHeader(__headerXRequestID)
}
