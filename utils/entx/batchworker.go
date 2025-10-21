package entx

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

// ---------------------------------------------
// 批量写事务相关
// ---------------------------------------------

// JobFunc 用户写任务函数类型
type JobFunc func(tx Txer) error

// BatchWorker 封装写入队列和批量写事务
type BatchWorker struct {
	txFactory func(ctx context.Context) (Txer, error) // 事务工厂
	writeChan chan JobFunc                            // 写任务队列
	batchSize int                                     // 批量提交条数
	interval  time.Duration                           // 批量提交间隔
	wg        sync.WaitGroup                          // 用于等待 worker 退出
	closeCh   chan struct{}                           // 用于安全关闭

	// 统计字段
	statsMu           sync.Mutex // 保护统计字段
	writeAttempts     int64      // 写入尝试次数
	writeSuccesses    int64      // 写入成功次数
	writeFailures     int64      // 写入失败次数
	dataRowsProcessed int64      // 处理的数据条数
}

// NewBatchWorker 创建批量写入队列
//
// 参数:
//
//	txFactory: 返回 Txer 的函数
//	batchSize: 批量提交条数
//	interval: 批量提交间隔
//
// 返回值:
//
//	*BatchWorker: 批量写入器
//
// 使用示例:
//
//	bw := entx.NewBatchWorker(txFactory, 100, 3*time.Second)
//	bw.Submit(func(tx entx.Txer) error { ... })
//	defer bw.Close()
func NewBatchWorker(txFactory func(ctx context.Context) (Txer, error), batchSize int, interval time.Duration) *BatchWorker {
	bw := &BatchWorker{
		txFactory: txFactory,
		writeChan: make(chan JobFunc, 1000),
		batchSize: batchSize,
		interval:  interval,
		closeCh:   make(chan struct{}),
	}
	bw.wg.Add(1)
	go bw.worker()
	return bw
}

// Submit 提交写任务到队列
func (bw *BatchWorker) Submit(job JobFunc) {
	bw.writeChan <- job
}

// worker 后台批量写处理协程
func (bw *BatchWorker) worker() {
	defer bw.wg.Done()

	var batch []JobFunc
	timer := time.NewTimer(bw.interval)
	defer timer.Stop()

	// flush 批量提交函数
	flush := func() {
		if len(batch) == 0 {
			return
		}
		if err := bw.safeBatchWrite(batch); err != nil {
			slog.Error("批量写入失败", slog.String("err", err.Error()))
		}
		batch = batch[:0]
		timer.Reset(bw.interval)
	}

	for {
		select {
		case job := <-bw.writeChan:
			batch = append(batch, job)
			if len(batch) >= bw.batchSize {
				flush()
			}
		case <-timer.C:
			if len(batch) > 0 {
				slog.Info("定时器触发，准备执行 flush()")
				flush()
			} else {
				// 批次为空也需要重置定时器
				timer.Reset(bw.interval)
			}
		case <-bw.closeCh:
			slog.Info("程序关闭，准备执行 flush()")
			flush()
			return
		}
	}
}

// safeBatchWrite 批量写事务
//
// 参数:
//
//	jobs: 需要执行的写任务数组
//
// 返回值:
//
//	error: 如果事务提交失败，则返回错误
func (bw *BatchWorker) safeBatchWrite(jobs []JobFunc) error {
	bw.statsMu.Lock()
	bw.writeAttempts++
	bw.statsMu.Unlock()

	tx, err := bw.txFactory(context.Background())
	if err != nil {
		bw.statsMu.Lock()
		bw.writeFailures++
		bw.statsMu.Unlock()
		return err
	}
	defer tx.Rollback() // 确保在提交或返回时回滚

	processedCount := 0
	for _, job := range jobs {
		if err := job(tx); err != nil {
			bw.statsMu.Lock()
			bw.writeFailures++
			bw.statsMu.Unlock()
			return err
		}
		processedCount++
	}

	if err := tx.Commit(); err != nil {
		bw.statsMu.Lock()
		bw.writeFailures++
		bw.statsMu.Unlock()
		return err
	}

	bw.statsMu.Lock()
	bw.writeSuccesses++
	bw.dataRowsProcessed += int64(processedCount)
	bw.statsMu.Unlock()
	return nil
}

// Close 安全关闭队列，确保剩余任务写入
func (bw *BatchWorker) Close() {
	close(bw.closeCh)
	bw.wg.Wait()
}

// BatchWorkerStats 批量写入器的统计信息
type BatchWorkerStats struct {
	WriteAttempts     int64 `json:"writeAttempts"`     // 写入尝试次数
	WriteSuccesses    int64 `json:"writeSuccesses"`    // 写入成功次数
	WriteFailures     int64 `json:"writeFailures"`     // 写入失败次数
	DataRowsProcessed int64 `json:"dataRowsProcessed"` // 处理的数据条数
}

// GetStats 获取当前批量写入器的统计信息
func (bw *BatchWorker) GetStats() BatchWorkerStats {
	bw.statsMu.Lock()
	defer bw.statsMu.Unlock()
	return BatchWorkerStats{
		WriteAttempts:     bw.writeAttempts,
		WriteSuccesses:    bw.writeSuccesses,
		WriteFailures:     bw.writeFailures,
		DataRowsProcessed: bw.dataRowsProcessed,
	}
}
