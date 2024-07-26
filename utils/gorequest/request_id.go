package gorequest

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gostring"
)

var (
	XRequestID = "X-Request-Id"
	TNil       = "%!s(<nil>)"
)

// SetRequestIDContext 设置请求编号
func SetRequestIDContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, XRequestID, gostring.GetUuId())
}

// GetRequestIDContext 获取请求编号
func GetRequestIDContext(ctx context.Context) string {
	traceId := fmt.Sprintf("%s", ctx.Value(XRequestID))
	if traceId == TNil {
		return ""
	}
	if len(traceId) <= 0 {
		return ""
	}
	return traceId
}
