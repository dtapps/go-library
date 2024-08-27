package fiber_requestid

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

var HeaderXRequestID = "X-Request-ID"

// NewKj 初始化RequestID中间件
func NewKj(config ...requestid.Config) fiber.Handler {
	return requestid.New(config...)
}

// New 初始化RequestID中间件
func New() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 从请求中获取ID
		requestId := c.Get(HeaderXRequestID)
		if requestId == "" {
			requestId = uuid.New().String()
		}

		// 设置id
		c.Set(HeaderXRequestID, requestId)
		c.Locals(HeaderXRequestID, requestId)

		// 继续
		return c.Next()
	}
}

// Get 返回请求标识符
func Get(c *fiber.Ctx) string {
	return c.GetRespHeader(fiber.HeaderXRequestID)
}

// GetX 返回请求标识符
func GetX(c *fiber.Ctx) string {
	return c.GetRespHeader(HeaderXRequestID)
}
