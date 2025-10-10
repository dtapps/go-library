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
		rid := c.Get(HeaderXRequestID)
		if rid == "" {
			rid = uuid.Must(uuid.NewV7()).String()
		}

		// 设置id
		c.Set(HeaderXRequestID, rid)
		c.Locals(HeaderXRequestID, rid)

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
	return c.Locals(HeaderXRequestID).(string)
}
