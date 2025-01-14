package dorm

import (
	"database/sql"
	"gorm.io/gorm"
)

// GetDb 获取驱动
func (c *GormClient) GetDb() *gorm.DB {
	return c.db
}

// GetSqlDb 获取驱动
func (c *GormClient) GetSqlDb() *sql.DB {
	return c.sqlDd
}
