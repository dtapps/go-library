package gojobs

import (
	"context"
	"fmt"
	"sync"
)

// ConcurrencyMap 使用 sync.Map 实现任务并发控制
type ConcurrencyMap struct {
	ctx       context.Context // 上下文
	wg        sync.WaitGroup  // 等待所有任务完成
	statusMap sync.Map        // 用来存储任务状态
	taskCount int             // 任务数量
}

// NewConcurrencyMap 创建
func NewConcurrencyMap(ctx context.Context, taskCount int) *ConcurrencyMap {
	return &ConcurrencyMap{
		ctx:       ctx,              // 上下文
		wg:        sync.WaitGroup{}, // 等待所有任务完成
		taskCount: taskCount,        // 任务数量
	}
}

// Add 添加任务
func (cm *ConcurrencyMap) Add(taskID int) {
	cm.wg.Add(1)
	cm.statusMap.Store(taskID, "处理中") // 标记任务开始
}

// Done 任务完成
func (cm *ConcurrencyMap) Done(taskID int) {
	cm.statusMap.Store(taskID, "已完成") // 更新任务状态
	cm.wg.Done()
}

// Wait 等待所有任务完成
func (cm *ConcurrencyMap) Wait() {
	cm.wg.Wait()
}

// ShowStatus 显示任务状态
func (cm *ConcurrencyMap) ShowStatus() {
	cm.statusMap.Range(func(key, value interface{}) bool {
		fmt.Printf("任务 %v 状态: %v\n", key, value)
		return true
	})
}
