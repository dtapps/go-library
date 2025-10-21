package entx

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

type JobFunc func(tx Txer) error

type BatchWorker struct {
	txFactory func(ctx context.Context) (Txer, error)
	writeChan chan JobFunc
	batchSize int
	interval  time.Duration
	wg        sync.WaitGroup
	closeCh   chan struct{}

	statsMu           sync.Mutex
	writeAttempts     int64
	writeSuccesses    int64
	writeFailures     int64
	dataRowsProcessed int64
}

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

func (bw *BatchWorker) Submit(job JobFunc) {
	bw.writeChan <- job
}

func (bw *BatchWorker) worker() {
	defer bw.wg.Done()

	var batch []JobFunc
	timer := time.NewTimer(bw.interval)
	defer timer.Stop()

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
				timer.Reset(bw.interval)
			}
		case <-bw.closeCh:
			slog.Info("程序关闭，准备执行 flush()")
			flush()
			return
		}
	}
}

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
	defer tx.Rollback()

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

func (bw *BatchWorker) Close() {
	close(bw.closeCh)
	bw.wg.Wait()
}

type BatchWorkerStats struct {
	WriteAttempts     int64 `json:"writeAttempts"`
	WriteSuccesses    int64 `json:"writeSuccesses"`
	WriteFailures     int64 `json:"writeFailures"`
	DataRowsProcessed int64 `json:"dataRowsProcessed"`
}

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
