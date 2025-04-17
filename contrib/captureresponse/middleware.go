package captureresponse

import (
	"go.dtapp.net/library/framework"
)

const __responseBodyKey = "__responseBody"

func New() framework.MiddlewareFunc {
	return func(c *framework.Context) {
		if c.GetGinContext() != nil {
			// 使用自定义的 ResponseCaptureWriter 来捕获响应体
			writer := framework.NewResponseCaptureWriter(c.GetGinContext())
			// 将捕获的 writer 存储到 context 中，以便后续访问
			c.Set(__responseBodyKey, writer)
			// 继续处理请求
			c.Next()
		}
		if c.GetHertzContext() != nil {
			c.Next()
		}
	}
}

func Get(c *framework.Context) []byte {
	if ginCtx := c.GetGinContext(); ginCtx != nil {
		if val, exists := ginCtx.Get(__responseBodyKey); exists {
			// Gin 默认是直接写到 ResponseWriter，需要通过自定义的 ResponseWriter 捕获响应体
			if writer, ok := val.(*framework.ResponseCaptureWriter); ok {
				return []byte(writer.GetBody())
			}
		}
		return nil
	}
	if hertzCtx := c.GetHertzContext(); hertzCtx != nil {
		return hertzCtx.Response.Body()
	}
	return nil
}
