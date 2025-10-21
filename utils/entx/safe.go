package entx

import (
	"context"
	"sync"
)

// ---------------------------------------------
// 全局写锁，保证写操作串行化
// ---------------------------------------------
var dbWriteLock sync.Mutex

// SafeWrite 封装写事务，避免并发写冲突
//
// 参数:
//
//	txFactory: 返回 Txer 的函数，用户负责创建事务
//	f: 用户自定义写操作函数，接收 Txer
//
// 返回值:
//
//	error: 如果事务或写操作失败，则返回错误
//
// 使用方法示例:
//
//	entx.SafeWrite(txFactory, func(tx entx.Txer) error {
//	    t := tx.(*ent.Tx)  // 断言为具体 ent.Tx
//	    return t.User.Create().SetName("Alice").Exec(ctx)
//	})
func SafeWrite(txFactory func(ctx context.Context) (Txer, error), f func(tx Txer) error) error {
	dbWriteLock.Lock()
	defer dbWriteLock.Unlock()

	tx, err := txFactory(context.Background())
	if err != nil {
		return err
	}

	// 确保事务最终被回滚（如果 Commit 已经执行，这里不会影响）
	defer tx.Rollback()

	if err := f(tx); err != nil {
		return err
	}

	// 提交事务
	return tx.Commit()
}

// SafeRead 封装只读事务，保证读取一致性
//
// 参数:
//
//	txFactory: 返回 Txer 的函数，用户负责创建事务
//	f: 用户自定义只读操作函数
//
// 返回值:
//
//	error: 读取失败时返回错误
//
// 说明:
//
//	SafeRead 会创建一个事务，但不加写锁。
//	适合多协程并发读取，保证读取的一致性
func SafeRead(txFactory func(ctx context.Context) (Txer, error), f func(tx Txer) error) error {
	tx, err := txFactory(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback()

	return f(tx)
}
