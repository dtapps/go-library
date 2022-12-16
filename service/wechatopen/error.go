package wechatopen

import "errors"

var (
	redisCachePrefixNoConfig = errors.New("请配置 RedisCachePrefix")
)

var (
	componentAppIdNoConfig     = errors.New("请配置 ComponentAppId")
	componentAppSecretNoConfig = errors.New("请配置 ComponentAppSecret")
)

var (
	authorizerAppidNoConfig = errors.New("请配置 AuthorizerAppid 或 ConfigAuthorizer")
)
