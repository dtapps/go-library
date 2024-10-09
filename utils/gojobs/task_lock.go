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

// Lock5Minute 上锁5分钟
func (tlo *TaskLockOperation) Lock5Minute(ctx context.Context, id any) error {
	_, err := tlo.client.LockMinute(ctx, tlo.task, id, 5)
	return err
}

// Lock10Minute 上锁10分钟
func (tlo *TaskLockOperation) Lock10Minute(ctx context.Context, id any) error {
	_, err := tlo.client.LockMinute(ctx, tlo.task, id, 10)
	return err
}

// Lock15Minute 上锁15分钟
func (tlo *TaskLockOperation) Lock15Minute(ctx context.Context, id any) error {
	_, err := tlo.client.LockMinute(ctx, tlo.task, id, 15)
	return err
}

// Lock30Minute 上锁30分钟
func (tlo *TaskLockOperation) Lock30Minute(ctx context.Context, id any) error {
	_, err := tlo.client.LockMinute(ctx, tlo.task, id, 30)
	return err
}

// Lock60Minute 上锁60分钟
func (tlo *TaskLockOperation) Lock60Minute(ctx context.Context, id any) error {
	_, err := tlo.client.LockMinute(ctx, tlo.task, id, 60)
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

// LockId5Minute 上锁5分钟
func (tlo *TaskLockOperation) LockId5Minute(ctx context.Context) error {
	_, err := tlo.client.LockIdMinute(ctx, tlo.task, 5)
	return err
}

// LockId10Minute 上锁10分钟
func (tlo *TaskLockOperation) LockId10Minute(ctx context.Context) error {
	_, err := tlo.client.LockIdMinute(ctx, tlo.task, 10)
	return err
}

// LockId15Minute 上锁15分钟
func (tlo *TaskLockOperation) LockId15Minute(ctx context.Context) error {
	_, err := tlo.client.LockIdMinute(ctx, tlo.task, 15)
	return err
}

// LockId30Minute 上锁30分钟
func (tlo *TaskLockOperation) LockId30Minute(ctx context.Context) error {
	_, err := tlo.client.LockIdMinute(ctx, tlo.task, 30)
	return err
}

// LockId60Minute 上锁60分钟
func (tlo *TaskLockOperation) LockId60Minute(ctx context.Context) error {
	_, err := tlo.client.LockIdMinute(ctx, tlo.task, 60)
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

// LockCustomId5Minute 上锁5分钟
func (tlo *TaskLockOperation) LockCustomId5Minute(ctx context.Context) error {
	_, err := tlo.client.LockCustomIdMinute(ctx, tlo.task, 5)
	return err
}

// LockCustomId10Minute 上锁10分钟
func (tlo *TaskLockOperation) LockCustomId10Minute(ctx context.Context) error {
	_, err := tlo.client.LockCustomIdMinute(ctx, tlo.task, 10)
	return err
}

// LockCustomId15Minute 上锁15分钟
func (tlo *TaskLockOperation) LockCustomId15Minute(ctx context.Context) error {
	_, err := tlo.client.LockCustomIdMinute(ctx, tlo.task, 15)
	return err
}

// LockCustomId30Minute 上锁30分钟
func (tlo *TaskLockOperation) LockCustomId30Minute(ctx context.Context) error {
	_, err := tlo.client.LockCustomIdMinute(ctx, tlo.task, 30)
	return err
}

// LockCustomId60Minute 上锁60分钟
func (tlo *TaskLockOperation) LockCustomId60Minute(ctx context.Context) error {
	_, err := tlo.client.LockCustomIdMinute(ctx, tlo.task, 60)
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
