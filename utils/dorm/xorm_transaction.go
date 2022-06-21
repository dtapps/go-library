package dorm

import (
	"xorm.io/xorm"
)

// XormClientSession https://xorm.io/zh/docs/chapter-10/readme/
type XormClientSession struct {
	*xorm.Session
}

// Begin 开始事务，需要创建 Session 对象
//func (c *XormClient) Begin() (*XormClientSession, error) {
//	session := c.Db.NewSession()
//	defer session.Close()
//	return &session, session.Begin()
//}

// Rollback 回滚事务
//func (c *XormClientSession) Rollback() error {
//	return c.Rollback()
//}

// Commit 提交事务
//func (c *XormClientSession) Commit() error {
//	return c.Commit()
//}
