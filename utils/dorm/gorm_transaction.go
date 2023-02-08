package dorm

import "gorm.io/gorm"

// Begin 开始事务，不需要创建 Session 对象
func (c *GormClient) Begin() *gorm.DB {
	return c.db.Begin()
}

// Rollback 回滚事务
func (c *GormClient) Rollback() *gorm.DB {
	return c.db.Rollback()
}

// Commit 提交事务
func (c *GormClient) Commit() *gorm.DB {
	return c.db.Commit()
}
