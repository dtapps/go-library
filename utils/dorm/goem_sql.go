package dorm

import "database/sql"

// Ping ping
func (c *GormClient) Ping() error {
	return c.sqlDd.Ping()
}

// Close 关闭
func (c *GormClient) Close() error {
	return c.sqlDd.Close()
}

// Stats 返回数据库统计信息
func (c *GormClient) Stats() sql.DBStats {
	return c.sqlDd.Stats()
}
