package dredis

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	v20210726 "gopkg.in/dtapps/go-library.v2/dredis/v20210726"
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {

}

func client() {
	// 连接
	err := v20210726.InitRedis("127.0.0.1", 6379, "", 2)
	if err != nil {
		fmt.Printf("err：%v\n", err)
	}
	jsonSimpleJson()
}

func set() {
	// 设置
	v20210726.NewStringOperation().Set("test", "test", v20210726.WithExpire(time.Second*1))
}

func mGet() {
	// 获取
	iter := v20210726.NewStringOperation().MGet("test1", "test2").Iter()
	for iter.HasNext() {
		log.Println("MGet：", iter.Next())
	}
}

func json() {
	newCache := v20210726.NewSimpleCache(v20210726.NewStringOperation(), time.Second*10, v20210726.SerializerJson)
	newCache.JsonGetter = func() interface{} {
		log.Println("【没有命中】SerializerJson")
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	cacheJSon := newCache.GetCache("test123")
	log.Printf("【GetCache】cacheJSon：%v\n", cacheJSon)
}

func dbString() {
	newCache := v20210726.NewSimpleCache(v20210726.NewStringOperation(), time.Second*10, v20210726.SerializerString)
	newCache.DBGetter = func() string {
		log.Println("【没有命中】SerializerString")
		return "data by id=123"
	}
	cacheString := newCache.GetCache("test456")
	log.Printf("【GetCache】cacheString：%v\n", cacheString)
}

func simpleJson() {
	newCache := v20210726.NewSimpleCache(v20210726.NewStringOperation(), time.Second*50, v20210726.SerializerSimpleJson)
	newCache.SimpleJsonGetter = func() *simplejson.Json {
		log.Println("_test【没有命中】SerializerSimpleJson")
		js := simplejson.New()
		js.Set("name", "test")
		return js
	}
	cacheSimpleJson := newCache.GetCacheSimpleJson("test789")
	log.Printf("_test【GetCache】cacheSimpleJson：%v\n", cacheSimpleJson.Get("name"))
}

func jsonSimpleJson() {
	newCache := v20210726.NewSimpleCache(v20210726.NewStringOperation(), time.Second*50, v20210726.SerializerJson)
	newCache.JsonGetter = func() interface{} {
		log.Println("【没有命中】SerializerJson")
		type a []string
		b := a{
			"me", "she", "you",
		}
		return b
	}
	cacheJson := newCache.GetCacheSimpleJson("test789")
	log.Printf("_test【JsonGetter GetCacheSimpleJson】jsonSimpleJson：%v\n", cacheJson.GetIndex(1))
}
