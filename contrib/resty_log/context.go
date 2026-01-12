package resty_log

import (
	"context"
	"time"
)

// startTimeKey 用于存储操作开始时间的键
type startTimeKey struct{}

// WithStartTimeKey 指定操作来源并返回新 context
func WithStartTimeKey(ctx context.Context, startTime time.Time) context.Context {
	return context.WithValue(ctx, startTimeKey{}, startTime)
}

// GetStartTimeKey 从 context 中获取时间
func GetStartTimeKey(ctx context.Context) time.Time {
	if v := ctx.Value(startTimeKey{}); v != nil {
		if t, ok := v.(time.Time); ok {
			return t
		}
	}
	return time.Now().UTC()
}

// 用于存储请求体的键
type requestBodyKey struct{}

// WithRequestBodyKey 指定请求体并返回新 context
func WithRequestBodyKey(ctx context.Context, requestBody []byte) context.Context {
	return context.WithValue(ctx, requestBodyKey{}, requestBody)
}

// GetRequestBodyKey 从 context 中获取请求体
func GetRequestBodyKey(ctx context.Context) []byte {
	if v := ctx.Value(requestBodyKey{}); v != nil {
		if body, ok := v.([]byte); ok {
			return body
		}
	}
	return nil
}

// 用于存储响应体的键
type responseBodyKey struct{}

// WithResponseBodyKey 指定响应体并返回新 context
func WithResponseBodyKey(ctx context.Context, responseBody []byte) context.Context {
	return context.WithValue(ctx, responseBodyKey{}, responseBody)
}

// GetResponseBodyKey 从 context 中获取响应体
func GetResponseBodyKey(ctx context.Context) []byte {
	if v := ctx.Value(responseBodyKey{}); v != nil {
		if body, ok := v.([]byte); ok {
			return body
		}
	}
	return nil
}
