package golog

import (
	"context"
)

func NewApiSlog(ctx context.Context) *ApiSLog {

	sl := &ApiSLog{}

	sl.setConfig(ctx)

	return sl
}
