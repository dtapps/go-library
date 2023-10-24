package golog

import (
	"context"
)

func NewGinCustomClient(ctx context.Context, config *GinCustomClientConfig) (*GinCustomClient, error) {
	c := &GinCustomClient{}
	c.ipService = config.IpService
	return c, nil
}
