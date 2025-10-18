package gin_ent_log

import (
	"context"
	"net/http"
	"time"
)

// RequestInfo 提供请求阶段可用的最小信息，避免依赖第三方框架类型
type RequestInfo struct {
	Version string // 版本号

	Header http.Header
	Start  time.Time

	Method    string
	Path      string
	URL       string
	UserAgent string

	Host string
}

// ResponseInfo 提供响应阶段可用的最小信息
type ResponseInfo struct {
	Status     int
	Header     http.Header
	Body       []byte
	End        time.Time
	DurationMs int64
}

// Tracer 可插拔追踪接口，不引入任何 OTel 依赖
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
