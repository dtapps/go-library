package gin_requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var headerXRequestID = "X-Request-ID"

// New 初始化RequestID中间件
func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中获取ID
		requestId := c.Request.Header.Get(headerXRequestID)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Set(headerXRequestID, requestId)
		// 设置id以确保请求id在响应中
		c.Header(headerXRequestID, requestId)
		c.Next()
	}
}

// Get 返回请求标识符
func Get(c *gin.Context) string {
	return c.GetHeader(headerXRequestID)
}
