package aswzk

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// SetTrace 设置OpenTelemetry链路追踪
func (c *Client) SetTrace(trace bool) {
	c.trace = trace
	c.httpClient.SetTrace(trace)
}

// TraceStartSpan 开始OpenTelemetry链路追踪状态
func (c *Client) TraceStartSpan(ctx context.Context, spanName string) context.Context {
	if c.trace {
		tr := otel.Tracer("go.dtapp.net/aswzk", trace.WithInstrumentationVersion(Version))
		ctx, c.span = tr.Start(ctx, "aswzk."+spanName, trace.WithSpanKind(trace.SpanKindClient))
	}
	return ctx
}

// TraceEndSpan 结束OpenTelemetry链路追踪状态
func (c *Client) TraceEndSpan() {
	if c.trace && c.span != nil {
		c.span.End()
	}
}

// TraceSetAttributes 设置OpenTelemetry链路追踪属性
func (c *Client) TraceSetAttributes(kv ...attribute.KeyValue) {
	if c.trace && c.span != nil {
		c.span.SetAttributes(kv...)
	}
}

// TraceSetStatus 设置OpenTelemetry链路追踪状态
func (c *Client) TraceSetStatus(code codes.Code, description string) {
	if c.trace && c.span != nil {
		c.span.SetStatus(code, description)
	}
}

// TraceRecordError 记录OpenTelemetry链路追踪错误
func (c *Client) TraceRecordError(err error, options ...trace.EventOption) {
	if c.trace && c.span != nil {
		c.span.RecordError(err, options...)
	}
}

// TraceGetTraceID 获取OpenTelemetry链路追踪TraceID
func (c *Client) TraceGetTraceID() (traceID string) {
	if c.trace && c.span != nil {
		traceID = c.span.SpanContext().TraceID().String()
	}
	return traceID
}

// TraceGetSpanID 获取OpenTelemetry链路追踪SpanID
func (c *Client) TraceGetSpanID() (spanID string) {
	if c.trace && c.span != nil {
		spanID = c.span.SpanContext().SpanID().String()
	}
	return spanID
}
