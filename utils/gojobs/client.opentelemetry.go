package gojobs

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// NewTraceStartSpan 开始OpenTelemetry链路追踪状态
func NewTraceStartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return gorequest.TraceNewSpan(ctx, "go.dtapp.net/gojobs", "", spanName, Version, trace.SpanKindClient)
}

// TraceStartSpan 开始OpenTelemetry链路追踪状态
func TraceStartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return gorequest.TraceNewSpan(ctx, "go.dtapp.net/gojobs", "gojobs.", spanName, Version, trace.SpanKindClient)
}

// TraceEndSpan 结束OpenTelemetry链路追踪状态
func TraceEndSpan(span trace.Span) {
	gorequest.TraceEndSpan(span)
}

// TraceSetAttributes 设置OpenTelemetry链路追踪属性
func TraceSetAttributes(ctx context.Context, kv ...attribute.KeyValue) {
	gorequest.TraceSetAttributes(ctx, kv...)
}

// TraceSetStatus 设置OpenTelemetry链路追踪状态
func TraceSetStatus(ctx context.Context, code codes.Code, description string) {
	gorequest.TraceSetStatus(ctx, code, description)
}

// TraceRecordError 记录OpenTelemetry链路追踪错误
func TraceRecordError(ctx context.Context, err error, options ...trace.EventOption) {
	gorequest.TraceRecordError(ctx, err, options...)
}

// TraceGetTraceID 获取OpenTelemetry链路追踪TraceID
func TraceGetTraceID(ctx context.Context) (traceID string) {
	return gorequest.TraceGetTraceID(ctx)
}

// TraceGetSpanID 获取OpenTelemetry链路追踪SpanID
func TraceGetSpanID(ctx context.Context) (spanID string) {
	return gorequest.TraceGetSpanID(ctx)
}
