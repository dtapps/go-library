package gorequest

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gostring"
)

var (
	xRequestID = "X-Request-Id"
	tNil       = "%!s(<nil>)"
)

// SetRequestIDContext 设置请求编号
func SetRequestIDContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, xRequestID, gostring.GetUuId())
}

// GetRequestIDContext 获取请求编号
func GetRequestIDContext(ctx context.Context) string {
	traceId := fmt.Sprintf("%s", ctx.Value(xRequestID))
	if traceId == tNil {
		return ""
	}
	if len(traceId) <= 0 {
		return ""
	}
	return traceId
}
