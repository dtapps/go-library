package entx

// Txer 是事务接口，抽象出 Commit/Rollback 方法和你需要的操作
type Txer interface {
	Commit() error
	Rollback() error
}
