package gin_requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var HeaderXRequestID = "X-Request-ID"

// New 初始化RequestID中间件
func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取ID
		requestId := c.Request.Header.Get(HeaderXRequestID)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(HeaderXRequestID, requestId)
		// 设置id以确保请求id在响应中
		//c.Header(HeaderXRequestID, requestId)
		c.Writer.Header().Set(HeaderXRequestID, requestId)
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
