package framework

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseWrapper 用于统一封装响应处理
type ResponseWrapper struct {
	Ctx      context.Context     // 统一的上下文
	ginCtx   *gin.Context        // Gin 上下文
	hertzCtx *app.RequestContext // Hertz 上下文
}

// Response 返回响应相关的封装方法
func (c *Context) Response() *ResponseWrapper {
	return &ResponseWrapper{
		Ctx:      c.Ctx,
		ginCtx:   c.ginCtx,
		hertzCtx: c.hertzCtx,
	}
}

// GetHeader 获取响应的Header
func (cr *ResponseWrapper) GetHeader(key string) string {
	if cr.ginCtx != nil {
		return cr.ginCtx.Writer.Header().Get(key)
	}
	if cr.hertzCtx != nil {
		return cr.hertzCtx.Response.Header.Get(key)
	}
	return ""
}

// StatusCode 获取响应的状态码
func (cr *ResponseWrapper) StatusCode() int {
	if cr.ginCtx != nil {
		return cr.ginCtx.Writer.Status()
	}
	if cr.hertzCtx != nil {
		return cr.hertzCtx.Response.StatusCode()
	}
	return 0
}

// Body 获取响应的 body 内容
func (cr *ResponseWrapper) Body() []byte {
	if cr.ginCtx != nil {
		// Gin 默认是直接写到 ResponseWriter，需要通过自定义的 ResponseWriter 捕获响应体
		if val, exists := cr.ginCtx.Get(__responseBodyKey); exists {
			if body, ok := val.([]byte); ok {
				return body
			}
		}
		return nil
	}
	if cr.hertzCtx != nil {
		return cr.hertzCtx.Response.Body()
	}
	return nil
}

const __responseBodyKey = "__responseBody"

// ResponseCaptureWriter 用于捕获响应体
type ResponseCaptureWriter struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

// NewResponseCaptureWriter 创建一个新的 ResponseCaptureWriter 实例
func NewResponseCaptureWriter(c *gin.Context) *ResponseCaptureWriter {
	// 创建一个新的 Buffer 用于捕获响应体
	body := bytes.NewBuffer([]byte{})
	// 创建一个新的 ResponseCaptureWriter
	writer := &ResponseCaptureWriter{
		ResponseWriter: c.Writer,
		body:           body,
	}
	// 替换 Gin 的原始 ResponseWriter
	c.Writer = writer
	return writer
}

// Write 捕获响应体并将数据写入响应
func (w *ResponseCaptureWriter) Write(b []byte) (int, error) {
	// 捕获响应体
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString 写入字符串并捕获响应体
func (w *ResponseCaptureWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// WriteHeader 设置响应状态码
func (w *ResponseCaptureWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Header 获取响应头
func (w *ResponseCaptureWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// GetBody 获取捕获的响应体
func (w *ResponseCaptureWriter) GetBody() string {
	return w.body.String()
}

// GetStatusCode 获取捕获的响应状态码
func (w *ResponseCaptureWriter) GetStatusCode() int {
	return w.statusCode
}
