package gojobs

import "errors"

var (
	currentIpNoConfig      = errors.New("请配置 CurrentIp")
	redisPrefixFunNoConfig = errors.New("请配置 RedisPrefixFun")
	gormClientFunNoConfig  = errors.New("请配置 GormClientFun")
	mongoClientFunNoConfig = errors.New("请配置 MongoClientFun")
	TaskIsExist            = errors.New("任务已存在")
)
