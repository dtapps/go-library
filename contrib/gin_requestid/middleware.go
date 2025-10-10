package gin_requestid

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var HeaderXRequestID = "X-Request-ID"

// NewKj 初始化RequestID中间件
func NewKj() gin.HandlerFunc {
	return requestid.New()
}

// GetKj 返回请求标识符
func GetKj(c *gin.Context) string {
	return requestid.Get(c)
}

// New 初始化RequestID中间件
func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取ID
		rid := c.Request.Header.Get(HeaderXRequestID)
		if rid == "" {
			rid = uuid.Must(uuid.NewV7()).String()
			//c.Request.Header.Add(HeaderXRequestID, rid)
		}

		// 设置id
		c.Set(HeaderXRequestID, rid)
		//c.Header(HeaderXRequestID, requestId)
		c.Writer.Header().Set(HeaderXRequestID, rid)

		// 继续
		c.Next()
	}
}

// Get 返回请求标识符
func Get(c *gin.Context) string {
	//return c.GetHeader(HeaderXRequestID)
	return c.Writer.Header().Get(HeaderXRequestID)
}

// GetX 返回请求标识符
func GetX(c *gin.Context) string {
	return c.GetString(HeaderXRequestID)
}
