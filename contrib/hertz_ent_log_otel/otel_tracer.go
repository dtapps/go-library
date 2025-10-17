package hertz_ent_log_otel

import (
	"context"

	"go.dtapp.net/library/contrib/hertz_ent_log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// otelTracer 使用 OpenTelemetry 实现 hertz_ent_log.Tracer
type otelTracer struct {
	tracer      trace.Tracer
	serviceName string
}

// Enable 在调用方启用 OTel 追踪（不改动原库的依赖）。
// 需要调用方在应用入口处先初始化全局 TracerProvider/Exporter/Propagator。
func Enable(serviceName string) {
	tp := otel.GetTracerProvider()
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	hertz_ent_log.SetTracer(&otelTracer{
		tracer:      tp.Tracer("go.dtapp.net/library/contrib/hertz_ent_log_otel"),
		serviceName: serviceName,
	})
}

func (t *otelTracer) Start(parent context.Context, req hertz_ent_log.RequestInfo) context.Context {
	// 从请求头提取上游上下文
	parent = otel.GetTextMapPropagator().Extract(parent, propagation.HeaderCarrier(req.Header))
	spanCtx, span := t.tracer.Start(parent, "HTTP "+req.Method, trace.WithSpanKind(trace.SpanKindServer))
	// 记录请求属性（响应属性待 End 阶段设置）
	span.SetAttributes(
		attribute.String("service.name", t.serviceName),
		attribute.String("http.method", req.Method),
		attribute.String("http.target", req.Path),
		attribute.String("http.host", req.Host),
	)
	return spanCtx
}

func (t *otelTracer) End(ctx context.Context, resp hertz_ent_log.ResponseInfo) {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return
	}
	span.SetAttributes(
		attribute.Int("http.status_code", resp.Status),
		attribute.Int("http.response_content_length", len(resp.Body)),
		attribute.Int64("http.elapsed_ms", resp.DurationMs),
	)
	if resp.Status >= 500 {
		span.SetAttributes(attribute.Bool("http.server.error", true))
	}
	span.End()
}
