package gorequest

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	cookiemonster "github.com/MercuryEngineering/CookieMonster"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Response 返回内容
type Response struct {
	RequestID             string      `json:"request_id"`              // 请求编号
	RequestUri            string      `json:"request_uri"`             // 请求链接
	RequestParams         *Params     `json:"request_params"`          // 请求参数
	RequestMethod         string      `json:"request_method"`          // 请求方式
	RequestHeader         *Headers    `json:"request_header"`          // 请求头部
	RequestCookie         string      `json:"request_cookie"`          // 请求Cookie
	RequestTime           time.Time   `json:"request_time"`            // 请求时间
	RequestCostTime       int64       `json:"request_cost_time"`       // 请求消耗时长
	ResponseHeader        http.Header `json:"response_header"`         // 响应头部
	ResponseStatus        string      `json:"response_status"`         // 响应状态
	ResponseStatusCode    int         `json:"response_status_code"`    // 响应状态码
	ResponseBody          []byte      `json:"response_body"`           // 响应内容
	ResponseContentLength int64       `json:"response_content_length"` // 响应大小
	ResponseTime          time.Time   `json:"response_time"`           // 响应时间
}

// LogFunc 日志函数
type LogFunc func(ctx context.Context, response *LogResponse)

// App 实例
type App struct {
	Uri                          string           // 全局请求地址，没有设置url才会使用
	httpUri                      string           // 请求地址
	httpMethod                   string           // 请求方法
	httpHeader                   *Headers         // 请求头
	httpParams                   *Params          // 请求参数
	httpCookie                   string           // Cookie
	responseContent              Response         // 返回内容
	httpContentType              string           // 请求内容类型
	p12Cert                      *tls.Certificate // p12证书内容
	tlsMinVersion, tlsMaxVersion uint16           // TLS版本
	clientIP                     string           // 客户端IP
	logFunc                      LogFunc          // 日志记录函数
}

// NewHttp 实例化
func NewHttp() *App {
	c := &App{
		httpHeader: NewHeaders(),
		httpParams: NewParams(),
	}
	return c
}

// Get 发起 GET 请求
func (c *App) Get(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodGet
	return request(c, ctx)
}

// Head 发起 HEAD 请求
func (c *App) Head(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodHead
	return request(c, ctx)
}

// Post 发起 POST 请求
func (c *App) Post(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodPost
	return request(c, ctx)
}

// Put 发起 PUT 请求
func (c *App) Put(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodPut
	return request(c, ctx)
}

// Patch 发起 PATCH 请求
func (c *App) Patch(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodPatch
	return request(c, ctx)
}

// Delete 发起 DELETE 请求
func (c *App) Delete(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodDelete
	return request(c, ctx)
}

// Connect 发起 CONNECT 请求
func (c *App) Connect(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodConnect
	return request(c, ctx)
}

// Options 发起 OPTIONS 请求
func (c *App) Options(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodOptions
	return request(c, ctx)
}

// Trace 发起 TRACE 请求
func (c *App) Trace(ctx context.Context, uri ...string) (Response, error) {
	if len(uri) == 1 {
		c.Uri = uri[0]
	}
	// 设置请求方法
	c.httpMethod = http.MethodTrace
	return request(c, ctx)
}

// Request 发起请求
func (c *App) Request(ctx context.Context) (Response, error) {
	return request(c, ctx)
}

// 请求接口
func request(c *App, ctx context.Context) (httpResponse Response, err error) {

	// 开始时间
	start := time.Now().UTC()

	// 赋值
	httpResponse.RequestTime = time.Now()
	httpResponse.RequestUri = c.httpUri
	httpResponse.RequestMethod = c.httpMethod
	httpResponse.RequestParams = c.httpParams.DeepCopy()
	httpResponse.RequestHeader = c.httpHeader.DeepCopy()
	httpResponse.RequestCookie = c.httpCookie

	// 判断网址
	if httpResponse.RequestUri == "" {
		httpResponse.RequestUri = c.Uri
	}
	if httpResponse.RequestUri == "" {
		err = errors.New("没有请求地址")
		return httpResponse, err
	}

	// 创建 http 客户端
	client := &http.Client{}

	transportStatus := false
	transport := &http.Transport{}
	transportTls := &tls.Config{}

	if c.p12Cert != nil {
		transportStatus = true
		// 配置
		transportTls.Certificates = []tls.Certificate{*c.p12Cert}
		transport.DisableCompression = true
	}

	if c.tlsMinVersion != 0 && c.tlsMaxVersion != 0 {
		transportStatus = true
		// 配置
		transportTls.MinVersion = c.tlsMinVersion
		transportTls.MaxVersion = c.tlsMaxVersion
	}

	if transportStatus {
		transport.TLSClientConfig = transportTls
		client = &http.Client{
			Transport: transport,
		}
	}

	// 请求类型
	if c.httpContentType == "" {
		c.httpContentType = httpParamsModeJson
	}
	switch c.httpContentType {
	case httpParamsModeJson:
		httpResponse.RequestHeader.Set("Content-Type", "application/json")
	case httpParamsModeForm:
		httpResponse.RequestHeader.Set("Content-Type", "application/x-www-form-urlencoded")
	case httpParamsModeXml:
		httpResponse.RequestHeader.Set("Content-Type", "text/xml")
	}

	// 请求编号
	httpResponse.RequestID = GetRequestIDContext(ctx)
	if httpResponse.RequestID != "" {
		httpResponse.RequestHeader.Set(XRequestID, httpResponse.RequestID)
	}

	// 请求内容
	var requestBody io.Reader

	if httpResponse.RequestMethod != http.MethodGet && c.httpContentType == httpParamsModeJson {
		jsonStr, err := json.Marshal(httpResponse.RequestParams.DeepGetAny())
		if err != nil {
			return httpResponse, err
		}
		// 赋值
		requestBody = bytes.NewBuffer(jsonStr)
	}

	if httpResponse.RequestMethod != http.MethodGet && c.httpContentType == httpParamsModeForm {
		// 携带 form 参数
		form := url.Values{}
		for k, v := range httpResponse.RequestParams.DeepGetString() {
			form.Add(k, v)
		}
		// 赋值
		requestBody = strings.NewReader(form.Encode())
	}

	if c.httpContentType == httpParamsModeXml {
		requestBody, err = ToXml(httpResponse.RequestParams.DeepGetAny())
		if err != nil {
			return httpResponse, err
		}
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, httpResponse.RequestMethod, httpResponse.RequestUri, requestBody)
	if err != nil {
		return httpResponse, err
	}

	// GET 请求携带查询参数
	if httpResponse.RequestMethod == http.MethodGet {
		q := req.URL.Query()
		for k, v := range httpResponse.RequestParams.DeepGetString() {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	// 设置请求头
	if len(httpResponse.RequestHeader.DeepGetString()) > 0 {
		for k, v := range httpResponse.RequestHeader.DeepGetString() {
			req.Header.Set(k, v)
		}
	}

	// 设置Cookie
	if httpResponse.RequestCookie != "" {
		cookies, err := cookiemonster.ParseString(httpResponse.RequestCookie)
		if err == nil {
			if len(cookies) > 0 {
				for _, c := range cookies {
					req.AddCookie(c)
				}
			} else {
				req.Header.Set("Cookie", httpResponse.RequestCookie)
			}
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return httpResponse, err
	}
	defer resp.Body.Close() // 关闭连接

	// 结束时间
	end := time.Now().UTC()

	httpResponse.RequestCostTime = end.Sub(start).Milliseconds() // 请求消耗时长

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			slog.ErrorContext(ctx, "[gorequest] gzip.NewReader",
				slog.String("err", err.Error()),
			)
		}
	case "deflate":
		reader = flate.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	defer reader.Close() // nolint

	// 读取内容
	body, err := io.ReadAll(reader)
	if err != nil {
		return httpResponse, err
	}

	// 赋值
	httpResponse.ResponseTime = time.Now()
	httpResponse.ResponseStatus = resp.Status
	httpResponse.ResponseStatusCode = resp.StatusCode
	httpResponse.ResponseHeader = resp.Header
	httpResponse.ResponseBody = body
	httpResponse.ResponseContentLength = resp.ContentLength

	// 调用日志记录函数
	if c.logFunc != nil {
		var responseBodyJson map[string]any
		_ = json.Unmarshal([]byte(httpResponse.ResponseBody), &responseBodyJson)
		c.logFunc(ctx, &LogResponse{
			RequestID:       httpResponse.RequestID,
			RequestTime:     httpResponse.RequestTime,
			RequestHost:     req.Host,
			RequestPath:     req.URL.String(),
			RequestQuery:    req.URL.Query(),
			RequestMethod:   req.Method,
			RequestIP:       c.clientIP,
			RequestBody:     httpResponse.RequestParams.DeepGetAny(),
			RequestHeader:   req.Header,
			RequestCostTime: httpResponse.RequestCostTime,

			ResponseTime:     httpResponse.ResponseTime,
			ResponseHeader:   resp.Header,
			ResponseCode:     resp.StatusCode,
			ResponseBody:     string(httpResponse.ResponseBody),
			ResponseBodyJson: responseBodyJson,
		})
	}

	return httpResponse, err
}
