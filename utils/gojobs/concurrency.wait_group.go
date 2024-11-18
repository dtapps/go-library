package gojobs

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

// ConcurrencyWaitGroup 使用 sync.WaitGroup 实现任务并发控制
type ConcurrencyWaitGroup struct {
	ctx            context.Context     // 上下文
	wg             sync.WaitGroup      // 等待所有任务完成
	sem            *semaphore.Weighted // 控制并发度
	maxConcurrency int64               // 最大并发数
}

// NewConcurrencyWaitGroup 创建
func (c *Client) NewConcurrencyWaitGroup(ctx context.Context, maxConcurrency int64) *ConcurrencyWaitGroup {
	return &ConcurrencyWaitGroup{
		ctx:            ctx,                                   // 上下文
		wg:             sync.WaitGroup{},                      // 等待
		sem:            semaphore.NewWeighted(maxConcurrency), // 控制并发度
		maxConcurrency: maxConcurrency,                        // 最大并发数
	}
}

// Add 添加任务
func (cw *ConcurrencyWaitGroup) Add() error {
	// 获取信号量，确保并发数量不超过限制
	if err := cw.sem.Acquire(cw.ctx, 1); err != nil {
		return err // 获取信号量失败
	}
	cw.wg.Add(1)
	return nil
}

// Done 完成任务
func (cw *ConcurrencyWaitGroup) Done() {
	cw.wg.Done()
	cw.sem.Release(1)
}

// Wait 等待所有任务完成
func (cw *ConcurrencyWaitGroup) Wait() {
	cw.wg.Wait()
}
