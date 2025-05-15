package framework

import (
	"bytes"
	"io"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
)

const __rawBodyKey = "__rawBody"

func (c *Context) GinCacheBody(ginCtx *gin.Context) ([]byte, error) {

	// 检查是否已经缓存了请求体
	if val, exists := ginCtx.Get(__rawBodyKey); exists {
		if body, ok := val.([]byte); ok {
			return body, nil
		}
	}

	// 读取并缓存请求体
	bodyBytes, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return nil, err
	}

	// 如果请求体为空，直接返回 nil
	if len(bodyBytes) == 0 {
		return nil, nil
	}

	// 重置请求体为缓存内容
	ginCtx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// 缓存请求体内容到 gin.Context
	ginCtx.Set(__rawBodyKey, bodyBytes)

	return bodyBytes, nil
}

func (c *Context) HertzCacheBody(hertzCtx *app.RequestContext) ([]byte, error) {

	// 检查是否已经缓存了请求体
	if val, exists := hertzCtx.Get(__rawBodyKey); exists {
		return val.([]byte), nil
	}

	// 读取并缓存请求体
	bodyBytes, err := hertzCtx.Body()
	if err != nil {
		return nil, err
	}

	// 如果请求体为空，直接返回 nil
	if len(bodyBytes) == 0 {
		return nil, nil
	}

	// 缓存请求体内容到 hertz.Context
	hertzCtx.Set(__rawBodyKey, bodyBytes)

	return bodyBytes, nil
}

func (c *Context) CacheBody() ([]byte, error) {
	if c.IsGin() {
		return c.GinCacheBody(c.ginCtx)
	}
	if c.IsHertz() {
		return c.HertzCacheBody(c.hertzCtx)
	}
	return nil, nil
}
