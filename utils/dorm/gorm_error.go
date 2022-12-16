package dorm

import "gorm.io/gorm"

var (
	// GormNotFound 没有数据
	GormNotFound = gorm.ErrRecordNotFound
)
