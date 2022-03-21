package gohttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/goheader"
	"github.com/dtapps/go-library/utils/gorequest"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Response struct {
	Status        string
	StatusCode    int
	Header        http.Header
	Body          []byte
	ContentLength int64
}

func Get(url string, params map[string]interface{}) (httpResponse Response, err error) {
	// 创建 http 客户端
	client := &http.Client{}
	// 创建请求
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if len(params) > 0 {
		// GET 请求携带查询参数
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, getString(v))
		}
		req.URL.RawQuery = q.Encode()
	}
	// 设置请求头
	req.Header.Set("User-Agent", gorequest.GetRandomUserAgent())
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		// 格式化返回错误
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}
	// 最后关闭连接
	defer resp.Body.Close()
	// 读取内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}
	httpResponse.Status = resp.Status
	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Header = resp.Header
	httpResponse.Body = respBody
	httpResponse.ContentLength = resp.ContentLength
	return httpResponse, err
}

func GetJsonHeader(url string, params map[string]interface{}, headers goheader.Headers) (httpResponse Response, err error) {
	// 创建 http 客户端
	client := &http.Client{}
	// 创建请求
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if len(params) > 0 {
		// GET 请求携带查询参数
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, getString(v))
		}
		req.URL.RawQuery = q.Encode()
	}
	// 设置请求头
	req.Header.Set("User-Agent", gorequest.GetRandomUserAgent())
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		// 格式化返回错误
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}
	// 最后关闭连接
	defer resp.Body.Close()
	// 读取内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}
	httpResponse.Status = resp.Status
	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Header = resp.Header
	httpResponse.Body = respBody
	httpResponse.ContentLength = resp.ContentLength
	return httpResponse, err
}

func PostForm(targetUrl string, params map[string]interface{}) (httpResponse Response, err error) {
	// 创建 http 客户端
	client := &http.Client{}
	// 携带 form 参数
	form := url.Values{}
	if len(params) > 0 {
		for k, v := range params {
			form.Add(k, getString(v))
		}
	}
	// 创建请求
	req, _ := http.NewRequest(http.MethodPost, targetUrl, strings.NewReader(form.Encode()))
	// 设置请求头
	req.Header.Set("User-Agent", gorequest.GetRandomUserAgent())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		// 格式化返回错误
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}
	// 最后关闭连接
	defer resp.Body.Close()
	// 读取内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}
	httpResponse.Status = resp.Status
	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Header = resp.Header
	httpResponse.Body = respBody
	httpResponse.ContentLength = resp.ContentLength
	return httpResponse, err
}

func PostJson(targetUrl string, paramsStr []byte) (httpResponse Response, err error) {
	// 创建请求
	req, _ := http.NewRequest(http.MethodPost, targetUrl, bytes.NewBuffer(paramsStr))
	// 设置请求头
	req.Header.Set("User-Agent", gorequest.GetRandomUserAgent())
	req.Header.Set("Content-Type", "application/json")
	// 创建 http 客户端
	client := &http.Client{}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		// 格式化返回错误
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}
	// 最后关闭连接
	defer resp.Body.Close()
	// 读取内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}
	httpResponse.Status = resp.Status
	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Header = resp.Header
	httpResponse.Body = respBody
	httpResponse.ContentLength = resp.ContentLength
	return httpResponse, err
}

func PostJsonHeader(targetUrl string, paramsStr []byte, headers goheader.Headers) (httpResponse Response, err error) {
	// 创建请求
	req, _ := http.NewRequest(http.MethodPost, targetUrl, bytes.NewBuffer(paramsStr))
	// 设置请求头
	req.Header.Set("User-Agent", gorequest.GetRandomUserAgent())
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}
	// 创建 http 客户端
	client := &http.Client{}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		// 格式化返回错误
		return httpResponse, errors.New(fmt.Sprintf("请求出错 %s", err))
	}
	// 最后关闭连接
	defer resp.Body.Close()
	// 读取内容
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return httpResponse, errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}
	httpResponse.Status = resp.Status
	httpResponse.StatusCode = resp.StatusCode
	httpResponse.Header = resp.Header
	httpResponse.Body = respBody
	httpResponse.ContentLength = resp.ContentLength
	return httpResponse, err
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		marshal, _ := json.Marshal(v)
		return string(marshal)
	}
}
