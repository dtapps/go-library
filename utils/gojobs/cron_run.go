package gojobs

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gotime"
	"time"
)

func (c *Client) StartHandle(ctx context.Context, key any, overdue int64) error {
	status, err := c.redisConfig.client.Get(ctx, fmt.Sprintf("%v", key)).Result()
	if status != "" {
		return fmt.Errorf("【%v】上一次还在执行中", fmt.Sprintf("%v", key))
	}

	err = c.redisConfig.client.Set(ctx, fmt.Sprintf("%v", key), gotime.Current().Format(), time.Duration(overdue)*time.Second).Err()
	if err != nil {
		err = fmt.Errorf("【%v】设置 %s", fmt.Sprintf("%v", key), err)
	}

	return nil
}
func (c *Client) EndHandle(ctx context.Context, key any) {
	err := c.redisConfig.client.Del(ctx, fmt.Sprintf("%v", key)).Err()
	if err != nil {
		err = fmt.Errorf("【%v】删除 %s", fmt.Sprintf("%v", key), err)
	}
}
