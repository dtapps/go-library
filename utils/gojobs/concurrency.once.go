package gojobs

import (
	"context"
	"sync"
)

// ConcurrencyOnce 使用 sync.Once 实现任务并发控制
type ConcurrencyOnce struct {
	ctx       context.Context // 上下文
	wg        sync.WaitGroup  // 等待所有任务完成
	once      sync.Once
	taskCount int // 任务数量
}

// NewConcurrencyOnce 创建
func NewConcurrencyOnce(ctx context.Context, taskCount int) *ConcurrencyOnce {
	return &ConcurrencyOnce{
		ctx:       ctx,              // 上下文
		wg:        sync.WaitGroup{}, // 等待所有任务完成
		taskCount: taskCount,        // 任务数量
	}
}

// Add 添加
func (co *ConcurrencyOnce) Add() {
	co.wg.Add(1)
}

// Done 完成任务
func (co *ConcurrencyOnce) Done() {
	co.wg.Done()
}

// InitOnce 初始化操作（只执行一次）
func (co *ConcurrencyOnce) InitOnce(f func()) {
	co.once.Do(f)
}

// Wait 等待所有任务完成
func (co *ConcurrencyOnce) Wait() {
	co.wg.Wait()
}
