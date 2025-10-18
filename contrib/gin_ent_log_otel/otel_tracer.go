package gin_ent_log_otel

import (
	"context"
	"net/http"

	"go.dtapp.net/library/contrib/gin_ent_log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// otelTracer 使用 OpenTelemetry 实现 gin_ent_log.Tracer
type otelTracer struct {
	tracer trace.Tracer
}

// Enable 在调用方启用 OTel 追踪
// 需要调用方在应用入口处先初始化全局 TracerProvider/Exporter/Propagator。
func Enable() {
	tp := otel.GetTracerProvider()
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	gin_ent_log.SetTracer(&otelTracer{
		tracer: tp.Tracer("go.dtapp.net/library/contrib/gin_ent_log_otel"),
	})
}

func (t *otelTracer) Start(parent context.Context, req gin_ent_log.RequestInfo) context.Context {
	// 从请求头提取上游上下文
	parent = otel.GetTextMapPropagator().Extract(parent, propagation.HeaderCarrier(req.Header))
	spanCtx, span := t.tracer.Start(parent, "Gin HTTP "+req.Method, trace.WithSpanKind(trace.SpanKindServer))
	// 记录请求属性（响应属性待 End 阶段设置）
	span.SetAttributes(
		// 库版本号
		attribute.String("instrumentation.version", req.Version),

		// HTTP 语义约定属性
		attribute.String("http.method", req.Method),
		attribute.String("http.path", req.Path),
		attribute.String("http.url", req.URL),
		attribute.String("http.user_agent", req.UserAgent),

		// 目标主机
		attribute.String("net.peer.name", req.Host),
		attribute.String("http.host", req.Host),
	)
	return spanCtx
}

func (t *otelTracer) End(ctx context.Context, resp gin_ent_log.ResponseInfo) {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return
	}
	span.SetAttributes(
		attribute.Int("http.status_code", resp.Status),                // 状态码
		attribute.Int("http.response_content_length", len(resp.Body)), // 响应长度
		attribute.Int64("http.elapsed_ms", resp.DurationMs),           // 耗时
	)
	if resp.Status >= 500 {
		span.SetStatus(codes.Error, http.StatusText(resp.Status))
	}
	span.End()
}
