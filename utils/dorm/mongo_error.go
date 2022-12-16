package dorm

import "errors"

var (
	NoConfigDatabaseName   = errors.New("没有配置库名")
	NoConfigCollectionName = errors.New("没有配置集合名")
)
