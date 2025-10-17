package resty_log

import (
	"context"
	"net/http"
	"time"
)

// RequestInfo 出站请求阶段的最小信息
type RequestInfo struct {
	Method string
	URL    string
	Host   string
	Header http.Header
	Start  time.Time
}

// ResponseInfo 出站响应阶段的最小信息
type ResponseInfo struct {
	Status     int
	Header     http.Header
	Body       []byte
	End        time.Time
	DurationMs int64
}

// Tracer 可插拔追踪接口（核心包不依赖任何 OTel 包）
type Tracer interface {
	// Start 返回带追踪的上下文（用于注入到请求或后续阶段）
	Start(parent context.Context, req RequestInfo) context.Context
	// End 在响应阶段收尾并记录属性
	End(ctx context.Context, resp ResponseInfo)
}

// 包级注册点
var tracer Tracer

// SetTracer 注册一个 Tracer；传 nil 可禁用追踪
func SetTracer(t Tracer) {
	tracer = t
}
