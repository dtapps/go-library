package dredis

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"time"
)

const (
	SerializerJson       = "json"
	SerializerSimpleJson = "simplejson"
	SerializerString     = "string"
)

type JsonGttFunc func() interface{}

type SimpleJsonGttFunc func() *simplejson.Json

type DBGttFunc func() string

// SimpleCache 缓存
type SimpleCache struct {
	Operation        *StringOperation  // 操作类
	Expire           time.Duration     // 过去时间
	DBGetter         DBGttFunc         // 缓存不存在的操作 DB
	JsonGetter       JsonGttFunc       // 缓存不存在的操作 JSON
	SimpleJsonGetter SimpleJsonGttFunc // 缓存不存在的操作 SimpleJson
	Serializer       string            // 序列化方式
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serializer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serializer: serializer}
}

// SetCache 设置缓存
func (c *SimpleCache) SetCache(key string, value interface{}) {
	c.Operation.Set(key, value, WithExpire(c.Expire)).Unwrap()
}

// GetCache 获取缓存
func (c *SimpleCache) GetCache(key string) (ret interface{}) {
	if c.Serializer == SerializerJson {
		f := func() string {
			obj := c.JsonGetter()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
	} else if c.Serializer == SerializerString {
		f := func() string {
			return c.DBGetter()
		}
		ret = c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
	} else if c.Serializer == SerializerSimpleJson {
		f := func() string {
			obj := c.SimpleJsonGetter()
			encode, err := obj.Encode()
			if err != nil {
				return ""
			}
			return string(encode)
		}
		ret = c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
	}
	return
}

// GetCacheSimpleJson 获取缓存配合SimpleJson插件
func (c *SimpleCache) GetCacheSimpleJson(key string) (js *simplejson.Json) {
	if c.Serializer == SerializerJson {
		f := func() string {
			obj := c.JsonGetter()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret := c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
		js, _ = simplejson.NewJson([]byte(ret))
	} else if c.Serializer == SerializerSimpleJson {
		f := func() string {
			obj := c.SimpleJsonGetter()
			encode, err := obj.Encode()
			if err != nil {
				return ""
			}
			return string(encode)
		}
		ret := c.Operation.Get(key).UnwrapOrElse(f)
		c.SetCache(key, ret)
		js, _ = simplejson.NewJson([]byte(ret))
	}
	return
}
