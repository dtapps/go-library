package goredis

import (
	"testing"
)

func TestSet(t *testing.T) {
	//// 连接
	//err := InitRedis("127.0.0.1", 6379, "", 2)
	//if err != nil {
	//	fmt.Printf("err：%v\n", err)
	//}
	//// 设置
	//NewStringOperation().Set("test", "test", WithExpire(time.Second*1))
}

func TestMGet(t *testing.T) {
	//// 连接
	//err := InitRedis("127.0.0.1", 6379, "", 2)
	//if err != nil {
	//	fmt.Printf("err：%v\n", err)
	//}
	//// 获取
	//iter := NewStringOperation().MGet("test1", "test2").Iter()
	//for iter.HasNext() {
	//	fmt.Println("MGet：", iter.Next())
	//}
}

func TestJson(t *testing.T) {
	//// 连接
	//err := InitRedis("127.0.0.1", 6379, "", 2)
	//if err != nil {
	//	fmt.Printf("err：%v\n", err)
	//}
	//newCache := NewSimpleCache(NewStringOperation(), time.Second*10, SerializerJson)
	//newCache.JsonGetter = func() interface{} {
	//	fmt.Println("【没有命中】SerializerJson")
	//	type a []string
	//	b := a{
	//		"me", "she", "you",
	//	}
	//	return b
	//}
	//cacheJSon := newCache.GetCache("test123")
	//fmt.Printf("【GetCache】cacheJSon：%v\n", cacheJSon)
}

func TestDbString(t *testing.T) {
	//// 连接
	//err := InitRedis("127.0.0.1", 6379, "", 2)
	//if err != nil {
	//	fmt.Printf("err：%v\n", err)
	//}
	//newCache := NewSimpleCache(NewStringOperation(), time.Second*10, SerializerString)
	//newCache.DBGetter = func() string {
	//	fmt.Println("【没有命中】SerializerString")
	//	return "data by id=123"
	//}
	//cacheString := newCache.GetCache("test456")
	//fmt.Printf("【GetCache】cacheString：%v\n", cacheString)
}
