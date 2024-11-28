package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"strings"
	"time"
)

// GetRedisKeyName 获取Redis键名
func GetRedisKeyName(taskType string) string {
	var builder strings.Builder
	builder.WriteString("task:run:")
	builder.WriteString(taskType)
	return builder.String()
}

// SetRedisKeyValue 返回设置Redis键值
func SetRedisKeyValue(ctx context.Context, taskType string) (context.Context, string, any, time.Duration) {
	return ctx,
		GetRedisKeyName(taskType),
		fmt.Sprintf(
			"%s-%s",
			gotime.Current().SetFormat(gotime.DateTimeZhFormat),
			gorequest.GetRequestIDContext(ctx),
		),
		0
}

// SetRedisKeyValueExpiration 返回设置Redis键值，有过分时间
func SetRedisKeyValueExpiration(ctx context.Context, taskType string, expiration int64) (context.Context, string, any, time.Duration) {
	return ctx,
		GetRedisKeyName(taskType),
		fmt.Sprintf(
			"%s-%s",
			gotime.Current().SetFormat(gotime.DateTimeZhFormat),
			gorequest.GetRequestIDContext(ctx),
		),
		time.Duration(expiration)
}
