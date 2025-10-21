package entx

// ---------------------------------------------
// 通用事务接口，抽象 Commit 和 Rollback 方法
// 用户传入的事务必须实现这个接口
// ---------------------------------------------
type Txer interface {
	Commit() error
	Rollback() error
}
