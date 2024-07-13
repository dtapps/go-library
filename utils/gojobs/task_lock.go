package gojobs

import (
	"context"
	"errors"
)

type TaskLockOperation struct {
	client *Client        // 实例
	task   *GormModelTask // 任务
}

func (c *Client) NewLock(task *GormModelTask) (*TaskLockOperation, error) {
	if task.ID == 0 {
		return nil, errors.New("任务数据不正常")
	}
	return &TaskLockOperation{
		client: c,
		task:   task,
	}, nil
}

// Lock 上锁
func (tlo *TaskLockOperation) Lock(ctx context.Context, id any) error {
	_, err := tlo.client.Lock(ctx, tlo.task, id)
	return err
}

// Unlock 解锁
func (tlo *TaskLockOperation) Unlock(ctx context.Context, id any) error {
	return tlo.client.Unlock(ctx, tlo.task, id)
}

// LockForever 永远上锁
func (tlo *TaskLockOperation) LockForever(ctx context.Context, id any) error {
	_, err := tlo.client.LockForever(ctx, tlo.task, id)
	return err
}

// LockId 上锁
func (tlo *TaskLockOperation) LockId(ctx context.Context) error {
	_, err := tlo.client.LockId(ctx, tlo.task)
	return err
}

// UnlockId 解锁
func (tlo *TaskLockOperation) UnlockId(ctx context.Context) error {
	return tlo.client.UnlockId(ctx, tlo.task)
}

// LockForeverId 永远上锁
func (tlo *TaskLockOperation) LockForeverId(ctx context.Context) error {
	_, err := tlo.client.LockForeverId(ctx, tlo.task)
	return err
}

// LockCustomId 上锁
func (tlo *TaskLockOperation) LockCustomId(ctx context.Context) error {
	_, err := tlo.client.LockCustomId(ctx, tlo.task)
	return err
}

// UnlockCustomId 解锁
func (tlo *TaskLockOperation) UnlockCustomId(ctx context.Context) error {
	return tlo.client.UnlockCustomId(ctx, tlo.task)
}

// LockForeverCustomId 永远上锁
func (tlo *TaskLockOperation) LockForeverCustomId(ctx context.Context) error {
	_, err := tlo.client.LockForeverCustomId(ctx, tlo.task)
	return err
}
