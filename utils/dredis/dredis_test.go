package dredis

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
	"time"
)

func TestName(t *testing.T) {

}

func client() {
	// 连接
	err := InitRedis("127.0.0.1", 6379, "", 2)
	if err != nil {
		fmt.Printf("err：%v\n", err)
	}
	jsonSimpleJson()
}

func set() {
	// 设置
	NewStringOperation().Set("test", "test", WithExpire(time.Second*1))
}

func mGet() {
	// 获取
	iter := NewStringOperation().MGet("test1", "test2").Iter()
	for iter.HasNext() {
		fmt.Println("MGet：", iter.Next())
	}
}

func json() {
	newCache := NewSimpleCache(NewStringOperation(), time.Second*10, SerializerJson)
	newCache.JsonGetter = func() interface{} {
		fmt.Println("【没有命中】SerializerJson")
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	cacheJSon := newCache.GetCache("test123")
	fmt.Printf("【GetCache】cacheJSon：%v\n", cacheJSon)
}

func dbString() {
	newCache := NewSimpleCache(NewStringOperation(), time.Second*10, SerializerString)
	newCache.DBGetter = func() string {
		fmt.Println("【没有命中】SerializerString")
		return "data by id=123"
	}
	cacheString := newCache.GetCache("test456")
	fmt.Printf("【GetCache】cacheString：%v\n", cacheString)
}

func simpleJson() {
	newCache := NewSimpleCache(NewStringOperation(), time.Second*50, SerializerSimpleJson)
	newCache.SimpleJsonGetter = func() *simplejson.Json {
		fmt.Println("_test【没有命中】SerializerSimpleJson")
		js := simplejson.New()
		js.Set("name", "test")
		return js
	}
	cacheSimpleJson := newCache.GetCacheSimpleJson("test789")
	fmt.Printf("_test【GetCache】cacheSimpleJson：%v\n", cacheSimpleJson.Get("name"))
}

func jsonSimpleJson() {
	newCache := NewSimpleCache(NewStringOperation(), time.Second*50, SerializerJson)
	newCache.JsonGetter = func() interface{} {
		fmt.Println("【没有命中】SerializerJson")
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	cacheJson := newCache.GetCacheSimpleJson("test789")
	fmt.Printf("_test【JsonGetter GetCacheSimpleJson】jsonSimpleJson：%v\n", cacheJson.GetIndex(1))
}
