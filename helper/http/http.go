package http

import (
	"encoding/json"
	"gitee.com/dtapps/go-library/helper/request"
	"github.com/bitly/go-simplejson"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

// GetJson 发起Get请求
func GetJson(url string, data string, headers map[string]interface{}) (res []byte, err error) {
	// 创建一个 HTTP 客户端cli
	cli := gentleman.New()
	// 设置要请求的 URL 基础地址
	cli.URL(url + "?" + data)
	// 创建一个请求对象req
	req := cli.Request()
	req.SetHeader("Content-Type", "application/json; charset=utf-8")
	for key, value := range headers {
		// 设置请求首部（Header）
		req.SetHeader(key, value.(string))
	}
	req.SetHeader("User-Agent", request.GetUserAgent())
	// 发送请求，获取响应对象res
	response, err := req.Send()
	if err != nil {
		panic(err)
	}
	res = response.Bytes()
	if json.Valid(res) == false {
		panic("http server json error GetJson " + url + "?" + data)
	}
	return
}

// PostJson 发起POST请求
func PostJson(url string, data map[string]interface{}, headers map[string]interface{}) (res []byte, err error) {
	// 创建一个 HTTP 客户端cli
	cli := gentleman.New()
	// 设置要请求的 URL 基础地址
	cli.URL(url)
	cli.Use(body.JSON(data))
	// 创建一个请求对象req
	req := cli.Request()
	req.Method("POST")
	req.SetHeader("Content-Type", "application/json; charset=utf-8")
	for key, value := range headers {
		// 设置请求首部（Header）
		req.SetHeader(key, value.(string))
	}
	req.SetHeader("User-Agent", request.GetUserAgent())
	// 发送请求，获取响应对象res
	response, err := req.Send()
	if err != nil {
		panic(err)
	}
	res = response.Bytes()
	if json.Valid(res) == false {
		panic("http server json error PostJson " + url)
	}
	return
}

// PostXml 发起POST请求
func PostXml(url string, data map[string]string, headers map[string]interface{}) (res []byte, err error) {
	// 创建一个 HTTP 客户端cli
	cli := gentleman.New()
	// 设置要请求的 URL 基础地址
	cli.URL(url)
	cli.Use(body.XML(data))
	// 创建一个请求对象req
	req := cli.Request()
	req.Method("POST")
	for key, value := range headers {
		// 设置请求首部（Header）
		req.SetHeader(key, value.(string))
	}
	req.SetHeader("User-Agent", request.GetUserAgent())
	// 发送请求，获取响应对象res
	response, err := req.Send()
	if err != nil {
		panic(err)
	}
	res = response.Bytes()
	return
}

// GetResponseBytes 兼容多层 Key 读取
func GetResponseBytes(data []byte, keys ...string) (b []byte, err error) {
	js, err := simplejson.NewJson(data)
	if err != nil {
		return
	}
	for _, key := range keys {
		js = js.Get(key)
	}
	b, err = js.Encode()
	return
}

// GetResponseArrayIndexBytes 兼容多层 Key 读取某个
func GetResponseArrayIndexBytes(data []byte, index int, keys ...string) (b []byte, err error) {
	js, err := simplejson.NewJson(data)
	if err != nil {
		return
	}
	for _, key := range keys {
		js = js.Get(key)
	}

	js = js.GetIndex(index)

	b, err = js.Encode()
	return
}
