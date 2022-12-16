package wechatunion

import "errors"

var (
	redisCachePrefixNoConfig = errors.New("请配置 RedisCachePrefix")
)
