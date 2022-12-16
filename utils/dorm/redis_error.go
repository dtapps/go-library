package dorm

import "errors"

var (
	// RedisKeysNotFound keys没有数据
	RedisKeysNotFound = errors.New("ERR wrong number of arguments for 'mget' command")
)
