package qxwlwagnt

import (
	"context"

	"resty.dev/v3"
)

type contextKey string

const (
	bodyMapKey contextKey = "bodyMap"
)

type Request struct {
	*resty.Request
}

func (r *Request) SetBodyMap(bodyMap map[string]any) *Request {
	ctx := context.WithValue(r.Context(), bodyMapKey, bodyMap)
	r.SetContext(ctx)
	// 注意：不设置 r.Body，留到中间件处理
	return r
}
