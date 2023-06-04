package golog

import "errors"

var (
	currentIpNoConfig   = errors.New("请配置 CurrentIp")
	dbClientFunNoConfig = errors.New("请配置 GormClientFun")
)
