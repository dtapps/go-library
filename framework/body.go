package framework

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

const __rawBodyKey = "__rawBody"

func (c *Context) GinCacheBody(ginCtx *gin.Context) ([]byte, error) {
	// 检查是否已经缓存了请求体
	if val, exists := ginCtx.Get(__rawBodyKey); exists {
		log.Printf("[GinCacheBody] 已缓存请求体: %s", __rawBodyKey)
		if body, ok := val.([]byte); ok {
			return body, nil
		}
		log.Printf("[GinCacheBody] 类型断言失败: %s", __rawBodyKey)
	}
	log.Printf("[GinCacheBody] 开始读取并缓存请求体: %s", __rawBodyKey)

	// 读取并缓存请求体
	bodyBytes, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("[GinCacheBody] 请求体读取成功: %s", __rawBodyKey)

	// 如果请求体为空，直接返回 nil
	if len(bodyBytes) == 0 {
		log.Printf("[GinCacheBody] 请求体为空: %s", __rawBodyKey)
		return nil, nil
	}

	// 重置请求体为缓存内容
	ginCtx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	log.Printf("[GinCacheBody] 重置请求体: %s", __rawBodyKey)

	// 缓存请求体内容到 gin.Context
	ginCtx.Set(__rawBodyKey, bodyBytes)
	log.Printf("[GinCacheBody] 请求体缓存成功: %s", __rawBodyKey)

	return bodyBytes, nil
}

func (c *Context) HertzCacheBody(hertzCtx *app.RequestContext) ([]byte, error) {
	// 检查是否已经缓存了请求体
	if val, exists := hertzCtx.Get(__rawBodyKey); exists {
		log.Printf("[HertzCacheBody] 已缓存请求体: %s", __rawBodyKey)
		return val.([]byte), nil
	}
	log.Printf("[HertzCacheBody] 开始读取并缓存请求体: %s", __rawBodyKey)

	// 读取并缓存请求体
	bodyBytes, err := hertzCtx.Body()
	if err != nil {
		return nil, err
	}
	log.Printf("[HertzCacheBody] 请求体读取成功: %s", __rawBodyKey)

	// 如果请求体为空，直接返回 nil
	if len(bodyBytes) == 0 {
		log.Printf("[HertzCacheBody] 请求体为空: %s", __rawBodyKey)
		return nil, nil
	}

	// 缓存请求体内容到 hertz.Context
	hertzCtx.Set(__rawBodyKey, bodyBytes)
	log.Printf("[HertzCacheBody] 请求体缓存成功: %s", __rawBodyKey)

	return bodyBytes, nil
}

func (c *Context) CacheBody() ([]byte, error) {
	if c.ginCtx != nil {
		return c.GinCacheBody(c.ginCtx)
	}
	if c.hertzCtx != nil {
		return c.HertzCacheBody(c.hertzCtx)
	}
	return nil, nil
}
